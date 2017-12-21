package tokendbbackendmock

import (
	"errors"
	"time"

	"runtime"

	"github.com/influx6/faux/reflection"

	context "context"

	pkg "github.com/gokit/tenancykit/pkg"
)

// MethodCallForCount defines a type which holds meta-details about the giving calls associated
// with the TokenDBBackend.Count() method.
type MethodCallForCount struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	// Return values.

	Ret1 int

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForCount) arguments
// values match this existing argument values.
func (me MethodCallForCount) MatchArguments(other MethodCallForCount) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	return true
}

// MethodCallForDelete defines a type which holds meta-details about the giving calls associated
// with the TokenDBBackend.Delete() method.
type MethodCallForDelete struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	PublicID string

	// Return values.

	Ret1 error
}

// MatchArguments returns true/false if provider other(MethodCallForDelete) arguments
// values match this existing argument values.
func (me MethodCallForDelete) MatchArguments(other MethodCallForDelete) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.PublicID, other.PublicID, true) {
		return false
	}

	return true
}

// MethodCallForCreate defines a type which holds meta-details about the giving calls associated
// with the TokenDBBackend.Create() method.
type MethodCallForCreate struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	Elem pkg.Token

	// Return values.

	Ret1 error
}

// MatchArguments returns true/false if provider other(MethodCallForCreate) arguments
// values match this existing argument values.
func (me MethodCallForCreate) MatchArguments(other MethodCallForCreate) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.Elem, other.Elem, true) {
		return false
	}

	return true
}

// MethodCallForGet defines a type which holds meta-details about the giving calls associated
// with the TokenDBBackend.Get() method.
type MethodCallForGet struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	PublicID string

	// Return values.

	Ret1 pkg.Token

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForGet) arguments
// values match this existing argument values.
func (me MethodCallForGet) MatchArguments(other MethodCallForGet) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.PublicID, other.PublicID, true) {
		return false
	}

	return true
}

// MethodCallForUpdate defines a type which holds meta-details about the giving calls associated
// with the TokenDBBackend.Update() method.
type MethodCallForUpdate struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	PublicID string

	Elem pkg.Token

	// Return values.

	Ret1 error
}

// MatchArguments returns true/false if provider other(MethodCallForUpdate) arguments
// values match this existing argument values.
func (me MethodCallForUpdate) MatchArguments(other MethodCallForUpdate) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.PublicID, other.PublicID, true) {
		return false
	}

	if !reflection.MatchElement(me.Elem, other.Elem, true) {
		return false
	}

	return true
}

// MethodCallForGetAllByOrder defines a type which holds meta-details about the giving calls associated
// with the TokenDBBackend.GetAllByOrder() method.
type MethodCallForGetAllByOrder struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	Order string

	OrderBy string

	// Return values.

	Ret1 []pkg.Token

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForGetAllByOrder) arguments
// values match this existing argument values.
func (me MethodCallForGetAllByOrder) MatchArguments(other MethodCallForGetAllByOrder) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.Order, other.Order, true) {
		return false
	}

	if !reflection.MatchElement(me.OrderBy, other.OrderBy, true) {
		return false
	}

	return true
}

// MethodCallForGetByField defines a type which holds meta-details about the giving calls associated
// with the TokenDBBackend.GetByField() method.
type MethodCallForGetByField struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	Key string

	Value interface{}

	// Return values.

	Ret1 pkg.Token

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForGetByField) arguments
// values match this existing argument values.
func (me MethodCallForGetByField) MatchArguments(other MethodCallForGetByField) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.Key, other.Key, true) {
		return false
	}

	if !reflection.MatchElement(me.Value, other.Value, true) {
		return false
	}

	return true
}

// MethodCallForGetAll defines a type which holds meta-details about the giving calls associated
// with the TokenDBBackend.GetAll() method.
type MethodCallForGetAll struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	Order string

	OrderBy string

	Page int

	ResponsePerPage int

	// Return values.

	Ret1 []pkg.Token

	Ret2 int

	Ret3 error
}

