package tenancykit

import (
	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg/db/types"
	"github.com/gokit/tenancykit/pkg/jwt/userclaimjwt"
	"github.com/gokit/tokens"
	"github.com/influx6/faux/metrics"
)

// NewUserSessionDB defines a function type that returns a new user session db for
// giving collection/table name.
type NewUserSessionDB func(metrics.Metrics, string) (types.UserSessionDBBackend, error)

// NewTokenset defines a function type that returns a new tokens.Tokenset for giving
// collection/table name.
type NewTokenset func(metrics.Metrics, string) (tokens.TokenSet, error)

// NewUserDB defines a function type that returns a new tenant db for giving
// collection/table name.
type NewUserDB func(metrics.Metrics, string) (types.UserDBBackend, error)

// NewTwoFactorSessionDB defines a function type that returns a new two factor sesion db
// for giving collection/table name.
type NewTwoFactorSessionDB func(metrics.Metrics, string) (types.TwoFactorSessionDBBackend, error)

// NewTenantDB defines a function type that returns a new tenant db for giving collection/table name.
type NewTenantDB func(metrics.Metrics, string) (types.TenantDBBackend, error)

// NewIdentityConfig defines a function which returns a new
type NewIdentityConfig func(metrics.Metrics, string) (userclaimjwt.JWTConfig, error)

// NewJWTIdentityDB defines a function type that returns a new userclaim jwt identity db for giving collection/table name.
type NewJWTIdentityDB func(metrics.Metrics, string) (userclaimjwt.IdentityDBBackend, error)

// NewTwoFactorRecordDB defines a function type that returns a new tfrecord db for giving
// collection/table name.
type NewTwoFactorRecordDB func(metrics.Metrics, string) (types.TFRecordDBBackend, error)

// TenancyOchestra embodies functions which are used for generating appropriate database instances
// for giving collection/table names.
type TenancyOchestra struct {
	Metrics              metrics.Metrics
	UserDBFunc           NewUserDB
	TokensetFunc         NewTokenset
	TenantDBFunc         NewTenantDB
	JWTIdentityFunc      NewJWTIdentityDB
	UserSessionFunc      NewUserSessionDB
	TwoFactorRecordFunc  NewTwoFactorRecordDB
	TwoFactorSessionFunc NewTwoFactorRecordDB
}

// TenancyDB embodies database types instances which are used in api backend calls.
type TenancyDB struct {
	Metrics            metrics.Metrics
	TwoFactorTokens    tokens.TokenSet
	IdentityBlackList  tokens.TokenSet
	IdentityActiveList tokens.TokenSet
	Users              types.UserDBBackend
	Tenants            types.TenantDBBackend
	TwoFactorRecords   types.TFRecordDBBackend
	UserSessions       types.UserSessionDBBackend
	Identities         userclaimjwt.IdentityDBBackend
	TwoFactorSessions  types.TwoFactorSessionDBBackend
}

// TenancyAPI embodies api instances which are used for routing needs for different api needs.
type TenancyAPI struct {
	Metrics           metrics.Metrics
	Users             api.UserAPI
	Tenants           api.TenantAPI
	Identities        api.IdentityAPI
	UserSessions      api.UserSessionAPI
	TwoFactorSessions api.TwoFactorSessionAPI
	TwoFactorRecords  api.TwoFactorAPI
}
