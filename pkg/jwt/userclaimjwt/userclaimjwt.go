// Package userclaimjwt provides a auto-generated package which contains a API for authentication using JWT.
//
//
package userclaimjwt

import (
	"time"

	"errors"

	"strings"

	"context"

	jwt "github.com/dgrijalva/jwt-go"

	uuid "github.com/satori/go.uuid"

	"github.com/influx6/faux/httputil"

	"github.com/gokit/tokens"

	"github.com/gokit/tenancykit/pkg"
)

// errors ...
var (
	ErrInvalidIdentity      = errors.New("provided Identity is invalid")
	ErrParseJWTToken        = errors.New("parse error: invalid jwt token")
	ErrInternalError        = errors.New("backend failed with error")
	ErrInvalidJWTToken      = errors.New("received jwt token is invalid")
	ErrUnexpectedJWTClaim   = errors.New("jwt token claim is not expected type")
	ErrExpiredJWTToken      = errors.New("received jwt token is expired")
	ErrInvalidRefreshToken  = errors.New("Invalid refresh token")
	ErrTokenRefreshDenied   = errors.New("refresh_token already blacklist")
	ErrExpiredRefreshToken  = errors.New("refresh_token already expired")
	ErrInvalidSigningMethod = errors.New("token signing method mismatched")
)

// Identity embodies data stored for a user's login credentials.
type Identity struct {
	PublicID        string        `json:"public_id"`
	RefreshToken    string        `json:"refresh_token"`
	Expires         time.Time     `json:"expires"`
	TargetID        string        `json:"target_id"`
	LastLogin       time.Time     `json:"last_login"`
	IssuedAt        time.Time     `json:"last_login"`
	RefreshInterval time.Duration `json:"refresh_interval"`
	Data            pkg.UserClaim `json:"data"`
}

// Validate returns an error if giving Identity does not match desired
// field value state.
func (id Identity) Validate() error {
	if strings.TrimSpace(id.PublicID) == "" {
		return errors.New("PublicID is required")
	}

	if strings.TrimSpace(id.RefreshToken) == "" {
		return errors.New("RefreshToken is required")
	}

	if strings.TrimSpace(id.TargetID) == "" {
		return errors.New("TargetID is required")
	}

	return nil
}

// JWTAuth embodies data provided as response to a token refresh or sso
// login operation.
type JWTAuth struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	TokenType      string `json:"token_type"`
	Expires        int64  `json:"expires"`
	RefreshExpires int64  `json:"refresh_expires"`
}

// JWTError embodies data sent as error for a operation.
type JWTError struct {
	Err    error `json:"err"`
	SrcErr error `json:"srcerr"`
}

// Error returns the underline combined src and err error values
// associated with the error instance.
func (t JWTError) Error() string {
	return t.Err.Error() + " (" + t.SrcErr.Error() + ")"
}

// Token embodies data desired for a field for storing key-value pairs in set.
type Token interface {
	ID() string
	Value() string
}

// IdentityDBBackend defines an interface which exposes a underline storage system
// for retrieving and storing identity records.
type IdentityDBBackend interface {
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, publicID string) error
	Create(ctx context.Context, elem Identity) error
	Get(ctx context.Context, publicID string) (Identity, error)
	Update(ctx context.Context, publicID string, elem Identity) error
	GetAllByOrder(ctx context.Context, order string, orderBy string) ([]Identity, error)
	GetByField(ctx context.Context, key string, value interface{}) (Identity, error)
	GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]Identity, int, error)
}

