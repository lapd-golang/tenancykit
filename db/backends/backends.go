package backends

import (
	"net/http"
	"strings"
	"time"

	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/db/types"
	"github.com/influx6/faux/context"
	"github.com/influx6/faux/httputil"
)

// TenantBackend is a wrapper to implement tenantapi.Backend interface methods for
// types.TenantBackend.
type TenantBackend struct {
	types.TenantBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (t TenantBackend) Create(ctx context.Context, nt tenancykit.CreateTenant) (tenancykit.Tenant, error) {
	newTenant := tenancykit.NewTenant(nt)
	return newTenant, t.TenantBackend.Create(ctx, newTenant)
}

// TwoFactorSessionBackend implements the api.twofactorsessionapi.Backend interface and wraps a types.TFRecordBackend.
type TwoFactorSessionBackend struct {
	types.TwoFactorSessionBackend
}

// Create attempts to create a new TFRecord in the db using the provided tenancykit.NewTF struct.
func (tf TwoFactorSessionBackend) Create(ctx context.Context, ntf tenancykit.TwoFactorSession) (tenancykit.TwoFactorSession, error) {
	return ntf, tf.TwoFactorSessionBackend.Create(ctx, ntf)
}

// TFBackend implements the api.tfrecordapi.Backend interface and wraps a types.TFRecordBackend.
type TFBackend struct {
	types.TFRecordBackend
}

// Create attempts to create a new TFRecord in the db using the provided tenancykit.NewTF struct.
func (tf TFBackend) Create(ctx context.Context, ntf tenancykit.NewTF) (tenancykit.TFRecord, error) {
	newTF, err := tenancykit.NewTFRecord(ntf.MaxLength, ntf.Domain, ntf.Tenant, ntf.User)
	if err != nil {
		return newTF, err
	}

	return newTF, tf.TFRecordBackend.Create(ctx, newTF)
}

// UserSessionBackend wraps types.UserSessionBackend to implement the api.usersessionapi.Backend
// create method.
type UserSessionBackend struct {
	types.UserSessionBackend
	UserBackend types.UserBackend
}

// Create attempts to create a user session for the incoming tenancykit.CreateUserSession
// and returns a new tenancykit.UserSession if successfully as returning error.
func (us UserSessionBackend) Create(ctx context.Context, cus tenancykit.CreateUserSession) (tenancykit.UserSession, error) {
	if err := cus.Validate(); err != nil {
		return tenancykit.UserSession{}, err
	}

	// If api is not provided a expiration time for session, then set to 24hrs.
	if cus.Expiration <= 0 {
		cus.Expiration = 24 * time.Hour
	}

	// Attempt to retrieve user from user backend.
	user, err := us.UserBackend.GetByField(ctx, "email", cus.Email)
	if err != nil {
		return tenancykit.UserSession{}, httputil.HTTPError{
			Err:  err,
			Code: http.StatusNotFound,
		}
	}

	// Attempt to validate user password.
	if err := user.Authenticate(cus.Password); err != nil {
		return tenancykit.UserSession{}, httputil.HTTPError{
			Err:  err,
			Code: http.StatusUnauthorized,
		}
	}

	ctx.Bag().Set(tenancykit.ContextKeyUser, user)

	// Check if old session exist and if not expired, then return that, else delete it.
	if oldSession, err := us.UserSessionBackend.GetByField(ctx, "user_id", user.PublicID); err == nil {
		if !oldSession.Expired() {
			return oldSession, nil
		}

		if err := us.UserSessionBackend.Delete(ctx, oldSession.PublicID); err != nil {
			return tenancykit.UserSession{}, err
		}
	}

	newSession := tenancykit.NewUserSession(user.PublicID, time.Now().Add(cus.Expiration))
	if err := us.UserSessionBackend.Create(ctx, newSession); err != nil {
		return tenancykit.UserSession{}, err
	}

	return newSession, nil
}

// UserBackend is a wrapper to implement userapi.Backend interface methods for
// types.UserBackend.
type UserBackend struct {
	types.UserBackend
}

// Create modifies the underline types.TenantBackend.Create method to implement
// the api.tenantapi.Backend interface.
func (u UserBackend) Create(ctx context.Context, nt tenancykit.CreateUser) (tenancykit.User, error) {
	if err := nt.Validate(); err != nil {
		return tenancykit.User{}, err
	}

	newUser, err := tenancykit.NewUser(nt)
	if err != nil {
		return newUser, err
	}

	return newUser, u.UserBackend.Create(ctx, newUser)
}

// Update attempts to update user password retrieved from underline types.UserBackend.
// It implements api.userapi.Backend.Update interface method.
func (u UserBackend) Update(ctx context.Context, id string, updater tenancykit.UpdateUser) error {
	if err := updater.Validate(); err != nil {
		return err
	}

	user, err := u.UserBackend.Get(ctx, id)
	if err != nil {
		return err
	}

	if !updater.IsPasswordUpdate && strings.TrimSpace(updater.Email) != "" {
		user.Email = updater.Email
		return u.UserBackend.Update(ctx, id, user)
	}

	if err = user.ChangePassword(updater.Password); err != nil {
		return err
	}

	return u.UserBackend.Update(ctx, id, user)
}
