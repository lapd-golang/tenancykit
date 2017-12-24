TenancyKit
-----------

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/tenancykit)](https://goreportcard.com/report/github.com/gokit/tenancykit)

TenancyKit applies [Solid Design Principles](https://dave.cheney.net/2016/08/20/solid-go-design), allowing every part of the API to be tested 
and validated as to it's individal unit parts. It's purpose is to provided combination of different packages which provides that can be used 
to build applications requiring tenancy/user management.

## Install

```
go get -u github.com/gokit/tenancykit
```

## Implemented

Tenancykit provides the following subpackages which can be used as desired:

- User API
- Tenant API
- Tranditional Session API
- JWT Identity API
- TwoFactor Record API
- TwoFactor Session API

### Vendoring

Project is vendored using [Govendor](https://github.com/kardianos/govendor).