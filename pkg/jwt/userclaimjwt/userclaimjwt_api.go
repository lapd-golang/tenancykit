package userclaimjwt

import (
	"errors"
	"strconv"

	"fmt"

	"net/http"

	"encoding/json"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/httputil"

	"github.com/gokit/tenancykit/pkg"
)

var (
	// ContextKeyJWTAuthClaim defines context used to store jwt claim
	// recovered from a authentication request.
	ContextKeyJWTAuthClaim = contextKey("auth-jwt-claim")
)

// contextKey contains provided key value for creating context based key names.
type contextKey string

// String returns string content printable for contextkey.
func (c contextKey) String() string {
	return "JWTKit context key: " + string(c)
}

// JWTAPI defines a struct which holds the http api handlers for providing CRUD
// operations for the provided "UserClaim" type.
type JWTAPI struct {
	metrics  metrics.Metrics
	operator IdentityBackend
}

// New returns a new JWTAPI instance using the provided operator and
// metric.
func New(m metrics.Metrics, backend IdentityBackend) *JWTAPI {
	return &JWTAPI{
		metrics:  m,
		operator: backend,
	}
}

// Info receives an http request to get record info for all available records of type "UserClaim".
//
// Route: /{Route}/info
// Method: INFO
// RESPONSE-BODY: JSON
func (api *JWTAPI) Info(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("JWTAPI.Info"), metrics.WithTrace(m.End()))

	ctx.Header().Set("Content-Type", "application/json")

	api.metrics.Emit(metrics.Info("Info request received"), metrics.WithFields(metrics.Field{
		"url": ctx.Request().URL.String(),
	}))

	total, err := api.operator.Count(ctx.Context())
	if err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to get IdentityRecord record count"), metrics.WithFields(metrics.Field{
			"error": err,
			"url":   ctx.Request().URL.String(),
		}))

		return err
	}

	if err := ctx.JSON(http.StatusOK, IdentityInfo{Total: total}); err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to get serialized Identity record to response writer"), metrics.WithFields(metrics.Field{
			"error": err,
			"url":   ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"url":    ctx.Request().URL.String(),
		"status": http.StatusOK,
	}))

	return nil
}

// Update receives an http request to create a new "UserClaim".
//
// Route: /{Route}/:public_id
// Method: PUT
// BODY: JSON
//
func (api *JWTAPI) Update(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("JWTAPI.Update"), metrics.WithTrace(m.End()))

	ctx.Header().Set("Content-Type", "application/json")

	api.metrics.Emit(metrics.Info("Update request received"), metrics.WithFields(metrics.Field{
		"url": ctx.Request().URL.String(),
	}))

	publicID, ok := ctx.Bag().GetString("public_id")
	if !ok {
		api.metrics.Emit(metrics.Errorf("No public_id provided in params"), metrics.WithFields(metrics.Field{
			"url": ctx.Request().URL.String(),
		}))

		return errors.New("publicId parameter not found")
	}

	var incoming Identity

	if err := json.NewDecoder(ctx.Body()).Decode(&incoming); err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to decode request body"), metrics.WithFields(metrics.Field{
			"error":     err.Error(),
			"public_id": publicID,
			"url":       ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("JSON received"), metrics.WithFields(metrics.Field{
		"data":      incoming,
		"url":       ctx.Request().URL.String(),
		"public_id": publicID,
	}))

	if err := api.operator.Update(ctx.Context(), publicID, incoming); err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to parse params and url.Values"), metrics.WithFields(metrics.Field{
			"error":     err,
			"public_id": publicID,
			"url":       ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"url":       ctx.Request().URL.String(),
		"public_id": publicID,
		"status":    http.StatusNoContent,
	}))

	return ctx.NoContent(http.StatusNoContent)
}

