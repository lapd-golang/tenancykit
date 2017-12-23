package backends

import (
	"context"
	"net/http"
	"time"

	"github.com/gokit/tenancykit/pkg"
	"github.com/gokit/tenancykit/pkg/db/types"
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
	if err := nt.Validate(); err != nil {
		return pkg.Tenant{}, err
	}

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

	user.Email = updater.Email
	if !updater.IsPassWordUpdate() {
		return u.UserDBBackend.Update(ctx, id, user)
	}

	if err = user.ChangePassword(updater.Password); err != nil {
		return err
	}

	return u.UserDBBackend.Update(ctx, id, user)
}
