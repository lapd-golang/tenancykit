package pkg

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserAuthorization embodies authenticate data sent to the
// user after login.
type UserAuthorization struct {
	Token string `json:"token"`
}

// UpdateUser defines the set of data sent when updating a users password.
type UpdateUserPassword struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

// Validate returns error if UpdateUser fails to match its value
// requirements.
func (u UpdateUserPassword) Validate() error {
	if u.Password == "" {
		return errors.New("Password can not be empty")
	}

	if u.PasswordConfirm == "" {
		return errors.New("PasswordConfirm can not be empty")
	}

	if u.Password != u.PasswordConfirm {
		return errors.New("Password must match PasswordConfirm exactly")
	}

	return nil
}

// CreateUser holds details necessary for creating a new user.
type CreateUser struct {
	TenantID        string `json:"tenant_id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	TwoFactorAuth   bool   `json:"twofactor_auth"`
}

// Validate returns error if any field does not match requirements.
func (cu CreateUser) Validate(multitenant bool) error {
	if cu.TenantID == "" && multitenant {
		return errors.New("TenantID is required")
	}

	if cu.Username == "" {
		return errors.New("Username is required")
	}

	if cu.Email == "" {
		return errors.New("Email is required")
	}

	if cu.Password == "" {
		return errors.New("Password can not be empty")
	}

	if cu.PasswordConfirm == "" {
		return errors.New("PasswordConfirm can not be empty")
	}

	if cu.Password != cu.PasswordConfirm {
		return errors.New("Password must match PasswordConfirm exactly")
	}

	return nil
}

// User is a type defining the given user related fields for a given.
// @mongoapi(ENVName => TENANCYKIT)
// @httpapi(New => CreateUser)
type User struct {
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	PublicID      string    `json:"public_id"`
	TenantID      string    `json:"tenant_id"`
	PrivateID     string    `json:"private_id"`
	Hash          string    `json:"hash"`
	Roles         []string  `json:"user_roles"`
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
	u.TwoFactorAuth = nw.TwoFactorAuth
	u.PublicID = uuid.NewV4().String()
	u.PrivateID = uuid.NewV4().String()
	if err := u.ChangePassword(nw.Password); err != nil {
		return u, err
	}
	return u, nil
}

// Validate returns an error if user is in a invalid state.
func (u User) Validate() error {
	return u.ValidateUser(false)
}

// ValidateUser returns an error if user is in a invalid state.
func (u User) ValidateUser(multitenant bool) error {
	if multitenant && u.TenantID == "" {
		return errors.New("user must have tenant_id")
	}
	if u.Created.IsZero() {
		return errors.New("user must have an created time")
	}
	if u.Updated.IsZero() {
		return errors.New("user must have an update time")
	}
	if u.PrivateID == "" {
		return errors.New("user must have an private_id")
	}
	if u.PublicID == "" {
		return errors.New("user must have an public_id")
	}
	if u.Username == "" {
		return errors.New("user must have an username")
	}
	if u.Email == "" {
		return errors.New("user must have an email")
	}
	return nil
}

// HasAllRole returns true/false if giving user has all provided roles.
func (u User) HasAllRole(roles ...Role) bool {
	for _, role := range roles {
		if !u.hasRole(role) {
			return false
		}
	}
	return true
}

// HasAnyrole returns true/false if giving user has any of provided roles.
func (u User) HasAnyRole(roles ...Role) bool {
	for _, role := range roles {
		if u.hasRole(role) {
			return true
		}
	}
	return false
}

// AddRole adds provided role public_id to user instance.
// Granting user all powers related to giving role.
func (u *User) AddRole(role Role) {
	u.Roles = append(u.Roles, role.PublicID)
}

// hasRole returns true/false if user has provided role.
func (u User) hasRole(r Role) bool {
	for _, role := range u.Roles {
		if role == r.PublicID {
			return true
		}
	}
	return false
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
