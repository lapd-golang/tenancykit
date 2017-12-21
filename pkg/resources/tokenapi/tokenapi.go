// Package tokenapi provides a auto-generated package which contains a http restful CRUD API for the specific Token struct in package pkg.
//
//
package tokenapi

import (
	"errors"
	"fmt"
	"strconv"

	"context"

	"net/http"

	"encoding/json"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/httputil"

	"github.com/gokit/tenancykit/pkg"
)

// TokenBackend defines an interface which allows the HTTPAPI to divert the final operation of
// the given CRUD request for the "Token" type. This is provided by the user.
// @implement
type TokenBackend interface {
	Delete(context.Context, string) error
	Get(context.Context, string) (pkg.Token, error)
	Update(context.Context, string, pkg.Token) error
	GetAll(context.Context, string, string, int, int) ([]pkg.Token, int, error)
	Create(context.Context, pkg.Token) (pkg.Token, error)
}

// TokenHTTP defines an interface which expose the methods provided by the http backend.
type TokenHTTP interface {
	Get(*httputil.Context) error
	Create(*httputil.Context) error
	Update(*httputil.Context) error
	Delete(*httputil.Context) error
	GetAll(*httputil.Context) error
}

// TokenRecords defines a type to represent the response given to a request for
// all records of the type pkg.Token.
type TokenRecords struct {
	Page            int         `json:"page"`
	ResponsePerPage int         `json:"responsePerPage"`
	TotalRecords    int         `json:"total_records"`
	Records         []pkg.Token `json:"records"`
}

// HTTPAPI defines a struct which holds the http api handlers for providing CRUD
// operations for the provided "Token" type.
type HTTPAPI struct {
	metrics  metrics.Metrics
	operator TokenBackend
}

// New returns a new HTTPAPI instance using the provided operator and
// metric.
func New(m metrics.Metrics, backend TokenBackend) *HTTPAPI {
	return &HTTPAPI{
		metrics:  m,
		operator: backend,
	}
}

// Create receives an http request to create a new "Token".
//
// Route: /{Route}/:public_id
// Method: POST
// BODY: JSON
//
func (api *HTTPAPI) Create(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("HTTPAPI.Create"), metrics.WithTrace(m.End()))

	ctx.Header().Set("Content-Type", "application/json")

	api.metrics.Emit(metrics.Info("Create request received"), metrics.WithFields(metrics.Field{
		"url": ctx.Request().URL.String(),
	}))

	var incoming pkg.Token

	if err := json.NewDecoder(ctx.Body()).Decode(&incoming); err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to parse params and url.Values"), metrics.WithFields(metrics.Field{
			"error": err,
			"url":   ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("JSON received"), metrics.WithFields(metrics.Field{
		"data": incoming,
		"url":  ctx.Request().URL.String(),
	}))

	response, err := api.operator.Create(ctx.Context(), incoming)
	if err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.WithFields(metrics.Field{
			"error": err,
			"url":   ctx.Request().URL.String(),
		}))

		return err
	}

	api.metrics.Emit(metrics.Info("Response Delivered"), metrics.WithFields(metrics.Field{
		"url":    ctx.Request().URL.String(),
		"status": http.StatusCreated,
	}))

	if err := ctx.JSON(http.StatusCreated, response); err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to deliver response"), metrics.WithFields(metrics.Field{
			"error": err,
			"url":   ctx.Request().URL.String(),
		}))
		return err
	}

	return nil
}

// Update receives an http request to create a new "Token".
//
// Route: /{Route}/:public_id
// Method: PUT
// BODY: JSON
//
func (api *HTTPAPI) Update(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("HTTPAPI.Update"), metrics.WithTrace(m.End()))

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

	var incoming pkg.Token

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

// Delete receives an http request to create a new "Token".
//
// Route: /{Route}/:public_id
// Method: DELETE
//
func (api *HTTPAPI) Delete(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("HTTPAPI.Delete"), metrics.WithTrace(m.End()))

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
		api.metrics.Emit(metrics.Errorf("Failed to delete pkg.Token record"), metrics.WithFields(metrics.Field{
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

// Get receives an http request to create a new "Token".
//
// Route: /{Route}/:public_id
// Method: GET
// RESPONSE-BODY: JSON
func (api *HTTPAPI) Get(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("HTTPAPI.Get"), metrics.WithTrace(m.End()))

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
		api.metrics.Emit(metrics.Errorf("Failed to get pkg.Token record"), metrics.WithFields(metrics.Field{
			"error":     err,
			"public_id": publicID,
			"url":       ctx.Request().URL.String(),
		}))

		return err
	}

	if err := ctx.JSON(http.StatusOK, requested); err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to get serialized pkg.Token record to response writer"), metrics.WithFields(metrics.Field{
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

// GetAll receives an http request to return all "Token" records.
//
// Route: /{Route}/
// Method: GET
// RESPONSE-BODY: JSON
func (api *HTTPAPI) GetAll(ctx *httputil.Context) error {
	m := metrics.NewTrace()
	defer api.metrics.Emit(metrics.Info("HTTPAPI.GetAll"), metrics.WithTrace(m.End()))

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
		api.metrics.Emit(metrics.Errorf("Failed to get all pkg.Token record"), metrics.WithFields(metrics.Field{
			"error": err,
			"url":   ctx.Request().URL.String(),
		}))

		return err
	}

	if err := ctx.JSON(http.StatusOK, TokenRecords{
		Page:            pageNo,
		Records:         requested,
		TotalRecords:    total,
		ResponsePerPage: responsePerPage,
	}); err != nil {
		api.metrics.Emit(metrics.Errorf("Failed to get serialized pkg.Token record to response writer"), metrics.WithFields(metrics.Field{
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
