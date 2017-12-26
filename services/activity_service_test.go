package services_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/influx6/faux/httputil/httptesting"

	"github.com/influx6/faux/httputil"

	"github.com/gokit/tenancykit/pkg"
	userFixtures "github.com/gokit/tenancykit/pkg/resources/userapi/fixtures"

	"github.com/gokit/tenancykit/api"
	"github.com/gokit/tenancykit/pkg/mock"
	"github.com/gokit/tenancykit/services"
	"github.com/influx6/faux/tests"
)

func TestActivityService(t *testing.T) {
	userBackend := mock.UserBackend()
	roleBackend := mock.RoleBackend()
	activityBackend := mock.ActivityBackend()
	activityService := services.NewActivityService(roleBackend, activityBackend)

	users := api.UserBackend{UserDBBackend: userBackend}

	editPostName := pkg.ActivityName("editPost")
	editPost := pkg.NewActivity(editPostName)
	if err := activityBackend.Create(context.Background(), editPost); err != nil {
		tests.FailedWithError(err, "Should have successfully created actity")
	}
	tests.Passed("Should have successfully created actity")

	userRole := pkg.NewRole("user")
	userRole.AddActivity(editPost)
	if err := roleBackend.Create(context.Background(), userRole); err != nil {
		tests.FailedWithError(err, "Should have successfully created normal user role")
	}
	tests.Passed("Should have successfully created normal user role")

	adminRole := pkg.NewRole("admin-user")
	adminRole.AddActivity(editPost)
	if err := roleBackend.Create(context.Background(), adminRole); err != nil {
		tests.FailedWithError(err, "Should have successfully created admin role")
	}
	tests.Passed("Should have successfully created admin role")

	userBody, err := userFixtures.LoadCreateJSON(userFixtures.UserCreateJSON)
	if err != nil {
		tests.FailedWithError(err, "Should successfully loaded create user fixture")
	}
	tests.Passed("Should successfully loaded create user fixture")

	userBody.PasswordConfirm = userBody.Password

	user, err := users.Create(context.Background(), userBody)
	if err != nil {
		tests.FailedWithError(err, "Should have successfully loaded user record")
	}
	tests.Passed("Should have successfully loaded user record")

	user.AddRole(userRole)

	resource := func(ctx *httputil.Context) error {
		if err := activityService.EnsureAuthorized(ctx, editPostName, func(ctx *httputil.Context, cu pkg.User) error {
			// If user is an admin, that is has adminRole, then allow.
			if cu.HasAnyRole(adminRole) {
				return nil
			}

			// If user is expected user with expected public_id, then allow.
			if cu.HasAnyRole(userRole) && cu.PublicID == user.PublicID {
				return nil
			}

			// if user is not expected user with expected public_id, then disallow nor admin then deny.
			return errors.New("insufficient privileges")
		}); err != nil {
			if herr, ok := err.(httputil.HTTPError); ok {
				ctx.Status(herr.Code)
			} else {
				ctx.Status(http.StatusUnauthorized)
			}

			return err
		}

		ctx.Status(http.StatusOK)
		return nil
	}

	tests.Header("Attempt to access post edit as user for user's posts")
	{
		response := httptest.NewRecorder()
		resourceReq := httptesting.NewRequest("INFO", "/users/post/"+user.PublicID, nil, response)
		resourceReq.Set(pkg.ContextKeyUser, user)

		if err := resource(resourceReq); err != nil {
			tests.FailedWithError(err, "Should have successfully accessed request")
		}
		tests.Passed("Should have successfully accessed request")

		if response.Code != http.StatusOK {
			tests.Info("Expected: %d", http.StatusOK)
			tests.Info("Received: %d", response.Code)
			tests.Failed("Should have successfully received 200 response")
		}
		tests.Passed("Should have successfully received 200 response")
	}

	tests.Header("Attempt to access post edit as admin for user's posts")
	{
		adminUser := user
		adminUser.Roles = nil
		adminUser.PublicID = "2323243twt44y4twtwttwsfdgddfgddf"
		adminUser.AddRole(adminRole)

		response := httptest.NewRecorder()
		resourceReq := httptesting.NewRequest("INFO", "/users/post/"+user.PublicID, nil, response)
		resourceReq.Set(pkg.ContextKeyUser, adminUser)

		if err := resource(resourceReq); err != nil {
			tests.FailedWithError(err, "Should have successfully accessed request")
		}
		tests.Passed("Should have successfully accessed request")

		if response.Code != http.StatusOK {
			tests.Info("Expected: %d", http.StatusOK)
			tests.Info("Received: %d", response.Code)
			tests.Failed("Should have successfully received 200 response")
		}
		tests.Passed("Should have successfully received 200 response")
	}

	tests.Header("Attempt to access post edit as some other normal user for user's posts")
	{
		anyUser := user
		anyUser.Roles = nil
		anyUser.PublicID = "2323243twt44y4twtwttwsfdgddfgddf"
		anyUser.AddRole(userRole)

		response := httptest.NewRecorder()
		resourceReq := httptesting.NewRequest("INFO", "/users/post/"+anyUser.PublicID, nil, response)
		resourceReq.Set(pkg.ContextKeyUser, anyUser)

		if err := resource(resourceReq); err == nil {
			tests.Failed("Should have successfully failed to access request")
		}
		tests.PassedWithError(err, "Should have successfully failed to access request")

		if response.Code == http.StatusOK {
			tests.Info("Expected: %d", http.StatusOK)
			tests.Info("Received: %d", response.Code)
			tests.Failed("Should have successfully failed to receive 200 response")
		}
		tests.Passed("Should have successfully failed receive 200 response")
	}
}
