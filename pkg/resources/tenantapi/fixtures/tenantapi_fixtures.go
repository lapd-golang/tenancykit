package fixtures

import (
	"encoding/json"

	"github.com/gokit/tenancykit/pkg"
)

// json fixtures ...
var (
	TenantJSON = `{


    "public_id":	"9v7hppfblc9b1h01x5vaduflkehhcu",

    "created_at":	"2017-12-21 15:13:27.87264 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:27.872647 +0000 UTC",

    "name":	"Kevin Wood",

    "email":	"0Oliver@Riffpath.net"

}`

	TenantCreateJSON = `{


    "email":	"8Larson@Teklist.edu",

    "name":	"Douglas Robinson"

}`

	TenantUpdateJSON = `{


    "created_at":	"2017-12-21 15:13:27.873194 +0000 UTC",

    "updated_at":	"2017-12-21 15:13:27.873199 +0000 UTC",

    "name":	"Keith Garza",

    "email":	"lGarrett@Vinder.org",

    "public_id":	"vzftudc2m6vfl72bw2c9rng3ir8vlz"

}`
)

// LoadCreateJSON returns a new instance of a pkg.CreateTenant.
func LoadCreateJSON(content string) (pkg.CreateTenant, error) {
	var elem pkg.CreateTenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.CreateTenant{}, err
	}

	return elem, nil
}

// LoadUpdateJSON returns a new instance of a pkg.Tenant.
func LoadUpdateJSON(content string) (pkg.Tenant, error) {
	var elem pkg.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Tenant{}, err
	}

	return elem, nil
}

// LoadTenantJSON returns a new instance of a pkg.Tenant.
func LoadTenantJSON(content string) (pkg.Tenant, error) {
	var elem pkg.Tenant

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return pkg.Tenant{}, err
	}

	return elem, nil
}