// IdentityBackend defines an interface that expose a backend interface which can
// expose methods that contain all necessary logic for interaction with api for http endpoints.
type IdentityBackend interface {
	Count(context.Context) (int, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (Identity, error)
	Update(context.Context, string, Identity) error
	Revoke(context.Context, string) error
	Refresh(context.Context, string) (JWTAuth, error)
	GetAll(context.Context, string, string, int, int) ([]Identity, int, error)
	Grant(context.Context, pkg.CreateUserSession) (JWTAuth, error)
	Authenticate(context.Context, string) (pkg.UserClaim, error)
}

// IdentityHTTP defines an interface which expose an interface for
// http api for Identity.
type IdentityHTTP interface {
	Get(*httputil.Context) error
	Info(*httputil.Context) error
	Update(*httputil.Context) error
	Delete(*httputil.Context) error
	GetAll(*httputil.Context) error
	Grant(*httputil.Context) error
	Revoke(*httputil.Context) error
	Refresh(*httputil.Context) error
	Authenticate(*httputil.Context) error
}

// IdentityInfo contains meta-data regarding all records in db of type Identity.
type IdentityInfo struct {
	Total int `json:"total"`
}

// Identities defines a type to represent the response given to a request for
// all records of the type pkg.Identity.
type Identities struct {
	Page            int        `json:"page"`
	ResponsePerPage int        `json:"responsePerPage"`
	TotalRecords    int        `json:"total_records"`
	Records         []Identity `json:"records"`
}

// Testimony embodies data returned by user for creation of new identity claim.
type Testimony struct {
	TargetID string
	Data     pkg.UserClaim
}

// JWTClaim embodies the data attached with the standard claims.
type JWTClaim struct {
	jwt.StandardClaims
	Data pkg.UserClaim
}

// IdentityMaker defines a function type provided by maker for generating identity claim.
type IdentityMaker func(context.Context, JWTConfig, pkg.CreateUserSession) (Testimony, error)

// TokenValidator defines a function type provided by user for validating incoming token.
type TokenValidator func(JWTConfig, *jwt.Token) (interface{}, error)

// JWTConfig embodies the field for configuring JWTBackend.
type JWTConfig struct {
	Signer              string
	SigningSecret       string
	Maker               IdentityMaker
	Method              jwt.SigningMethod
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
	Authenticator       TokenValidator
}

// JWTIdentity implements the IdentityBackend interface and embodies all
// necessary method for granting, revoking and freshing jwt access and refresh tokens.
type JWTIdentity struct {
	config    JWTConfig
	blacklist tokens.TokenSet
	IdentityDBBackend
}

// NewJWTIdentity returns a new JWTIdentity instance which embodies and implements the IdentityBackend
// interface.
func NewJWTIdentity(config JWTConfig, blacklist tokens.TokenSet, backend IdentityDBBackend) JWTIdentity {
	return JWTIdentity{config: config, blacklist: blacklist, IdentityDBBackend: backend}
}

// Authenticate attempts to authenticate users access token to validate user's
func (jwa JWTIdentity) Authenticate(ctx context.Context, accessToken string) (pkg.UserClaim, error) {
	var claim JWTClaim

	token, err := jwt.ParseWithClaims(accessToken, &claim, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwa.config.Method {
			return nil, ErrInvalidSigningMethod
		}

		if jwa.config.Authenticator != nil {
			return jwa.config.Authenticator(jwa.config, t)
		}

		return jwa.config.Method, nil
	})

	if err != nil {
		if tokenErr, ok := err.(*jwt.ValidationError); ok {
			if tokenErr.Errors&jwt.ValidationErrorMalformed != 0 {
				return pkg.UserClaim{}, JWTError{
					SrcErr: tokenErr,
					Err:    ErrInvalidJWTToken,
				}
			}

			// Token is either expired or not active yet
			if tokenErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return pkg.UserClaim{}, JWTError{
					SrcErr: tokenErr,
					Err:    ErrExpiredJWTToken,
				}
			}
		}

		return pkg.UserClaim{}, JWTError{
			SrcErr: err,
			Err:    ErrParseJWTToken,
		}
	}

	if !token.Valid {
		return pkg.UserClaim{}, JWTError{
			Err:    ErrParseJWTToken,
			SrcErr: errors.New("jwt token is not valid"),
		}
	}

	jwtClaim, ok := token.Claims.(JWTClaim)
	if !ok {
		return pkg.UserClaim{}, JWTError{
			Err:    ErrUnexpectedJWTClaim,
			SrcErr: errors.New("invalid jwt.Claim type"),
		}
	}

	return jwtClaim.Data, nil
}

// Revoke exists for the purpose to actively revoke a an existing jwt record refresh token
// which then revokes all present and pending
func (jwa JWTIdentity) Revoke(ctx context.Context, refreshToken string) error {
	// Retrieve jwt record from backend.
	record, err := jwa.IdentityDBBackend.GetByField(ctx, "refresh_token", refreshToken)
	if err != nil {
		return JWTError{
			SrcErr: err,
			Err:    ErrInvalidRefreshToken,
		}
	}

	// Add refresh token into in-memory
	if _, err = jwa.blacklist.Add(ctx, record.PublicID, record.RefreshToken); err != nil {
		return JWTError{
			SrcErr: err,
			Err:    ErrInternalError,
		}
	}

	if err = jwa.IdentityDBBackend.Delete(ctx, record.PublicID); err != nil {
		return JWTError{
			SrcErr: err,
			Err:    ErrInternalError,
		}
	}

	return nil
}

