// Package tenancykit implements structures for creating a multi-tenant api.
package tenancykit

//go:generate mgokit -generate.dest=./db generate
//go:generate sqlkit -generate.dest=./db generate
//go:generate httpkit -generate.dest=./api generate
