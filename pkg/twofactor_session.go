package pkg

import uuid "github.com/satori/go.uuid"

// TwoFactorSession embodies the data used to representing the
// status of a users twofactor authentication authorization.
// It indicates into db that user has done necessary bits
// to satisfy twofactor authorization. It should be removed
// from db after user logout or immediately after api requests
// for apis.
// @mongoapi
// @httpapi
type TwoFactorSession struct {
	UserID   string `json:"user_id"`
	PublicID string `json:"public_id"`
	Done     bool   `json:"bool"`
}

// NewTwoFactorSession returns new instance of TwoFactorSession.
func NewTwoFactorSession(userID string, status bool) TwoFactorSession {
	return TwoFactorSession{
		Done:     status,
		UserID:   userID,
		PublicID: uuid.NewV4().String(),
	}
}

// Consume consumes data from map into instance fieldtfs.
func (tfs *TwoFactorSession) Consume(data map[string]interface{}) error {
	return mapper.MapTo("json", tfs, data)
}

// Fields returns a map containing the instance fields as key-value pairtfs.
func (tfs TwoFactorSession) Fields() (map[string]interface{}, error) {
	return mapper.MapFrom("json", &tfs)
}
