package pkg

import (
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sec51/twofactor"
)

// errors ...
var (
	ErrInvalidUserCodeLength = errors.New("user code length does not match expected")
)

const (
	// GoogleAuthenticatorUserCodeLength sets the length expected by
	// google authenticator
	GoogleAuthenticatorUserCodeLength = 6
)

// NewTF defines a struct type which contains necessary field data
// needed for the creation of TFRecords.
type NewTF struct {
	MaxLength int    `json:"max_length"`
	User      User   `json:"user"`
	Domain    string `json:"domain"`
}

// TFRecord embodies the record of a tenants's or user's twofactor
// authorization data. It stores the generated data and security details related
// to the account and is used to validate users code after login if enabled.
// TwoFactor is enabled if user has associated record in twofactor table or collection.
//
// @mongoapi
// @sqlapi
// @httpapi(New => NewTF)
type TFRecord struct {
	Key        string    `json:"key"`
	Domain     string    `json:"domain"`
	UserID     string    `json:"user_id"`
	TOTP       string    `json:"totp"`
	PublicID   string    `json:"tenant_id"`
	CodeLength int       `json:"code_length"`
	Created    time.Time `json:"created_at"`
	Updated    time.Time `json:"updated_at"`
}

// NewTFRecord returns a new TFRecord which is created for login authentication using the
// user's and tenant's record. It requires providing domain name to associated with
// record as it affects key display on google authenticator.
func NewTFRecord(codeLen int, domainName string, user User) (TFRecord, error) {
	var record TFRecord
	record.Domain = domainName
	record.UserID = user.PublicID
	record.CodeLength = codeLen
	record.PublicID = uuid.NewV4().String()
	record.Key = fmt.Sprintf("%s - (%s)", domainName, user.Email)
	record.Created = time.Now()
	record.Updated = record.Created

	totp, err := twofactor.NewTOTP(record.UserID, record.Key, crypto.SHA512, codeLen)
	if err != nil {
		return record, err
	}

	if err := record.serializeAndUpdateTOTP(totp); err != nil {
		return record, err
	}

	return record, nil
}

// ValidateOTP validates provided OTP user code from external source.
func (tf *TFRecord) ValidateOTP(userCode string) error {
	if codeLen := len(userCode); codeLen > tf.CodeLength || codeLen < tf.CodeLength {
		return ErrInvalidUserCodeLength
	}

	totp, err := tf.deserializeTOTP()
	if err != nil {
		return err
	}

	if err = totp.Validate(userCode); err != nil {
		return err
	}

	return tf.serializeAndUpdateTOTP(totp)
}

// OTP returns the next one-time-password from the
func (tf *TFRecord) OTP() (string, error) {
	totp, err := tf.deserializeTOTP()
	if err != nil {
		return "", err
	}

	otp, err := totp.OTP()
	if err != nil {
		return "", err
	}

	return otp, tf.serializeAndUpdateTOTP(totp)
}

// SecretCode returns QR secret code for TFRecord.
func (tf *TFRecord) SecetCode() (string, error) {
	totp, err := tf.deserializeTOTP()
	if err != nil {
		return "", err
	}

	return totp.Secret(), nil
}

// QR returns QR png byte slice for TFRecord.
func (tf *TFRecord) QR() ([]byte, error) {
	totp, err := tf.deserializeTOTP()
	if err != nil {
		return nil, err
	}

	return totp.QR()
}

// serializeAndUpdateTOTP takes provided twofactor.Totp instance, serializing
// and updating it's internal reference with provided one.
func (tf *TFRecord) serializeAndUpdateTOTP(totp *twofactor.Totp) error {
	totpbytes, err := totp.ToBytes()
	if err != nil {
		return err
	}

	tf.TOTP = base64.StdEncoding.EncodeToString(totpbytes)
	return nil
}

// deserializeTOTP attempts to deserialize internal totp data into twofactor.Totp
// struct.
func (tf TFRecord) deserializeTOTP() (*twofactor.Totp, error) {
	data, err := base64.StdEncoding.DecodeString(tf.TOTP)
	if err != nil {
		return nil, err
	}

	totp, err := twofactor.TOTPFromBytes(data, tf.Key)
	if err != nil {
		return nil, err
	}

	return totp, nil
}

// Consume consumes data from map into instance fields.
func (tf *TFRecord) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", tf, data)
}

// Fields returns a map containing the instance fields as key-value pairs.
func (tf TFRecord) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &tf)
}

// FilteredFields returns maps where all secret data as being stripped
// for safe public viewing.
func (tf TFRecord) FilteredFields() (map[string]interface{}, error) {
	fields, err := tf.Fields()
	if err != nil {
		return nil, err
	}

	delete(fields, "totp")
	return fields, nil
}
