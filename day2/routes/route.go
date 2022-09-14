package routes

import (
	"net/http"

	"github.com/dakasakti/day2/controllers/book"
	"github.com/dakasakti/day2/controllers/user"
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
	books.POST("", book.CreateBook)

	book_id := books.Group("/:id")
	book_id.GET("", book.GetBook)
	book_id.PUT("", book.UpdateBook)
	book_id.DELETE("", book.DeleteBook)

	// users
	users := e.Group("/users")
	users.GET("", user.GetUsers)
	users.POST("", user.CreateUser)

	user_id := users.Group("/:id")
	user_id.GET("", user.GetUser)
	user_id.PUT("", user.UpdateUser)
	user_id.DELETE("", user.DeleteUser)

	return e
}
