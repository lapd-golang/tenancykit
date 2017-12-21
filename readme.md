TenancyKit
-----------

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit)](https://goreportcard.com/report/github.com/gokit/tenancykit)

TenancyKit is a package for the different structural parts needed to create multitenant apis.

TenancyKit applies [Solid Design Principles](https://dave.cheney.net/2016/08/20/solid-go-design), allowing every part of the API to be tested 
and validated as to it's individal units behaviour. It provides a solid foundation for building either a tenant microservice or as a drop-in 
for a projects user and tenant management module for your backend application.

*TenancyKit also adopts code generation at it's core and you will find a lot of packages that make up the core are actually generated, quickly providing
cohesive but flexible customization as to structure fields and behaviours.*

## Install

```
go get -u github.com/gokit/tenancykit
```