// Delete receives an http request to create a new "UserClaim".
//
// Route: /{Route}/:public_id
// Method: DELETE
//
func (api *JWTAPI) Delete(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("JWTAPI.Delete"), metrics.WithTrace(m.End()))

	api.metrics.Emit(metrics.Info("Delete request received"), metrics.WithFields(metrics.Field{
		"url": ctx.Request().URL.String(),
	}))

	publicID, ok := ctx.Bag().GetString("public_id")
	if !ok {
		api.metrics.Emit(metrics.Errorf("No public_id provided in params"), metrics.WithFields(metrics.Field{
			"url": ctx.Request().URL.String(),
		}))

		return fmt.Errorf("No public_id provided in params")
	}

	api.metrics.Emit(metrics.Info("JSON received"), metrics.WithFields(metrics.Field{
		"url":       ctx.Request().URL.String(),
		"public_id": publicID,
	}))

	if err := api.operator.Delete(ctx.Context(), publicID); err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to delete IdentityRecord record"), metrics.WithFields(metrics.Field{
			"error":     err,
			"public_id": publicID,
			"url":       ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"url":       ctx.Request().URL.String(),
		"public_id": publicID,
		"status":    http.StatusNoContent,
	}))

	return ctx.NoContent(http.StatusNoContent)
}

// Get receives an http request to create a new "UserClaim".
//
// Route: /{Route}/:public_id
// Method: GET
// RESPONSE-BODY: JSON
func (api *JWTAPI) Get(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("JWTAPI.Get"), metrics.WithTrace(m.End()))

	ctx.Header().Set("Content-Type", "application/json")

	api.metrics.Emit(metrics.Info("Get request received"), metrics.WithFields(metrics.Field{
		"url": ctx.Request().URL.String(),
	}))

	publicID, ok := ctx.Bag().GetString("public_id")
	if !ok {
		api.metrics.Emit(metrics.Errorf("No public_id provided in params"), metrics.WithFields(metrics.Field{
			"url": ctx.Request().URL.String(),
		}))

		return errors.New("public_id parameter not found")
	}

	requested, err := api.operator.Get(ctx.Context(), publicID)
	if err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to get IdentityRecord record"), metrics.WithFields(metrics.Field{
			"error":     err,
			"public_id": publicID,
			"url":       ctx.Request().URL.String(),
		}))

		return err
	}

	if err := ctx.JSON(http.StatusOK, requested); err != nil {
		api.metrics.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"error":     err,
			"public_id": publicID,
			"url":       ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"url":       ctx.Request().URL.String(),
		"public_id": publicID,
		"status":    http.StatusOK,
	}))

	return nil
}

// GetAll receives an http request to return all "UserClaim" records.
//
// Route: /{Route}/
// Method: GET
// RESPONSE-BODY: JSON
func (api *JWTAPI) GetAll(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("JWTAPI.GetAll"), metrics.WithTrace(m.End()))

	ctx.Header().Set("Content-Type", "application/json")

	api.metrics.Emit(metrics.Info("GetAll request received"), metrics.WithFields(metrics.Field{
		"url": ctx.Request().URL.String(),
	}))

	var order, orderBy string

	if od, ok := ctx.Bag().Get("order"); ok {
		if ordr, ok := od.(string); ok {
			order = ordr
		} else {
			order = "asc"
		}
	}

	if od, ok := ctx.Bag().Get("orderBy"); ok {
		if ordr, ok := od.(string); ok {
			orderBy = ordr
		} else {
			orderBy = "public_id"
		}
	}

	var err error
	var pageNo, responsePerPage int

	if rpp, ok := ctx.Bag().GetString("responsePerPage"); ok {
		responsePerPage, err = strconv.Atoi(rpp)
		if err != nil {
			api.metrics.Emit(metrics.Errorf("Failed to retrieve responserPerPage number details"), metrics.WithFields(metrics.Field{
				"error": err,
				"url":   ctx.Request().URL.String(),
			}))
		}
	} else {
		responsePerPage = -1
	}

	if pg, ok := ctx.Bag().GetString("page"); ok {
		pageNo, err = strconv.Atoi(pg)
		if err != nil {
			api.metrics.Emit(metrics.Errorf("Failed to retrieve page number details"), metrics.WithFields(metrics.Field{
				"error": err,
				"url":   ctx.Request().URL.String(),
			}))
		}
	} else {
		pageNo = -1
	}

	requested, total, err := api.operator.GetAll(ctx.Context(), order, orderBy, pageNo, responsePerPage)
	if err != nil {
		api.metrics.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"url": ctx.Request().URL.String(),
		}))

		return err
	}

	if err := ctx.JSON(http.StatusOK, Identities{
		Page:            pageNo,
		Records:         requested,
		TotalRecords:    total,
		ResponsePerPage: responsePerPage,
	}); err != nil {
		api.metrics.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"error": err,
			"url":   ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"url":    ctx.Request().URL.String(),
		"status": http.StatusOK,
	}))

	return nil
}

