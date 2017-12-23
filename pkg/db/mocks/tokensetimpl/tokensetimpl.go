package tokensetimpl

import (
	"github.com/gokit/tenancykit/pkg"

	context "context"
)

// TokenSetImpl defines a concrete struct which implements the methods for the
// TokenSet interface. All methods will panic, so add necessary internal logic.
type TokenSetImpl struct {
	RemoveFunc func(ctx context.Context, target string, token string) error

	HasFunc func(ctx context.Context, target string, token string) (bool, error)

	AddFunc func(ctx context.Context, target string, token string) (pkg.Token, error)
}

// Remove implements the TokenSet.Remove() method for TokenSetImpl.
func (impl TokenSetImpl) Remove(ctx context.Context, target string, token string) error {

	ret1 := impl.RemoveFunc(ctx, target, token)
	return ret1

}

// Has implements the TokenSet.Has() method for TokenSetImpl.
func (impl TokenSetImpl) Has(ctx context.Context, target string, token string) (bool, error) {

	ret1, ret2 := impl.HasFunc(ctx, target, token)
	return ret1, ret2

}

// Add implements the TokenSet.Add() method for TokenSetImpl.
func (impl TokenSetImpl) Add(ctx context.Context, target string, token string) (pkg.Token, error) {

	ret1, ret2 := impl.AddFunc(ctx, target, token)
	return ret1, ret2

}
