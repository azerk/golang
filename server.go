// server.go

package main

import (
	"net/http"

	//	"github.com/codegangsta/martini"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	//	m := martini.Classic()
	//	//	m.Get("/", func() string {
	//	//		return "Hello world!"
	//	//	})
	//	m.NotFound(func() string {
	//		return "i am error"
	//	})
	//	m.Get("/", func() string {
	//		return "i am get"
	//	})
	//	m.Get("/post/:name", func(parmas martini.Params) string {
	//		return "i am " + parmas["name"]
	//	})
	//	m.Run()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":3000")

}