// Grant receives an http request to grant access and request token after sso login.
//
// Route: /{Route}/grant
// Method: POST
// BODY: JSON
//
func (api *JWTAPI) Grant(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("JWTAPI.Grant"), metrics.WithTrace(m.End()))

	ctx.Header().Set("Content-Type", "application/json")

	api.metrics.Emit(metrics.Info("Grant request received"), metrics.WithFields(metrics.Field{
		"url": ctx.Request().URL.String(),
	}))

	var incoming pkg.CreateUserSession

	if err := json.NewDecoder(ctx.Body()).Decode(&incoming); err != nil {
		api.metrics.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"error": err.Error(),
			"url":   ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("JSON received"), metrics.WithFields(metrics.Field{
		"data": incoming,
		"url":  ctx.Request().URL.String(),
	}))

	res, err := api.operator.Grant(ctx.Context(), incoming)
	if err != nil {
		api.metrics.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"error": err,
			"url":   ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"res":    res,
		"url":    ctx.Request().URL.String(),
		"status": http.StatusNoContent,
	}))

	return ctx.JSON(http.StatusOK, res)
}

// Refresh receives an http request to refresh a existing jwt access token using
// provided refresh token.
//
// Route: /{Route}/refresh
// Method: POST
// BODY: JSON (JSON String)
//
func (api *JWTAPI) Refresh(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("JWTAPI.Refresh"), metrics.WithTrace(m.End()))

	ctx.Header().Set("Content-Type", "application/json")
	api.metrics.Emit(metrics.Info("Refresh request received"), metrics.WithFields(metrics.Field{
		"url": ctx.Request().URL.String(),
	}))

	var incoming string
	if err := json.NewDecoder(ctx.Body()).Decode(&incoming); err != nil {
		api.metrics.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"url": ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("JSON received"), metrics.WithFields(metrics.Field{
		"data": incoming,
		"url":  ctx.Request().URL.String(),
	}))

	res, err := api.operator.Refresh(ctx.Context(), incoming)
	if err != nil {
		api.metrics.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"url": ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"url":    ctx.Request().URL.String(),
		"status": http.StatusNoContent,
		"res":    res,
	}))

	return ctx.JSON(http.StatusOK, res)
}

// Authenticate receives an http request to authenticate user request using provided Authorization header.
//
// Route: /{Route}/Authenticate
// Method: GET
//
func (api *JWTAPI) Authenticate(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("JWTAPI.Authenticate"), metrics.WithTrace(m.End()))

	var authHeader string

	if ctx.HasHeader("Authorization", "") {
		authHeader = ctx.GetHeader("Authorization")
	}

	api.metrics.Emit(metrics.Info("Refresh request received"), metrics.WithFields(metrics.Field{
		"header": authHeader,
		"url":    ctx.Request().URL.String(),
	}))

	authType, token, err := httputil.ParseAuthorization(authHeader)
	if err != nil {
		return err
	}

	if authType != "Bearer" {
		return errors.New("only 'Bearer' authorization allowed")
	}

	api.metrics.Emit(metrics.Info("JSON received"), metrics.WithFields(metrics.Field{
		"type":          authType,
		"authorization": token,
		"url":           ctx.Request().URL.String(),
	}))

	res, err := api.operator.Authenticate(ctx.Context(), token)
	if err != nil {
		api.metrics.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"type":          authType,
			"error":         err,
			"url":           ctx.Request().URL.String(),
			"authorization": token,
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"type":          authType,
		"authorization": token,
		"url":           ctx.Request().URL.String(),
		"status":        http.StatusNoContent,
	}))

	ctx.Status(http.StatusNoContent)
	ctx.Set(ContextKeyJWTAuthClaim, res)
	return nil
}
