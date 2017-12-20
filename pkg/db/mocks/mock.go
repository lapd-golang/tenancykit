package mocks

import "errors"

// ErrNotFound identifies giving item has an error not found. Useful has a generic not found error.
var ErrNotFound = errors.New("not found")
