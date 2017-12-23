package tokensetmock

import (
	"errors"
	"time"

	"runtime"

	"github.com/influx6/faux/reflection"

	"github.com/gokit/tenancykit/pkg"

	context "context"
)

// MethodCallForRemove defines a type which holds meta-details about the giving calls associated
// with the TokenSet.Remove() method.
type MethodCallForRemove struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	Target string

	Token string

	// Return values.

	Ret1 error
}

// MatchArguments returns true/false if provider other(MethodCallForRemove) arguments
// values match this existing argument values.
func (me MethodCallForRemove) MatchArguments(other MethodCallForRemove) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.Target, other.Target, true) {
		return false
	}

	if !reflection.MatchElement(me.Token, other.Token, true) {
		return false
	}

	return true
}

// MethodCallForHas defines a type which holds meta-details about the giving calls associated
// with the TokenSet.Has() method.
type MethodCallForHas struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	Target string

	Token string

	// Return values.

	Ret1 bool

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForHas) arguments
// values match this existing argument values.
func (me MethodCallForHas) MatchArguments(other MethodCallForHas) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.Target, other.Target, true) {
		return false
	}

	if !reflection.MatchElement(me.Token, other.Token, true) {
		return false
	}

	return true
}

// MethodCallForAdd defines a type which holds meta-details about the giving calls associated
// with the TokenSet.Add() method.
type MethodCallForAdd struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	Ctx context.Context

	Target string

	Token string

	// Return values.

	Ret1 pkg.Token

	Ret2 error
}

// MatchArguments returns true/false if provider other(MethodCallForAdd) arguments
// values match this existing argument values.
func (me MethodCallForAdd) MatchArguments(other MethodCallForAdd) bool {

	if !reflection.MatchElement(me.Ctx, other.Ctx, true) {
		return false
	}

	if !reflection.MatchElement(me.Target, other.Target, true) {
		return false
	}

	if !reflection.MatchElement(me.Token, other.Token, true) {
		return false
	}

	return true
}

// TokenSetMock defines a type which implements a struct with the
// methods for the TokenSet as fields which allows you provide implementations of
// these functions to provide flexible testing.
type TokenSetMock struct {
	RemoveMethodCalls []MethodCallForRemove

	HasMethodCalls []MethodCallForHas

	AddMethodCalls []MethodCallForAdd
}

// Remove implements the TokenSet.Remove() method for the TokenSet.
func (impl *TokenSetMock) Remove(ctx context.Context, target string, token string) (MethodCallForRemove, error) {
	var caller MethodCallForRemove

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Target = target

	caller.Token = token

	var found bool
	for _, possibleCall := range impl.RemoveMethodCalls {
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

// Has implements the TokenSet.Has() method for the TokenSet.
func (impl *TokenSetMock) Has(ctx context.Context, target string, token string) (MethodCallForHas, error) {
	var caller MethodCallForHas

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Target = target

	caller.Token = token

	var found bool
	for _, possibleCall := range impl.HasMethodCalls {
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

// Add implements the TokenSet.Add() method for the TokenSet.
func (impl *TokenSetMock) Add(ctx context.Context, target string, token string) (MethodCallForAdd, error) {
	var caller MethodCallForAdd

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Target = target

	caller.Token = token

	var found bool
	for _, possibleCall := range impl.AddMethodCalls {
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

// TokenSetSnitch defines a type which implements a struct with the
// methods for the TokenSet as fields which allows you provide implementations of
// these functions to provide flexible testing.
type TokenSetSnitch struct {
	RemoveMethodCalls []MethodCallForRemove
	RemoveFunc        func(ctx context.Context, target string, token string) error

	HasMethodCalls []MethodCallForHas
	HasFunc        func(ctx context.Context, target string, token string) (bool, error)

	AddMethodCalls []MethodCallForAdd
	AddFunc        func(ctx context.Context, target string, token string) (pkg.Token, error)
}

// Remove implements the TokenSet.Remove() method for the TokenSet.
func (impl *TokenSetSnitch) Remove(ctx context.Context, target string, token string) error {
	var caller MethodCallForRemove

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.RemoveMethodCalls = append(impl.RemoveMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Target = target

	caller.Token = token

	ret1 := impl.RemoveFunc(ctx, target, token)

	caller.Ret1 = ret1

	return ret1
}

// Has implements the TokenSet.Has() method for the TokenSet.
func (impl *TokenSetSnitch) Has(ctx context.Context, target string, token string) (bool, error) {
	var caller MethodCallForHas

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.HasMethodCalls = append(impl.HasMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Target = target

	caller.Token = token

	ret1, ret2 := impl.HasFunc(ctx, target, token)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}

// Add implements the TokenSet.Add() method for the TokenSet.
func (impl *TokenSetSnitch) Add(ctx context.Context, target string, token string) (pkg.Token, error) {
	var caller MethodCallForAdd

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.AddMethodCalls = append(impl.AddMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	caller.Ctx = ctx

	caller.Target = target

	caller.Token = token

	ret1, ret2 := impl.AddFunc(ctx, target, token)

	caller.Ret1 = ret1

	caller.Ret2 = ret2

	return ret1, ret2
}
