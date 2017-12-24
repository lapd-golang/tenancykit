// Package userclaimjwt provides a auto-generated package which contains a API for authentication using JWT.
//
//
package userclaimjwt

import (
	"encoding/base64"
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
	ErrNotFound                = errors.New("not found")
	ErrInvalidIdentity         = errors.New("provided Identity is invalid")
	ErrParseJWTToken           = errors.New("parse error: invalid jwt token")
	ErrNoJWTAuthorizationToken = errors.New("no jwt authorization token")
	ErrInternalError           = errors.New("backend failed with error")
	ErrInvalidJWTToken         = errors.New("received jwt token is invalid")
	ErrUnexpectedJWTClaim      = errors.New("jwt token claim is not expected type")
	ErrExpiredJWTToken         = errors.New("received jwt token is expired")
	ErrInvalidRefreshToken     = errors.New("Invalid refresh token")
	ErrTokenRefreshDenied      = errors.New("refresh_token already blacklist")
	ErrExpiredRefreshToken     = errors.New("refresh_token already expired")
	ErrNoTargetHeaderInToken   = errors.New("token.Header has no 'jwt-target-id'")
	ErrExpiredAccessToken      = errors.New("access_token already expired")
	ErrInvalidSigningMethod    = errors.New("token signing method mismatched")
	ErrFailedToGetSecret       = errors.New("target-id failed to get secret from secrets function")
)

