package routes

import (
	"net/http"

	"github.com/dakasakti/day2/controllers/book"
	"github.com/dakasakti/day2/controllers/user"
	"github.com/dakasakti/day2/middlewares"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// books
	books := e.Group("/books")
	books.GET("", book.GetBooks)
	books.POST("", book.CreateBook, middlewares.JWTSign())

	book_id := books.Group("/:id")
	book_id.GET("", book.GetBook)
	book_id.PUT("", book.UpdateBook, middlewares.JWTSign())
	book_id.DELETE("", book.DeleteBook, middlewares.JWTSign())

	// users
	users := e.Group("/users")
	users.GET("", user.GetUsers, middlewares.JWTSign())
	users.POST("", user.CreateUser)

	//profile
	profile := e.Group("/users/profile", middlewares.JWTSign())
	profile.GET("", user.GetUser)
	profile.PUT("", user.UpdateUser)
	profile.DELETE("", user.DeleteUser)

	// login
	e.POST("/login", user.Login)

	return e
}
