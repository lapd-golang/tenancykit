package backends

import (
	"net/http"
	"strings"
	"time"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/tokenmgo"
	"github.com/gokit/tenancykit/pkg/db/tokensql"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/influx6/faux/context"
	"github.com/influx6/faux/httputil"
)

// TenantBackend is a wrapper to implement tenantapi.Backend interface methods for
// types.TenantBackend.
type TenantBackend struct {
	types.TenantDBBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (t TenantBackend) Create(ctx context.Context, nt pkg.CreateTenant) (pkg.Tenant, error) {
	newTenant := pkg.NewTenant(nt)
	return newTenant, t.TenantDBBackend.Create(ctx, newTenant)
}

// TwoFactorSessionBackend implements the api.twofactorsessionapi.Backend interface and wraps a types.TFRecordBackend.
type TwoFactorSessionBackend struct {
	types.TwoFactorSessionDBBackend
}

// Create attempts to create a new TFRecord in the db using the provided pkg.NewTF struct.
func (tf TwoFactorSessionBackend) Create(ctx context.Context, ntf pkg.TwoFactorSession) (pkg.TwoFactorSession, error) {
	return ntf, tf.TwoFactorSessionDBBackend.Create(ctx, ntf)
}

// TFBackend implements the api.tfrecordapi.Backend interface and wraps a types.TFRecordBackend.
type TFBackend struct {
	types.TFRecordDBBackend
}

// Create attempts to create a new TFRecord in the db using the provided pkg.NewTF struct.
func (tf TFBackend) Create(ctx context.Context, ntf pkg.NewTF) (pkg.TFRecord, error) {
	newTF, err := pkg.NewTFRecord(ntf.MaxLength, ntf.Domain, ntf.Tenant, ntf.User)
	if err != nil {
		return newTF, err
	}

	return newTF, tf.TFRecordDBBackend.Create(ctx, newTF)
}

// UserSessionBackend wraps types.UserSessionBackend to implement the api.usersessionapi.Backend
// create method.
type UserSessionBackend struct {
	types.UserSessionDBBackend
	UserBackend types.UserDBBackend
}

// Create attempts to create a user session for the incoming pkg.CreateUserSession
// and returns a new pkg.UserSession if successfully as returning error.
func (us UserSessionBackend) Create(ctx context.Context, cus pkg.CreateUserSession) (pkg.UserSession, error) {
	if err := cus.Validate(); err != nil {
		return pkg.UserSession{}, err
	}

	// If api is not provided a expiration time for session, then set to 24hrs.
	if cus.Expiration <= 0 {
		cus.Expiration = 24 * time.Hour
	}

	// Attempt to retrieve user from user backend.
	user, err := us.UserBackend.GetByField(ctx, "email", cus.Email)
	if err != nil {
		return pkg.UserSession{}, httputil.HTTPError{
			Err:  err,
			Code: http.StatusNotFound,
		}
	}

	// Attempt to validate user password.
	if err := user.Authenticate(cus.Password); err != nil {
		return pkg.UserSession{}, httputil.HTTPError{
			Err:  err,
			Code: http.StatusUnauthorized,
		}
	}

	ctx.Bag().Set(pkg.ContextKeyUser, user)

	// Check if old session exist and if not expired, then return that, else delete it.
	if oldSession, err := us.UserSessionDBBackend.GetByField(ctx, "user_id", user.PublicID); err == nil {
		if !oldSession.Expired() {
			return oldSession, nil
		}

		if err := us.UserSessionDBBackend.Delete(ctx, oldSession.PublicID); err != nil {
			return pkg.UserSession{}, err
		}
	}

	newSession := pkg.NewUserSession(user.PublicID, time.Now().Add(cus.Expiration))
	if err := us.UserSessionDBBackend.Create(ctx, newSession); err != nil {
		return pkg.UserSession{}, err
	}

	return newSession, nil
}

// UserBackend is a wrapper to implement userapi.Backend interface methods for
// types.UserBackend.
type UserBackend struct {
	types.UserDBBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (u UserBackend) Create(ctx context.Context, nt pkg.CreateUser) (pkg.User, error) {
	if err := nt.Validate(); err != nil {
		return pkg.User{}, err
	}

	newUser, err := pkg.NewUser(nt)
	if err != nil {
		return newUser, err
	}

	return newUser, u.UserDBBackend.Create(ctx, newUser)
}

// Update attempts to update user password retrieved from underline types.UserBackend.
// It implements api.userapi.Backend.Update interface method.
func (u UserBackend) Update(ctx context.Context, id string, updater pkg.UpdateUser) error {
	if err := updater.Validate(); err != nil {
		return err
	}

	user, err := u.UserDBBackend.Get(ctx, id)
	if err != nil {
		return err
	}

	if !updater.IsPasswordUpdate && strings.TrimSpace(updater.Email) != "" {
		user.Email = updater.Email
		return u.UserDBBackend.Update(ctx, id, user)
	}

	if err = user.ChangePassword(updater.Password); err != nil {
		return err
	}

	return u.UserDBBackend.Update(ctx, id, user)
}

// TokenBackend implements tokenapi.TokenBackend.
type TokenBackend struct {
	types.TokenDBBackend
}

// Create attempts to save provided token into db and returning token once done.
// It implements tokenapi.TokenBackend interface.
func (tf TokenBackend) Create(ctx context.Context, elem pkg.Token) (pkg.Token, error) {
	return elem, tf.TokenDBBackend.Create(ctx, elem)
}

// Has returns true/false if giving token exists or not within underline
// db. Returns an error if call to db failed.
func (tf TokenBackend) Has(ctx context.Context, targetID string, token string) (bool, error) {
	_, err := tf.TokenDBBackend.GetByField(ctx, "target_id", token)
	if err != nil && (err == tokensql.ErrNotFound || err == tokenmgo.ErrNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// Add adds giving underline tokendb into db, returning error if
// it fails to do so, or call to db errors out.
func (tf TokenBackend) Add(ctx context.Context, targetID string, token string) (pkg.Token, error) {
	newtoken := pkg.NewToken(targetID, token)
	return tf.Create(ctx, newtoken)
}