// Identity embodies data stored for a user's login credentials.
type Identity struct {
	PublicID        string        `json:"public_id"`
	RefreshToken    string        `json:"refresh_token"`
	Expires         int64         `json:"expires"`
	TargetID        string        `json:"target_id"`
	LastLogin       int64         `json:"last_login"`
	IssuedAt        int64         `json:"last_login"`
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

// Validate returns an error if giving JWTAuth does not match desired
// field value state.
func (ja JWTAuth) Validate() error {
	if strings.TrimSpace(ja.TokenType) == "" {
		return errors.New("TokenType is required")
	}

	if strings.TrimSpace(ja.AccessToken) == "" {
		return errors.New("AccessToken is required")
	}

	if strings.TrimSpace(ja.RefreshToken) == "" {
		return errors.New("RefreshToken is required")
	}

	if ja.Expires <= 0 {
		return errors.New("expiration time in unix is required")
	}

	if ja.RefreshExpires <= 0 {
		return errors.New("refresh expiration time in unix is required")
	}
	return nil
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

// IdentityDBBackend defines an interface which exposes a underline storage system
// for retrieving and storing identity records.
// @implement
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

// IdentityOps embodies method specific for grant, authenticating, revoking and refreshing identities.
type IdentityOps interface {
	Revoke(context.Context, string) error
	Refresh(context.Context, string) (JWTAuth, error)
	Grant(context.Context, pkg.CreateUserSession) (JWTAuth, error)
	Authenticate(context.Context, string) (pkg.UserClaim, error)
}

// IdentityBackend defines an interface that expose a backend interface which can
// expose methods that contain all necessary logic for interaction with api for http endpoints.
type IdentityBackend interface {
	IdentityOps
	Count(context.Context) (int, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (Identity, error)
	Update(context.Context, string, Identity) error
	GetAll(context.Context, string, string, int, int) ([]Identity, int, error)
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
	AuthenticateJSON(*httputil.Context) error
}

// IdentityToken embodies data received over api calls to refresh or revoke a identity token.
type IdentityToken struct {
	RefreshToken string `json:"refresh_token"`
}

// IdentityAccess embodies data received over api calls to revoke or refresh a identity token.
type IdentityAccess struct {
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
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
	SpecialID string `json:"special_id"`
	Data      pkg.UserClaim
}

// IdentityMaker defines a function type provided by maker for generating identity claim.
type IdentityMaker func(context.Context, JWTConfig, pkg.CreateUserSession) (Testimony, error)

// TokenValidator defines a function type provided by user for custom validating incoming token.
type TokenValidator func(context.Context, JWTConfig, *jwt.Token) error

// TokenSecrets defines a function type when provided the config and target id received from
// the Identity maker will return a byte slice which represent the secret the token is signed with.
type TokenSecrets func(context.Context, JWTConfig, string) ([]byte, error)

// JWTConfig embodies the field for configuring JWTBackend.
type JWTConfig struct {
	Signer              string
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
	Maker               IdentityMaker
	Secrets             TokenSecrets
	Validator           TokenValidator
	Method              jwt.SigningMethod
}

// JWTIdentity implements the IdentityBackend interface and embodies all
// necessary method for granting, revoking and refreshing jwt access and refresh tokens.
type JWTIdentity struct {
	config    JWTConfig
	blacklist tokens.TokenSet
	whitelist tokens.TokenSet
	IdentityDBBackend
}

// NewJWTIdentity returns a new JWTIdentity instance which embodies and implements the IdentityBackend
// interface.
func NewJWTIdentity(config JWTConfig, whitelist tokens.TokenSet, blacklist tokens.TokenSet, backend IdentityDBBackend) JWTIdentity {
	return JWTIdentity{
		config:            config,
		blacklist:         blacklist,
		whitelist:         whitelist,
		IdentityDBBackend: backend,
	}
}

// Authenticate attempts to authenticate users access token to validate user's
func (jwa JWTIdentity) Authenticate(ctx context.Context, accessToken string) (pkg.UserClaim, error) {
	var claim JWTClaim

	token, err := jwt.ParseWithClaims(accessToken, &claim, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwa.config.Method {
			return nil, ErrInvalidSigningMethod
		}

		// If user provides validator then run to allow user validate token.
		if jwa.config.Validator != nil {
			if err := jwa.config.Validator(ctx, jwa.config, t); err != nil {
				return nil, err
			}
		}

		targetID, ok := t.Header["jwt-target-id"].(string)
		if !ok {
			return nil, ErrNoTargetHeaderInToken
		}

		// Retrieve secret for target from secrets function.
		secret, err := jwa.config.Secrets(ctx, jwa.config, targetID)
		if err != nil {
			return nil, err
		}

		return secret, nil
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

	// check in-memory cache.
	has, err := jwa.whitelist.Has(ctx, claim.Id, claim.SpecialID)
	if err != nil {
		return pkg.UserClaim{}, JWTError{
			Err:    ErrParseJWTToken,
			SrcErr: errors.New("jwt token is not valid"),
		}
	}

	if !has {
		return pkg.UserClaim{}, JWTError{
			Err:    ErrExpiredAccessToken,
			SrcErr: errors.New("jwt token is not valid"),
		}
	}

	return claim.Data, nil
}

// Revoke exists for the purpose to actively revoke a an existing jwt record refresh token
// which then revokes all present and pending
func (jwa JWTIdentity) Revoke(ctx context.Context, encodedRefreshToken string) error {
	unbased64Token, err := base64.StdEncoding.DecodeString(encodedRefreshToken)
	if err != nil {
		return JWTError{
			SrcErr: err,
			Err:    ErrInvalidRefreshToken,
		}
	}

	decodedRefToken := strings.Split(string(unbased64Token), ":")
	if len(decodedRefToken) != 2 {
		return JWTError{
			Err:    ErrInvalidRefreshToken,
			SrcErr: errors.New("refresh_token does not match wanted format"),
		}
	}

	refreshToken, specialID := decodedRefToken[0], decodedRefToken[1]

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

	// remove record target and public key to current active lists.
	if err = jwa.whitelist.Remove(ctx, record.PublicID, specialID); err != nil && err != tokens.ErrNotFound {
		return JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	return nil
}

// Refresh attempts to provide a new access token through the provided refresh token if valid.
// Allow the user to get new token for access to underline resources.
func (jwa JWTIdentity) Refresh(ctx context.Context, encodedRefreshToken string) (JWTAuth, error) {
	unbased64Token, err := base64.StdEncoding.DecodeString(encodedRefreshToken)
	if err != nil {
		return JWTAuth{}, JWTError{
			SrcErr: err,
			Err:    ErrInvalidRefreshToken,
		}
	}

	decodedRefToken := strings.Split(string(unbased64Token), ":")
	if len(decodedRefToken) != 2 {
		return JWTAuth{}, JWTError{
			Err:    ErrInvalidRefreshToken,
			SrcErr: errors.New("refresh_token does not match wanted format"),
		}
	}

	refreshToken, specialID := decodedRefToken[0], decodedRefToken[1]

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
	expires := time.Unix(record.Expires, 0)
	if expires.Before(now) {
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

		// remove record target and public key to current active lists.
		if err = jwa.whitelist.Remove(ctx, record.PublicID, specialID); err != nil && err != tokens.ErrNotFound {
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

	// Retrieve secret for target from secrets function.
	secret, err := jwa.config.Secrets(ctx, jwa.config, record.TargetID)
	if err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrFailedToGetSecret,
			SrcErr: err,
		}
	}

	// Delete old token existing for record for access.
	if err = jwa.whitelist.Remove(ctx, record.PublicID, specialID); err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	accessExpires := now.Add(jwa.config.AccessTokenExpires)

	var claim JWTClaim
	claim.Data = record.Data
	claim.Id = record.PublicID
	claim.IssuedAt = now.Unix()
	claim.Subject = record.TargetID
	claim.ExpiresAt = accessExpires.Unix()
	claim.SpecialID = uuid.NewV4().String()

	// add record target and public key to current active lists.
	if _, err = jwa.whitelist.Add(ctx, record.PublicID, claim.SpecialID); err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	jwtToken := jwt.NewWithClaims(jwa.config.Method, claim)
	jwtToken.Header["jwt-target-id"] = record.TargetID
	jwtAccessToken, err := jwtToken.SignedString(secret)
	if err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	// create new refresh token with special id and record refresh token.
	drefreshToken := base64.StdEncoding.EncodeToString([]byte(record.RefreshToken + ":" + claim.SpecialID))

	return JWTAuth{
		RefreshToken:   drefreshToken,
		AccessToken:    jwtAccessToken,
		Expires:        claim.ExpiresAt,
		RefreshExpires: record.Expires,
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

	// Retrieve secret for target from secrets function.
	secret, err := jwa.config.Secrets(ctx, jwa.config, rclaim.TargetID)
	if err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrFailedToGetSecret,
			SrcErr: err,
		}
	}

	now := time.Now()
	accessExpires := now.Add(jwa.config.AccessTokenExpires)
	refreshExpires := now.Add(jwa.config.RefreshTokenExpires)

	var record Identity
	record.IssuedAt = now.Unix()
	record.LastLogin = now.Unix()
	record.Data = rclaim.Data
	record.Expires = refreshExpires.Unix()
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
	claim.SpecialID = uuid.NewV4().String()

	// Save new user record and send out refresh token and access token.
	if err := jwa.IdentityDBBackend.Create(ctx, record); err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	// add record target and public key to current active lists.
	if _, err = jwa.whitelist.Add(ctx, record.PublicID, claim.SpecialID); err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	jwtToken := jwt.NewWithClaims(jwa.config.Method, claim)
	jwtToken.Header["jwt-target-id"] = record.TargetID
	jwtAccessToken, err := jwtToken.SignedString(secret)
	if err != nil {
		return JWTAuth{}, JWTError{
			Err:    ErrInternalError,
			SrcErr: err,
		}
	}

	// create new refresh token with special id and record refresh token.
	refreshToken := base64.StdEncoding.EncodeToString([]byte(record.RefreshToken + ":" + claim.SpecialID))

	return JWTAuth{
		TokenType:      "Bearer",
		RefreshToken:   refreshToken,
		AccessToken:    jwtAccessToken,
		Expires:        claim.ExpiresAt,
		RefreshExpires: refreshExpires.Unix(),
	}, nil
}
