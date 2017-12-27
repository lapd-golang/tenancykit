package pkg

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// CreateUserSession defines the set of data received to create a new user's session.
type CreateUserSession struct {
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	Expiration time.Duration `json:"expiration"`
}

// Validate attempts to validate struct has appropriate data.
func (cus CreateUserSession) Validate() error {
	if cus.Email == "" {
		return errors.New("Email is required")
	}
	if cus.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}

// EndUserSession defines the set of data received to end a user's session.
type EndUserSession struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

// Validate attempts to validate struct has appropriate data.
func (eus EndUserSession) Validate() error {
	if eus.UserID == "" {
		return errors.New("UserID is required")
	}
	if eus.Token == "" {
		return errors.New("User Token is required")
	}
	return nil
}

// UserSession embodies the data which is used to identify a logged in user's session.
// @mongoapi(ENVName => TENANCYKIT)
//
// @httpapi(New => CreateUserSession)
type UserSession struct {
	UserID   string    `json:"user_id"`
	PublicID string    `json:"public_id"`
	Token    string    `json:"token"`
	Expires  time.Time `json:"expires"`
}

// NewUserSession returns a new instance of a UserSession.
func NewUserSession(userID string, expiration time.Time) UserSession {
	return UserSession{
		UserID:   userID,
		PublicID: uuid.NewV4().String(),
		Token:    uuid.NewV4().String(),
		Expires:  expiration,
	}
}

// ValidateToken validates the provide base64 encoded token, that it matches the
// expected token value with that of the session.
func (us UserSession) ValidateToken(token string) bool {
	return us.Token == token
}

// UserSessionToken returns the UserSession.Token has a base64 encoded string.
// It returns a base64 encoded version where it contains the UserID:UserSessionToken.
func (us UserSession) UserSessionToken() string {
	token := []byte(fmt.Sprintf("%s:%s", us.UserID, us.Token))
	return base64.StdEncoding.EncodeToString(token)
}

// Expired returns true/false if the the given session is expired.
func (us UserSession) Expired() bool {
	return time.Now().After(us.Expires)
}

// Consume consumes data from map into instance fieldus.
func (us *UserSession) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", us, data)
}

// Fields returns a map containing the instance fields as key-value pairus.
func (us UserSession) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &us)
}