// Refresh attempts to provide a new access token through the provided refresh token if valid.
// Allow the user to get new token for access to underline resources.
func (jwa JWTIdentity) Refresh(ctx context.Context, refreshToken string) (JWTAuth, error) {
	// Retrieve jwt record from backend.
	record, err := jwa.IdentityDBBackend.GetByField(ctx, "refresh_token", refreshToken)
	if err != nil {
		return JWTAuth{}, JWTError{
			SrcErr: err,
			Err:    ErrInvalidRefreshToken,
		}
	}

	// check that records refresh token is not in blacklist.
	found, err := jwa.blacklist.Has(ctx, record.PublicID, record.RefreshToken)
	if err != nil {
		return JWTAuth{}, JWTError{
			SrcErr: err,
			Err:    ErrInternalError,
		}
	}

	// If in blacklist, then return denied error.
	if found {
		return JWTAuth{}, JWTError{
			SrcErr: errors.New("refresh denied"),
			Err:    ErrTokenRefreshDenied,
		}
	}

	now := time.Now()
	if record.Expires.Before(now) {
		if err := jwa.IdentityDBBackend.Delete(ctx, record.PublicID); err != nil {
			return JWTAuth{}, JWTError{
				Err:    ErrInternalError,
				SrcErr: err,
			}
		}

		// Black list refresh token.
		if _, err := jwa.blacklist.Add(ctx, record.PublicID, record.RefreshToken); err != nil {
			return JWTAuth{}, JWTError{
				Err:    ErrInternalError,
				SrcErr: err,
			}
		}

		return JWTAuth{}, JWTError{
			SrcErr: errors.New("refresh has expired, relogin"),
			Err:    ErrExpiredRefreshToken,
		}
	}

	accessExpires := now.Add(jwa.config.AccessTokenExpires)

	var claim JWTClaim
	claim.Data = record.Data
	claim.Id = record.PublicID
	claim.IssuedAt = now.Unix()
	claim.Subject = record.TargetID
	claim.ExpiresAt = accessExpires.Unix()

	jwtToken := jwt.NewWithClaims(jwa.config.Method, claim)
	jwtAccessToken, err := jwtToken.SignedString(jwa.config.SigningSecret)
	if err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	return JWTAuth{
		RefreshToken:   record.RefreshToken,
		AccessToken:    jwtAccessToken,
		Expires:        claim.ExpiresAt,
		RefreshExpires: record.Expires.Unix(),
		TokenType:      "Bearer",
	}, nil
}

// Grant generates a new jwt token for provided credential contract if valiated to be correct and authorizable.
// It returns a json of JWTClaim.
func (jwa JWTIdentity) Grant(ctx context.Context, cr pkg.CreateUserSession) (JWTAuth, error) {
	rclaim, err := jwa.config.Maker(ctx, jwa.config, cr)
	if err != nil {
		return JWTAuth{}, err
	}

	now := time.Now()
	accessExpires := now.Add(jwa.config.AccessTokenExpires)
	refreshExpires := now.Add(jwa.config.RefreshTokenExpires)

	var record Identity
	record.IssuedAt = now
	record.LastLogin = now
	record.Data = rclaim.Data
	record.Expires = refreshExpires
	record.TargetID = rclaim.TargetID
	record.PublicID = uuid.NewV4().String()
	record.RefreshToken = uuid.NewV4().String()
	record.RefreshInterval = jwa.config.AccessTokenExpires

	if err := record.Validate(); err != nil {
		return JWTAuth{}, JWTError{
			SrcErr: err,
			Err:    ErrInvalidIdentity,
		}
	}

	var claim JWTClaim
	claim.Data = rclaim.Data
	claim.Id = record.PublicID
	claim.IssuedAt = now.Unix()
	claim.Subject = record.TargetID
	claim.ExpiresAt = accessExpires.Unix()

	// If a hold an existing record, then destroy it, we get two benefits:
	// 1. Ensure users refresh_token becomes absolutely invalidated and un-usable.
	// 2. Reduce risk of usage of old access token by attacker with nefarious desires as
	// token stops working.
	if beforeRecord, err := jwa.IdentityDBBackend.GetByField(ctx, "target_id", rclaim.TargetID); err == nil {
		if err = jwa.IdentityDBBackend.Delete(ctx, beforeRecord.PublicID); err != nil {
			return JWTAuth{}, JWTError{
				Err:    ErrInternalError,
				SrcErr: err,
			}
		}

		// Black list refresh token.
		if _, err = jwa.blacklist.Add(ctx, beforeRecord.PublicID, beforeRecord.RefreshToken); err != nil {
			return JWTAuth{}, JWTError{
				Err:    ErrInternalError,
				SrcErr: err,
			}
		}
	}

	// Save new user record and send out refresh token and access token.
	if err := jwa.IdentityDBBackend.Create(ctx, record); err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	jwtToken := jwt.NewWithClaims(jwa.config.Method, claim)
	jwtAccessToken, err := jwtToken.SignedString(jwa.config.SigningSecret)
	if err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	return JWTAuth{
		TokenType:      "Bearer",
		RefreshToken:   record.RefreshToken,
		AccessToken:    jwtAccessToken,
		Expires:        claim.ExpiresAt,
		RefreshExpires: refreshExpires.Unix(),
	}, nil
}