// MatchArguments returns true/false if provider other(MethodCallForGetAll) arguments
// values match this existing argument values.
func (me MethodCallForGetAll) MatchArguments(other MethodCallForGetAll) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.Order, other.Order, true) {
		return false
	}

	if !reflection.MatchElement(me.OrderBy, other.OrderBy, true) {
		return false
	}

	if !reflection.MatchElement(me.Page, other.Page, true) {
		return false
	}

	if !reflection.MatchElement(me.ResponsePerPage, other.ResponsePerPage, true) {
		return false
	}

	return true
}

// TokenDBBackendMock defines a type which implements a struct with the
// methods for the TokenDBBackend as fields which allows you provide implementations of
// these functions to provide flexible testing.
type TokenDBBackendMock struct {
	CountMethodCalls []MethodCallForCount

	DeleteMethodCalls []MethodCallForDelete

	CreateMethodCalls []MethodCallForCreate

	GetMethodCalls []MethodCallForGet

	UpdateMethodCalls []MethodCallForUpdate

	GetAllByOrderMethodCalls []MethodCallForGetAllByOrder

	GetByFieldMethodCalls []MethodCallForGetByField

	GetAllMethodCalls []MethodCallForGetAll
}

// Count implements the TokenDBBackend.Count() method for the TokenDBBackend.
func (impl *TokenDBBackendMock) Count(ctx context.Context) (MethodCallForCount, error) {
	var caller MethodCallForCount

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

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

// Delete implements the TokenDBBackend.Delete() method for the TokenDBBackend.
func (impl *TokenDBBackendMock) Delete(ctx context.Context, publicID string) (MethodCallForDelete, error) {
	var caller MethodCallForDelete

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.PublicID = publicID

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

// Create implements the TokenDBBackend.Create() method for the TokenDBBackend.
func (impl *TokenDBBackendMock) Create(ctx context.Context, elem pkg.Token) (MethodCallForCreate, error) {
	var caller MethodCallForCreate

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Elem = elem

	var found bool
	for _, possibleCall := range impl.CreateMethodCalls {
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

// Get implements the TokenDBBackend.Get() method for the TokenDBBackend.
func (impl *TokenDBBackendMock) Get(ctx context.Context, publicID string) (MethodCallForGet, error) {
	var caller MethodCallForGet

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.PublicID = publicID

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

// Update implements the TokenDBBackend.Update() method for the TokenDBBackend.
func (impl *TokenDBBackendMock) Update(ctx context.Context, publicID string, elem pkg.Token) (MethodCallForUpdate, error) {
	var caller MethodCallForUpdate

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.PublicID = publicID

	caller.Elem = elem

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

// GetAllByOrder implements the TokenDBBackend.GetAllByOrder() method for the TokenDBBackend.
func (impl *TokenDBBackendMock) GetAllByOrder(ctx context.Context, order string, orderBy string) (MethodCallForGetAllByOrder, error) {
	var caller MethodCallForGetAllByOrder

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Order = order

	caller.OrderBy = orderBy

	var found bool
	for _, possibleCall := range impl.GetAllByOrderMethodCalls {
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

// GetByField implements the TokenDBBackend.GetByField() method for the TokenDBBackend.
func (impl *TokenDBBackendMock) GetByField(ctx context.Context, key string, value interface{}) (MethodCallForGetByField, error) {
	var caller MethodCallForGetByField

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Key = key

	caller.Value = value

	var found bool
	for _, possibleCall := range impl.GetByFieldMethodCalls {
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

// GetAll implements the TokenDBBackend.GetAll() method for the TokenDBBackend.
func (impl *TokenDBBackendMock) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) (MethodCallForGetAll, error) {
	var caller MethodCallForGetAll

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Order = order

	caller.OrderBy = orderBy

	caller.Page = page

	caller.ResponsePerPage = responsePerPage

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

// TokenDBBackendSnitch defines a type which implements a struct with the
// methods for the TokenDBBackend as fields which allows you provide implementations of
// these functions to provide flexible testing.
type TokenDBBackendSnitch struct {
	CountMethodCalls []MethodCallForCount
	CountFunc        func(ctx context.Context) (int, error)

	DeleteMethodCalls []MethodCallForDelete
	DeleteFunc        func(ctx context.Context, publicID string) error

	CreateMethodCalls []MethodCallForCreate
	CreateFunc        func(ctx context.Context, elem pkg.Token) error

	GetMethodCalls []MethodCallForGet
	GetFunc        func(ctx context.Context, publicID string) (pkg.Token, error)

	UpdateMethodCalls []MethodCallForUpdate
	UpdateFunc        func(ctx context.Context, publicID string, elem pkg.Token) error

	GetAllByOrderMethodCalls []MethodCallForGetAllByOrder
	GetAllByOrderFunc        func(ctx context.Context, order string, orderBy string) ([]pkg.Token, error)

	GetByFieldMethodCalls []MethodCallForGetByField
	GetByFieldFunc        func(ctx context.Context, key string, value interface{}) (pkg.Token, error)

	GetAllMethodCalls []MethodCallForGetAll
	GetAllFunc        func(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Token, int, error)
}

// Count implements the TokenDBBackend.Count() method for the TokenDBBackend.
func (impl *TokenDBBackendSnitch) Count(ctx context.Context) (int, error) {
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

	caller.Ctx = ctx

	ret1, ret2 := impl.CountFunc(ctx)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// Delete implements the TokenDBBackend.Delete() method for the TokenDBBackend.
func (impl *TokenDBBackendSnitch) Delete(ctx context.Context, publicID string) error {
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

	caller.Ctx = ctx

	caller.PublicID = publicID

	ret1 := impl.DeleteFunc(ctx, publicID)

	caller.Ret1 = ret1

	return ret1
}

// Create implements the TokenDBBackend.Create() method for the TokenDBBackend.
func (impl *TokenDBBackendSnitch) Create(ctx context.Context, elem pkg.Token) error {
	var caller MethodCallForCreate

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.CreateMethodCalls = append(impl.CreateMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Elem = elem

	ret1 := impl.CreateFunc(ctx, elem)

	caller.Ret1 = ret1

	return ret1
}

// Get implements the TokenDBBackend.Get() method for the TokenDBBackend.
func (impl *TokenDBBackendSnitch) Get(ctx context.Context, publicID string) (pkg.Token, error) {
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

	caller.Ctx = ctx

	caller.PublicID = publicID

	ret1, ret2 := impl.GetFunc(ctx, publicID)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// Update implements the TokenDBBackend.Update() method for the TokenDBBackend.
func (impl *TokenDBBackendSnitch) Update(ctx context.Context, publicID string, elem pkg.Token) error {
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

	caller.Ctx = ctx

	caller.PublicID = publicID

	caller.Elem = elem

	ret1 := impl.UpdateFunc(ctx, publicID, elem)

	caller.Ret1 = ret1

	return ret1
}

// GetAllByOrder implements the TokenDBBackend.GetAllByOrder() method for the TokenDBBackend.
func (impl *TokenDBBackendSnitch) GetAllByOrder(ctx context.Context, order string, orderBy string) ([]pkg.Token, error) {
	var caller MethodCallForGetAllByOrder

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.GetAllByOrderMethodCalls = append(impl.GetAllByOrderMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Order = order

	caller.OrderBy = orderBy

	ret1, ret2 := impl.GetAllByOrderFunc(ctx, order, orderBy)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// GetByField implements the TokenDBBackend.GetByField() method for the TokenDBBackend.
func (impl *TokenDBBackendSnitch) GetByField(ctx context.Context, key string, value interface{}) (pkg.Token, error) {
	var caller MethodCallForGetByField

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.GetByFieldMethodCalls = append(impl.GetByFieldMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Key = key

	caller.Value = value

	ret1, ret2 := impl.GetByFieldFunc(ctx, key, value)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// GetAll implements the TokenDBBackend.GetAll() method for the TokenDBBackend.
func (impl *TokenDBBackendSnitch) GetAll(ctx context.Context, order string, orderBy string, page int, responsePerPage int) ([]pkg.Token, int, error) {
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

	caller.Ctx = ctx

	caller.Order = order

	caller.OrderBy = orderBy

	caller.Page = page

	caller.ResponsePerPage = responsePerPage

	ret1, ret2, ret3 := impl.GetAllFunc(ctx, order, orderBy, page, responsePerPage)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	caller.Ret3 = ret3

	return ret1, ret2, ret3
}
