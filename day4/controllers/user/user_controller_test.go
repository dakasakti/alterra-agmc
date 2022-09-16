package user_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/dakasakti/day2/controllers/user"
	"github.com/dakasakti/day2/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	token, _ = middlewares.CreateToken(1, "dakasakti")
)

func httptestcontext(method string, path string, body io.Reader, token *string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, body)

	if token != nil {
		req.Header.Set("Authorization", "Bearer "+*token)
	}

	res := httptest.NewRecorder()

	context := e.NewContext(req, res)
	context.SetPath(path)
	return context, res
}

func TestUserController(t *testing.T) {
	asserts := assert.New(t)

	t.Run("Test get All User", func(t *testing.T) {
		context, res := httptestcontext("GET", "/users", nil, &token)

		if asserts.NoError(user.GetUsers(context)) {
			asserts.Equal(200, res.Code)
		}
	})

	t.Run("Test get User", func(t *testing.T) {
		context, res := httptestcontext("GET", "/users/profile", nil, &token)

		println(context.Request().Header.Get(echo.HeaderAuthorization))
		context.Request().Header.Get(echo.HeaderAuthorization)
		if asserts.NoError(user.GetUser(context)) {
			asserts.Equal(200, res.Code)
		}
	})
}
