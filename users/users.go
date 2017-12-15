package users

import (
	"reflect"
	"time"

	"github.com/influx6/faux/reflection"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// consts ...
const (
	hashComplexity = 10
)

var (
	mapper = (func() reflection.Mapper {
		layout := "Mon Jan 2 2006 15:04:05 -0700 MST"
		timeType := reflect.TypeOf((*time.Time)(nil))
		baseMapper := reflection.NewStructMapper()
		baseMapper.AddAdapter(timeType, reflection.TimeMapper(layout))
		baseMapper.AddInverseAdapter(timeType, reflection.TimeInverseMapper(layout))
		return baseMapper
	}())
)

// UpdateUserPassword defines the set of data sent when updating a users password.
type UpdateUserPassword struct {
	PublicID        string `json:"public_id"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

// NewUser holds details necessary for creating a new user.
type NewUser struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

// User is a type defining the given user related fields for a given.
// @mongoapi
// @sqlapi
type User struct {
	Username      string    `json:"username"`
	PublicID      string    `json:"public_id"`
	PrivateID     string    `json:"private_id"`
	Hash          string    `json:"hash"`
	TwoFactorAuth bool      `json:"two_factor_auth"`
	Created       time.Time `json:"created_at"`
	Updated       time.Time `json:"updated_at"`
}

// New returns a new User instance based on the provided data.
func New(nw NewUser) (User, error) {
	var u User
	u.Created = time.Now()
	u.Updated = time.Now()
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

// SafeFields returns a map representing the data of the user with important
// security fields removed.
func (u User) SafeFields() map[string]interface{} {
	fields := u.Fields()
	delete(fields, "hash")
	delete(fields, "private_id")
	return fields
}

// Consume sets values of User with provided data map.
func (u *User) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", u, data)
}

// Fields returns a map representing the data of the user.
func (u User) Fields() map[string]interface{} {
	mapped, err := mapper.MapFrom("json", &u)
	if err != nil {
		panic(err)
	}
	return mapped
}

// ChangePassword uses the provided password to set the users password hash.
func (u *User) ChangePassword(password string) error {
	pass := []byte(u.PrivateID + ":" + password)
	hash, err := bcrypt.GenerateFromPassword(pass, hashComplexity)
	if err != nil {
		return err
	}

	u.Hash = string(hash)
	return nil
}
