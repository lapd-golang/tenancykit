package identitybackendmock

import (
	"errors"
	"time"

	"runtime"

	"github.com/influx6/faux/reflection"

	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"

	context "context"

	pkg "github.com/gokit/tenancykit/pkg"
)

// MethodCallForCount defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.Count() method.
type MethodCallForCount struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	// Return values.

	Ret1 int

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForCount) arguments
// values match this existing argument values.
func (me MethodCallForCount) MatchArguments(other MethodCallForCount) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	return true
}

// MethodCallForDelete defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.Delete() method.
type MethodCallForDelete struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	Var2 string

	// Return values.

	Ret1 error
}

// MatchArguments returns true/false if provider other(MethodCallForDelete) arguments
// values match this existing argument values.
func (me MethodCallForDelete) MatchArguments(other MethodCallForDelete) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	if !reflection.MatchElement(me.Var2, other.Var2, true) {
		return false
	}

	return true
}

// MethodCallForGet defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.Get() method.
type MethodCallForGet struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	Var2 string

	// Return values.

	Ret1 userclaimjwt.Identity

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForGet) arguments
// values match this existing argument values.
func (me MethodCallForGet) MatchArguments(other MethodCallForGet) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	if !reflection.MatchElement(me.Var2, other.Var2, true) {
		return false
	}

	return true
}

// MethodCallForUpdate defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.Update() method.
type MethodCallForUpdate struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	Var2 string

	Var3 userclaimjwt.Identity

	// Return values.

	Ret1 error
}

// MatchArguments returns true/false if provider other(MethodCallForUpdate) arguments
// values match this existing argument values.
func (me MethodCallForUpdate) MatchArguments(other MethodCallForUpdate) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	if !reflection.MatchElement(me.Var2, other.Var2, true) {
		return false
	}

	if !reflection.MatchElement(me.Var3, other.Var3, true) {
		return false
	}

	return true
}

// MethodCallForRevoke defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.Revoke() method.
type MethodCallForRevoke struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	Var2 string

	// Return values.

	Ret1 error
}

// MatchArguments returns true/false if provider other(MethodCallForRevoke) arguments
// values match this existing argument values.
func (me MethodCallForRevoke) MatchArguments(other MethodCallForRevoke) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	if !reflection.MatchElement(me.Var2, other.Var2, true) {
		return false
	}

	return true
}

// MethodCallForRefresh defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.Refresh() method.
type MethodCallForRefresh struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	Var2 string

	// Return values.

	Ret1 userclaimjwt.JWTAuth

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForRefresh) arguments
// values match this existing argument values.
func (me MethodCallForRefresh) MatchArguments(other MethodCallForRefresh) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	if !reflection.MatchElement(me.Var2, other.Var2, true) {
		return false
	}

	return true
}

// MethodCallForGetAll defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.GetAll() method.
type MethodCallForGetAll struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	Var2 string

	Var3 string

	Var4 int

	Var5 int

	// Return values.

	Ret1 []userclaimjwt.Identity

	Ret2 int

	Ret3 error
}

// MatchArguments returns true/false if provider other(MethodCallForGetAll) arguments
// values match this existing argument values.
func (me MethodCallForGetAll) MatchArguments(other MethodCallForGetAll) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	if !reflection.MatchElement(me.Var2, other.Var2, true) {
		return false
	}

	if !reflection.MatchElement(me.Var3, other.Var3, true) {
		return false
	}

	if !reflection.MatchElement(me.Var4, other.Var4, true) {
		return false
	}

	if !reflection.MatchElement(me.Var5, other.Var5, true) {
		return false
	}

	return true
}

// MethodCallForGrant defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.Grant() method.
type MethodCallForGrant struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	Var2 pkg.CreateUserSession

	// Return values.

	Ret1 userclaimjwt.JWTAuth

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForGrant) arguments
// values match this existing argument values.
func (me MethodCallForGrant) MatchArguments(other MethodCallForGrant) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	if !reflection.MatchElement(me.Var2, other.Var2, true) {
		return false
	}

	return true
}

// MethodCallForAuthenticate defines a type which holds meta-details about the giving calls associated
// with the IdentityBackend.Authenticate() method.
type MethodCallForAuthenticate struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Var1 context.Context

	Var2 string

	// Return values.

	Ret1 pkg.UserClaim

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForAuthenticate) arguments
// values match this existing argument values.
func (me MethodCallForAuthenticate) MatchArguments(other MethodCallForAuthenticate) bool {

	if !reflection.MatchElement(me.Var1, other.Var1, true) {
		return false
	}

	if !reflection.MatchElement(me.Var2, other.Var2, true) {
		return false
	}

	return true
}

