package user

import (
	"strconv"

	"github.com/dakasakti/day2/config"
	"github.com/dakasakti/day2/lib/database"
	"github.com/dakasakti/day2/models"
	"github.com/labstack/echo/v4"
)

var (
	db = database.InitMySQL(config.GetConfig())
)

func GetUsers(c echo.Context) error {
	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": err.Error(),
		})
	}

	if len(users) == 0 {
		return c.JSON(404, echo.Map{
			"message": "data tidak ditemukan",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil get all data user",
		"data":    users,
	})
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "id yang dimasukkan salah",
		})
	}

	var user models.User
	err = db.First(&user, "id = ?", conv_id).Error
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(404, echo.Map{
				"message": "data tidak ditemukan",
			})
		}

		return c.JSON(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil get data user",
		"data":    user,
	})
}

func CreateUser(c echo.Context) error {
	var req models.UserRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "data yang dimasukkan salah",
		})
	}

	user := models.User{
		Fullname: req.Fullname,
		Username: req.Username,
		Password: req.Password,
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil membuat data user",
		"data":    user,
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "id yang dimasukkan salah",
		})
	}

	var req models.UserUpdateRequest

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": err.Error(),
		})
	}

	user := models.User{
		ID:       uint(conv_id),
		Fullname: req.Fullname,
		Username: req.Username,
		Password: req.Password,
	}

	err = db.Updates(&user).Error
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil update data user",
		"data":    user,
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "id yang dimasukkan salah",
		})
	}

	var user models.User
	err = db.Where("id = ?", conv_id).Delete(&user).Error
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "berhasil menghapus data user",
	})
}
