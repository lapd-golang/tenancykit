// Package tenancykit implements structures for creating a multi-tenant api.
package tenancykit

//go:generate mgokit -generate.dest=./pkg/db -generate.target=./pkg generate
//go:generate httpkit -generate.dest=./pkg/resources -generate.target=./pkg generate

//go:generate mockkit -generate.target=./pkg/db/types -generate.dest=./pkg/db/mocks generate
//go:generate mockkit -generate.target=./pkg -generate.dest=./pkg/db/mocks generate

//go:generate jwtkit -generate.target=./pkg -generate.dest=./pkg/jwt generate
//go:generate mgokit -generate.target=./pkg/jwt/userclaimjwt -generate.dest=./pkg/db generate
//go:generate mockkit -generate.target=./pkg/jwt/userclaimjwt -generate.dest=./pkg/db/mocks generate
