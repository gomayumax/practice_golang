package user

import (
	user "echo-api/service"
	"fmt"
	"github.com/labstack/echo/v4"
)

// Controller is user controlller
type Controller struct{}

// Index action: GET /users
func (pc Controller) Index(c echo.Context) error {
	var s user.Service
	p, _ := s.GetAll()

	//if err != nil {
	//	fmt.Println(err)
	//	// TODO 404
	//} else {
	return c.JSON(200, p)
	//}
}

// Show action: GET /users/:id
func (pc Controller) Show(c echo.Context) error {
	// in the handler for /users?id=<userID>
	//　引数のバインドができる user側で`query:"id"`しておく
	//var user User
	//err := c.Bind(&user)
	//if err != nil {
	//	return c.String(http.StatusBadRequest, "bad request")
	//}

	id := c.Param("id")
	var s user.Service
	p, err := s.GetByID(id)

	if err != nil {
		// TODO 404
		fmt.Println(err)
	} else {
		return c.JSON(200, p)
	}
	return nil
}
