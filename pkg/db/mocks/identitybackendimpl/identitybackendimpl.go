package identitybackendimpl

import (
	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"

	context "context"

	pkg "github.com/gokit/tenancykit/pkg"
)

// IdentityBackendImpl defines a concrete struct which implements the methods for the
// IdentityBackend interface. All methods will panic, so add necessary internal logic.
type IdentityBackendImpl struct {
	CountFunc func(var1 context.Context) (int, error)

	DeleteFunc func(var1 context.Context, var2 string) error

	GetFunc func(var1 context.Context, var2 string) (userclaimjwt.Identity, error)

	UpdateFunc func(var1 context.Context, var2 string, var3 userclaimjwt.Identity) error

	RevokeFunc func(var1 context.Context, var2 string) error

	RefreshFunc func(var1 context.Context, var2 string) (userclaimjwt.JWTAuth, error)

	GetAllFunc func(var1 context.Context, var2 string, var3 string, var4 int, var5 int) ([]userclaimjwt.Identity, int, error)

	GrantFunc func(var1 context.Context, var2 pkg.CreateUserSession) (userclaimjwt.JWTAuth, error)

	AuthenticateFunc func(var1 context.Context, var2 string) (pkg.UserClaim, error)
}

// Count implements the IdentityBackend.Count() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) Count(var1 context.Context) (int, error) {

	ret1, ret2 := impl.CountFunc(var1)
	return ret1, ret2

}

// Delete implements the IdentityBackend.Delete() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) Delete(var1 context.Context, var2 string) error {

	ret1 := impl.DeleteFunc(var1, var2)
	return ret1

}

// Get implements the IdentityBackend.Get() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) Get(var1 context.Context, var2 string) (userclaimjwt.Identity, error) {

	ret1, ret2 := impl.GetFunc(var1, var2)
	return ret1, ret2

}

// Update implements the IdentityBackend.Update() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) Update(var1 context.Context, var2 string, var3 userclaimjwt.Identity) error {

	ret1 := impl.UpdateFunc(var1, var2, var3)
	return ret1

}

// Revoke implements the IdentityBackend.Revoke() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) Revoke(var1 context.Context, var2 string) error {

	ret1 := impl.RevokeFunc(var1, var2)
	return ret1

}

// Refresh implements the IdentityBackend.Refresh() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) Refresh(var1 context.Context, var2 string) (userclaimjwt.JWTAuth, error) {

	ret1, ret2 := impl.RefreshFunc(var1, var2)
	return ret1, ret2

}

// GetAll implements the IdentityBackend.GetAll() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) GetAll(var1 context.Context, var2 string, var3 string, var4 int, var5 int) ([]userclaimjwt.Identity, int, error) {

	ret1, ret2, ret3 := impl.GetAllFunc(var1, var2, var3, var4, var5)
	return ret1, ret2, ret3

}

// Grant implements the IdentityBackend.Grant() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) Grant(var1 context.Context, var2 pkg.CreateUserSession) (userclaimjwt.JWTAuth, error) {

	ret1, ret2 := impl.GrantFunc(var1, var2)
	return ret1, ret2

}

// Authenticate implements the IdentityBackend.Authenticate() method for IdentityBackendImpl.
func (impl IdentityBackendImpl) Authenticate(var1 context.Context, var2 string) (pkg.UserClaim, error) {

	ret1, ret2 := impl.AuthenticateFunc(var1, var2)
	return ret1, ret2

}
