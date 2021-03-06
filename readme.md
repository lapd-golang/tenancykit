TenancyKit
-----------

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit)](https://goreportcard.com/report/github.com/gokit/tenancykit)
[![Travis Build](https://travis-ci.org/gokit/tenancykit.svg?branch=master)](https://travis-ci.org/gokit/tenancykit#)

TenancyKit applies [Solid Design Principles](https://dave.cheney.net/2016/08/20/solid-go-design), allowing every part of the API to be tested 
and validated as to it's individal unit parts. It's purpose is to provided combination of different packages which provides that can be used 
to build applications requiring tenancy/user management.

TenancyKit is rather opinionate and is not for everyone has it has specific opinions about things that would not be changable due to the choices made
with the code. In such cases, building a custom tenant package that fits your need is most suited than using tenancyKit. This does not mean the code is not
extensible, but that due to specific choices, internal cores are not likely to change and my differ from intended approach of user, where in such 
cases a custom solution is more better.

*In the future, I hope to extract as much repeatable parts of tenancykit into small code generators that can be used 
to code generate different areas of projects to better allow extensability for users.*

## Install

```
go get -u github.com/gokit/tenancykit
```

## Implemented

Tenancykit provides the following subpackages which can be used as desired:

- User API
- Tenant API
- JWT Identity API
- TwoFactor Record API
- TwoFactor Session API
- Activity based Roles Authorization
- Session API (SingleUsers and MultiTenant Users)

### Vendoring

Project is vendored using [Govendor](https://github.com/kardianos/govendor).