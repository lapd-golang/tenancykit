package userclaimjwt_test

import (
	"encoding/json"
	"fmt"

	"testing"

	"time"

	"context"

	"github.com/influx6/faux/tests"

	jwt "github.com/dgrijalva/jwt-go"

	userclaimjwt "github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"

	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt/mock"

	"github.com/gokit/tenancykit/pkg"
)

var (
	claimDataJSON = `{


    "user":	null,

    "tenant":	null

}`
	contractDataJSON = `{


    "email":	"JeremyLawson@Babbleopia.mil",

    "password":	"ku2ikso0bji55bjiilb6",

    "expiration":	null

}`
	userClaim = userclaimjwt.Testimony{
		TargetID: "7fd15938c823cf58e78019dfddf2af142f9449696a",
	}
)

func noSecureUser(ctx context.Context, config userclaimjwt.JWTConfig, cr pkg.CreateUserSession) (userclaimjwt.Testimony, error) {
	var item pkg.UserClaim
	if err := json.Unmarshal([]byte(claimDataJSON), &item); err != nil {
		return userClaim, err
	}
	userClaim.Data = item
	return userClaim, nil
}

func secretFunc(ctx context.Context, config userclaimjwt.JWTConfig, targetID string) ([]byte, error) {
	return []byte("All we want is to sign this"), nil
}

func TestUserClaimJWT(t *testing.T) {
	blacklist := mock.TokenBackend()
	whitelist := mock.TokenBackend()
	idb := mock.IdentityBackend()
	jwter := userclaimjwt.NewJWTIdentity(userclaimjwt.JWTConfig{
		Maker:               noSecureUser,
		Signer:              "wellington",
		Secrets:             secretFunc,
		Method:              jwt.SigningMethodHS256,
		AccessTokenExpires:  600 * time.Millisecond,
		RefreshTokenExpires: 1 * time.Second,
	}, blacklist, whitelist, idb)

	var cred pkg.CreateUserSession
	if jsnerr := json.Unmarshal([]byte(contractDataJSON), &cred); jsnerr != nil {
		tests.FailedWithError(jsnerr, "Should have successfully parsed pkg.CreateUserSession json")
	}
	tests.Passed("Should have successfully parsed pkg.CreateUserSession json")

	auth, err := jwter.Grant(context.Background(), cred)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully authenticated credentials with auth response")
	}
	tests.Passed("Should have successfully authenticated credentials with auth response")

	if err := auth.Validate(); err != nil {
		tests.FailedWithError(err, "Should have received valid auth token")
	}
	tests.Passed("Should have received valid auth token")

	claim, err := jwter.Authenticate(context.Background(), auth.AccessToken)
	if err != nil {
		tests.FailedWithError(err, "Should have succesfully authenticate access token")
	}
	tests.Passed("Should have succesfully authenticate access token")

	if claim.User.PublicID != userClaim.Data.User.PublicID {
		tests.Failed("Should have received same user claim data")
	}
	tests.Passed("Should have received same user claim data")

	newAuth, err := jwter.Refresh(context.Background(), auth.RefreshToken)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully refreshed token")
	}
	tests.Passed("Should have successfully refreshed token")

	if err := newAuth.Validate(); err != nil {
		tests.FailedWithError(err, "Should have received valid auth token")
	}
	tests.Passed("Should have received valid auth token")

	if _, err := jwter.Authenticate(context.Background(), auth.AccessToken); err == nil {
		tests.FailedWithError(err, "Should have failed to authenticate access token after refresh")
	}
	tests.Passed("Should have failed to authenticate access token after refresh")

	claim, err = jwter.Authenticate(context.Background(), newAuth.AccessToken)
	if err != nil {
		tests.FailedWithError(err, "Should have succesfully authenticate access token")
	}
	tests.Passed("Should have succesfully authenticate access token")

	if claim.User.PublicID != userClaim.Data.User.PublicID {
		tests.Failed("Should have received same user claim data")
	}
	tests.Passed("Should have received same user claim data")

	if err := jwter.Revoke(context.Background(), newAuth.RefreshToken); err != nil {
		tests.FailedWithError(err, "Should have successfully revoked token")
	}
	tests.Passed("Should have successfully revoked token")

	if _, err = jwter.Authenticate(context.Background(), newAuth.AccessToken); err == nil {
		tests.Failed("Should have failed to authenticate with access token from revoked user")
	}
	tests.Passed("Should have failed to authenticate with access token from revoked user")
}

