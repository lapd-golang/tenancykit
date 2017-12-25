package pkg

// CurrentUser embodies a struct with 3 distinct fields representing
// the current user, it's tenant and twofactor record if enabled.
type CurrentUser struct {
	User      User
	Roles     []Role
	Tenant    *Tenant
	TwoFactor *TFRecord
	TFSession *TwoFactorSession
	Session   *UserSession
}

// UserClaim embodies data attached to a jwt claim which is received
// when using jwt as session handler for SSO Login.
// @jwt(Contract => CreateUserSession)
type UserClaim struct {
	User   User    `json:"user"`
	Tenant ITenant `json:"tenant"`
	Roles  []Role  `json:"roles"`
}
