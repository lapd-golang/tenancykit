// Package tenancykit implements structures for creating a multi-tenant api.
package tenancykit

//go:generate sqlkit -generate.dest=./pkg/db -generate.target=./pkg generate
//go:generate mgokit -generate.dest=./pkg/db -generate.target=./pkg generate
//go:generate httpkit -generate.dest=./pkg/resources -generate.target=./pkg generate
//go:generate mockkit -generate.target=./pkg/db/types -generate.dest=./pkg/db/mocks generate