func TestUserClaimJWTExpiration(t *testing.T) {
	blacklist := mock.TokenBackend()
	whitelist := mock.TokenBackend()
	idb := mock.IdentityBackend()
	jwter := userclaimjwt.NewJWTIdentity(userclaimjwt.JWTConfig{
		Maker:               noSecureUser,
		Signer:              "wellington",
		Secrets:             secretFunc,
		Method:              jwt.SigningMethodHS256,
		AccessTokenExpires:  100 * time.Millisecond,
		RefreshTokenExpires: 300 * time.Millisecond,
	}, blacklist, whitelist, idb)

	var cred pkg.CreateUserSession
	if jsnerr := json.Unmarshal([]byte(contractDataJSON), &cred); jsnerr != nil {
		tests.FailedWithError(jsnerr, "Should have successfully parsed pkg.CreateUserSession json")
	}
	tests.Passed("Should have successfully parsed pkg.CreateUserSession json")

	auth, err := jwter.Grant(context.Background(), cred)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully authenticated credentials with auth response")
	}
	tests.Passed("Should have successfully authenticated credentials with auth response")

	if err := auth.Validate(); err != nil {
		tests.FailedWithError(err, "Should have received valid auth token")
	}
	tests.Passed("Should have received valid auth token")

	claim, err := jwter.Authenticate(context.Background(), auth.AccessToken)
	if err != nil {
		tests.FailedWithError(err, "Should have succesfully authenticate access token")
	}
	tests.Passed("Should have succesfully authenticate access token")

	if claim.User.PublicID != userClaim.Data.User.PublicID {
		tests.Failed("Should have received same user claim data")
	}
	tests.Passed("Should have received same user claim data")

	time.Sleep(2 * time.Second)

	_, err = jwter.Authenticate(context.Background(), auth.AccessToken)
	fmt.Printf("AuthErr: %+q\n", err)
	if err == nil {
		tests.Failed("Should have failed to authenticate with access token from revoked user")
	}
	tests.Passed("Should have failed to authenticate with access token from revoked user")
}

func TestUserClaimJWTUserFlow(t *testing.T) {
	blacklist := mock.TokenBackend()
	whitelist := mock.TokenBackend()
	idb := mock.IdentityBackend()
	jwter := userclaimjwt.NewJWTIdentity(userclaimjwt.JWTConfig{
		Maker:               noSecureUser,
		Signer:              "wellington",
		Secrets:             secretFunc,
		Method:              jwt.SigningMethodHS256,
		AccessTokenExpires:  700 * time.Millisecond,
		RefreshTokenExpires: 4 * time.Second,
	}, blacklist, whitelist, idb)

	var cred pkg.CreateUserSession
	if jsnerr := json.Unmarshal([]byte(contractDataJSON), &cred); jsnerr != nil {
		tests.FailedWithError(jsnerr, "Should have successfully parsed pkg.CreateUserSession json")
	}
	tests.Passed("Should have successfully parsed pkg.CreateUserSession json")

	auth, err := jwter.Grant(context.Background(), cred)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully authenticated credentials with auth response")
	}
	tests.Passed("Should have successfully authenticated credentials with auth response")

	if err := auth.Validate(); err != nil {
		tests.FailedWithError(err, "Should have received valid auth token")
	}
	tests.Passed("Should have received valid auth token")

	claim, err := jwter.Authenticate(context.Background(), auth.AccessToken)
	if err != nil {
		tests.FailedWithError(err, "Should have succesfully authenticate access token")
	}
	tests.Passed("Should have succesfully authenticate access token")

	if claim.User.PublicID != userClaim.Data.User.PublicID {
		tests.Failed("Should have received same user claim data")
	}
	tests.Passed("Should have received same user claim data")

	time.Sleep(2 * time.Second)

	if _, err = jwter.Authenticate(context.Background(), auth.AccessToken); err == nil {
		tests.Failed("Should have failed to authenticate with access token after expiration")
	}
	tests.Passed("Should have failed to authenticate with access token after expiration")

	newAuth, err := jwter.Refresh(context.Background(), auth.RefreshToken)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully refreshed token")
	}
	tests.Passed("Should have successfully refreshed token")

	if err := newAuth.Validate(); err != nil {
		tests.FailedWithError(err, "Should have received valid auth token")
	}
	tests.Passed("Should have received valid auth token")

	if _, err := jwter.Authenticate(context.Background(), newAuth.AccessToken); err != nil {
		tests.FailedWithError(err, "Should have succeeded in authenticating access token after refresh")
	}
	tests.Passed("Should have succeeded in authenticating access token after refresh")

	time.Sleep(7 * time.Second)

	if _, err = jwter.Authenticate(context.Background(), auth.AccessToken); err == nil {
		tests.Failed("Should have failed to authenticate with access token from revoked user")
	}
	tests.Passed("Should have failed to authenticate with access token from revoked user")

	if _, err = jwter.Authenticate(context.Background(), newAuth.AccessToken); err == nil {
		tests.Failed("Should have failed to authenticate with access token from revoked user")
	}
	tests.Passed("Should have failed to authenticate with access token from revoked user")
}
