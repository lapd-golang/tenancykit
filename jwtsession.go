package tenancykit

import "github.com/gokit/tenancykit/pkg/db/types"

// JWTSession implements session management using jwt tokens.
// It provides an alternative from the underline session system
// built with UserSession, and focus on jwt claims and encoded
// data as it's basis.
type JWTSession struct {
	UserBackend         types.UserDBBackend
	TenantBackend       types.TenantDBBackend
	TFBackend           types.TFRecordDBBackend
	TFSessionsBackend   types.TwoFactorSessionDBBackend
	IsNotFoundErrorFunc func(error) bool
}
