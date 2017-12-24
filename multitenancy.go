package tenancykit

// MultiTenancyAPI exposes a API for a multi-tenant api backend, it embodies the necessary
// logic to expose appropriate methods to handle per tenant.
type MultiTenancyAPI struct {
	ochestra TenancyOchestra
	//cache map[string]
}
