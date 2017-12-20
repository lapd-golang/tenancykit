package routes

import (
	"github.com/gokit/tenancykit"
	"github.com/gokit/tenancykit/pkg/resources/tenantapi"
	"github.com/gokit/tenancykit/pkg/resources/tfrecordapi"
	"github.com/gorilla/mux"
	"github.com/influx6/faux/httputil"
)

// RegisterUsers registers users api routes with the provided gorilla router.
func RegisterUsers(version string, r *mux.Router, users tenancykit.UserAPI, mw httputil.Middleware) {
	mw = httputil.MWi(httputil.GorillaMW, mw)

	r.PathPrefix(version).Path("/users").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			users.GetAll,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users").Methods("POST").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			users.Create,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users/{public_id}").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			users.Get,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users/{public_id}").Methods("PUT").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			users.Update,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users/{public_id}/qr").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			users.TwoFactorQRImage,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users/{public_id}/twofactor/enable").Methods("POST").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			users.EnableTwoFactor,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users/{public_id}/twofactor/disable").Methods("POST").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			users.DisableTwoFactor,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users/{public_id}/").Methods("DELETE").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			users.Delete,
			httputil.BadRequest,
		),
	)))
}

// RegisterTenants registers tenants api routes with the provided gorilla router.
func RegisterTenants(version string, r *mux.Router, tenants tenantapi.TenantHTTP, mw httputil.Middleware) {
	mw = httputil.MWi(httputil.GorillaMW, mw)

	r.PathPrefix(version).Path("/tenants").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			tenants.GetAll,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/tenants/").Methods("POST").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			tenants.Create,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/tenants/{public_id}").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			tenants.Get,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users/{public_id}").Methods("PUT").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			tenants.Update,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/users/{public_id}").Methods("DELETE").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			tenants.Delete,
			httputil.BadRequest,
		),
	)))
}

// RegisterTwoFactorRecords registers tenants api routes with the provided gorilla router.
func RegisterTwoFactorRecords(version string, r *mux.Router, tf tfrecordapi.TFRecordHTTP, mw httputil.Middleware) {
	mw = httputil.MWi(httputil.GorillaMW, mw)

	r.PathPrefix(version).Path("/tfrecords").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			tf.GetAll,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/tfrecords/{public_id}").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			tf.Get,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/tfrecords/{public_id}").Methods("DELETE").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			tf.Delete,
			httputil.BadRequest,
		),
	)))
}

// RegisterUserSessions registers tenants api routes with the provided gorilla router.
func RegisterUserSessions(version string, r *mux.Router, us tenancykit.UserSessionAPI, mw httputil.Middleware) {
	mw = httputil.MWi(httputil.GorillaMW, mw)

	r.PathPrefix(version).Path("/sessions").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			us.GetAll,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/sessions/{public_id}").Methods("GET").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			us.Get,
			httputil.BadRequest,
		),
	)))

	r.PathPrefix(version).Path("/sessions/{public_id}").Methods("DELETE").HandlerFunc(httputil.HTTPFunc(mw(
		httputil.OnError(
			us.Delete,
			httputil.BadRequest,
		),
	)))
}
