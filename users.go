package tenancykit

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// UpdateUser defines the set of data sent when updating a users password.
type UpdateUser struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

// CreateUser holds details necessary for creating a new user.
type CreateUser struct {
	TenantID        string `json:"tenant_id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

// User is a type defining the given user related fields for a given.
// @mongoapi
// @sqlapi
// @httpapi(New => CreateUser, Update => UpdateUser)
type User struct {
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	PublicID      string    `json:"public_id"`
	TenantID      string    `json:"tenant_id"`
	PrivateID     string    `json:"private_id"`
	Hash          string    `json:"hash"`
	TwoFactorAuth bool      `json:"two_factor_auth"`
	Created       time.Time `json:"created_at"`
	Updated       time.Time `json:"updated_at"`
}

// NewUser returns a new User instance based on the provided data.
func NewUser(nw CreateUser) (User, error) {
	var u User
	u.Email = nw.Email
	u.Created = time.Now()
	u.Updated = time.Now()
	u.TenantID = nw.TenantID
	u.Username = nw.Username
	u.PublicID = uuid.NewV4().String()
	u.PrivateID = uuid.NewV4().String()
	if err := u.ChangePassword(nw.Password); err != nil {
		return u, err
	}
	return u, nil
}

// Authenticate attempts to authenticate the giving password to the provided user.
func (u User) Authenticate(password string) error {
	pass := []byte(u.PrivateID + ":" + password)
	return bcrypt.CompareHashAndPassword([]byte(u.Hash), pass)
}

// FilteredFields returns maps where all secret data as being stripped
// for safe public viewing.
func (u User) FilteredFields() (map[string]interface{}, error) {
	fields, err := u.Fields()
	if err != nil {
		return nil, err
	}

	delete(fields, "hash")
	delete(fields, "private_id")
	return fields, nil
}

// Consume consumes data from map into instance fields.
func (u *User) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", u, data)
}

// Fields returns a map containing the instance fields as key-value pairs.
func (u User) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &u)
}

// ChangePassword uses the provided password to set the users password hash.
func (u *User) ChangePassword(password string) error {
	pass := []byte(u.PrivateID + ":" + password)
	hash, err := bcrypt.GenerateFromPassword(pass, UserHashComplexity)
	if err != nil {
		return err
	}

	u.Hash = string(hash)
	return nil
}
