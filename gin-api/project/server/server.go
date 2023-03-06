package server

import (
	"fmt"
	user "gin-api/project/controller"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	err := r.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := user.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
	}

	return r
}