// IdentityBackendMock defines a type which implements a struct with the
// methods for the IdentityBackend as fields which allows you provide implementations of
// these functions to provide flexible testing.
type IdentityBackendMock struct {
	CountMethodCalls []MethodCallForCount

	DeleteMethodCalls []MethodCallForDelete

	GetMethodCalls []MethodCallForGet

	UpdateMethodCalls []MethodCallForUpdate

	RevokeMethodCalls []MethodCallForRevoke

	RefreshMethodCalls []MethodCallForRefresh

	GetAllMethodCalls []MethodCallForGetAll

	GrantMethodCalls []MethodCallForGrant

	AuthenticateMethodCalls []MethodCallForAuthenticate
}

// Count implements the IdentityBackend.Count() method for the IdentityBackend.
func (impl *IdentityBackendMock) Count(var1 context.Context) (MethodCallForCount, error) {
	var caller MethodCallForCount

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	var found bool
	for _, possibleCall := range impl.CountMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.Ret2 = possibleCall.Ret2

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// Delete implements the IdentityBackend.Delete() method for the IdentityBackend.
func (impl *IdentityBackendMock) Delete(var1 context.Context, var2 string) (MethodCallForDelete, error) {
	var caller MethodCallForDelete

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	var found bool
	for _, possibleCall := range impl.DeleteMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// Get implements the IdentityBackend.Get() method for the IdentityBackend.
func (impl *IdentityBackendMock) Get(var1 context.Context, var2 string) (MethodCallForGet, error) {
	var caller MethodCallForGet

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	var found bool
	for _, possibleCall := range impl.GetMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.Ret2 = possibleCall.Ret2

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// Update implements the IdentityBackend.Update() method for the IdentityBackend.
func (impl *IdentityBackendMock) Update(var1 context.Context, var2 string, var3 userclaimjwt.Identity) (MethodCallForUpdate, error) {
	var caller MethodCallForUpdate

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	caller.Var3 = var3

	var found bool
	for _, possibleCall := range impl.UpdateMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// Revoke implements the IdentityBackend.Revoke() method for the IdentityBackend.
func (impl *IdentityBackendMock) Revoke(var1 context.Context, var2 string) (MethodCallForRevoke, error) {
	var caller MethodCallForRevoke

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	var found bool
	for _, possibleCall := range impl.RevokeMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// Refresh implements the IdentityBackend.Refresh() method for the IdentityBackend.
func (impl *IdentityBackendMock) Refresh(var1 context.Context, var2 string) (MethodCallForRefresh, error) {
	var caller MethodCallForRefresh

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	var found bool
	for _, possibleCall := range impl.RefreshMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.Ret2 = possibleCall.Ret2

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// GetAll implements the IdentityBackend.GetAll() method for the IdentityBackend.
func (impl *IdentityBackendMock) GetAll(var1 context.Context, var2 string, var3 string, var4 int, var5 int) (MethodCallForGetAll, error) {
	var caller MethodCallForGetAll

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	caller.Var3 = var3

	caller.Var4 = var4

	caller.Var5 = var5

	var found bool
	for _, possibleCall := range impl.GetAllMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.Ret2 = possibleCall.Ret2

			caller.Ret3 = possibleCall.Ret3

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// Grant implements the IdentityBackend.Grant() method for the IdentityBackend.
func (impl *IdentityBackendMock) Grant(var1 context.Context, var2 pkg.CreateUserSession) (MethodCallForGrant, error) {
	var caller MethodCallForGrant

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	var found bool
	for _, possibleCall := range impl.GrantMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.Ret2 = possibleCall.Ret2

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// Authenticate implements the IdentityBackend.Authenticate() method for the IdentityBackend.
func (impl *IdentityBackendMock) Authenticate(var1 context.Context, var2 string) (MethodCallForAuthenticate, error) {
	var caller MethodCallForAuthenticate

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	var found bool
	for _, possibleCall := range impl.AuthenticateMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.Ret2 = possibleCall.Ret2

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// IdentityBackendSnitch defines a type which implements a struct with the
// methods for the IdentityBackend as fields which allows you provide implementations of
// these functions to provide flexible testing.
type IdentityBackendSnitch struct {
	CountMethodCalls []MethodCallForCount
	CountFunc        func(var1 context.Context) (int, error)

	DeleteMethodCalls []MethodCallForDelete
	DeleteFunc        func(var1 context.Context, var2 string) error

	GetMethodCalls []MethodCallForGet
	GetFunc        func(var1 context.Context, var2 string) (userclaimjwt.Identity, error)

	UpdateMethodCalls []MethodCallForUpdate
	UpdateFunc        func(var1 context.Context, var2 string, var3 userclaimjwt.Identity) error

	RevokeMethodCalls []MethodCallForRevoke
	RevokeFunc        func(var1 context.Context, var2 string) error

	RefreshMethodCalls []MethodCallForRefresh
	RefreshFunc        func(var1 context.Context, var2 string) (userclaimjwt.JWTAuth, error)

	GetAllMethodCalls []MethodCallForGetAll
	GetAllFunc        func(var1 context.Context, var2 string, var3 string, var4 int, var5 int) ([]userclaimjwt.Identity, int, error)

	GrantMethodCalls []MethodCallForGrant
	GrantFunc        func(var1 context.Context, var2 pkg.CreateUserSession) (userclaimjwt.JWTAuth, error)

	AuthenticateMethodCalls []MethodCallForAuthenticate
	AuthenticateFunc        func(var1 context.Context, var2 string) (pkg.UserClaim, error)
}

// Count implements the IdentityBackend.Count() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) Count(var1 context.Context) (int, error) {
	var caller MethodCallForCount

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.CountMethodCalls = append(impl.CountMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	ret1, ret2 := impl.CountFunc(var1)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// Delete implements the IdentityBackend.Delete() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) Delete(var1 context.Context, var2 string) error {
	var caller MethodCallForDelete

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.DeleteMethodCalls = append(impl.DeleteMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	ret1 := impl.DeleteFunc(var1, var2)

	caller.Ret1 = ret1

	return ret1
}

// Get implements the IdentityBackend.Get() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) Get(var1 context.Context, var2 string) (userclaimjwt.Identity, error) {
	var caller MethodCallForGet

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.GetMethodCalls = append(impl.GetMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	ret1, ret2 := impl.GetFunc(var1, var2)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// Update implements the IdentityBackend.Update() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) Update(var1 context.Context, var2 string, var3 userclaimjwt.Identity) error {
	var caller MethodCallForUpdate

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.UpdateMethodCalls = append(impl.UpdateMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	caller.Var3 = var3

	ret1 := impl.UpdateFunc(var1, var2, var3)

	caller.Ret1 = ret1

	return ret1
}

// Revoke implements the IdentityBackend.Revoke() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) Revoke(var1 context.Context, var2 string) error {
	var caller MethodCallForRevoke

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.RevokeMethodCalls = append(impl.RevokeMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	ret1 := impl.RevokeFunc(var1, var2)

	caller.Ret1 = ret1

	return ret1
}

// Refresh implements the IdentityBackend.Refresh() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) Refresh(var1 context.Context, var2 string) (userclaimjwt.JWTAuth, error) {
	var caller MethodCallForRefresh

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.RefreshMethodCalls = append(impl.RefreshMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	ret1, ret2 := impl.RefreshFunc(var1, var2)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// GetAll implements the IdentityBackend.GetAll() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) GetAll(var1 context.Context, var2 string, var3 string, var4 int, var5 int) ([]userclaimjwt.Identity, int, error) {
	var caller MethodCallForGetAll

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.GetAllMethodCalls = append(impl.GetAllMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	caller.Var3 = var3

	caller.Var4 = var4

	caller.Var5 = var5

	ret1, ret2, ret3 := impl.GetAllFunc(var1, var2, var3, var4, var5)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	caller.Ret3 = ret3

	return ret1, ret2, ret3
}

// Grant implements the IdentityBackend.Grant() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) Grant(var1 context.Context, var2 pkg.CreateUserSession) (userclaimjwt.JWTAuth, error) {
	var caller MethodCallForGrant

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.GrantMethodCalls = append(impl.GrantMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	ret1, ret2 := impl.GrantFunc(var1, var2)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// Authenticate implements the IdentityBackend.Authenticate() method for the IdentityBackend.
func (impl *IdentityBackendSnitch) Authenticate(var1 context.Context, var2 string) (pkg.UserClaim, error) {
	var caller MethodCallForAuthenticate

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.AuthenticateMethodCalls = append(impl.AuthenticateMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Var1 = var1

	caller.Var2 = var2

	ret1, ret2 := impl.AuthenticateFunc(var1, var2)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}
