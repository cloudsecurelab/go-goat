package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gogo/protobuf/proto"
)

const jsonTest = `{"name":{"first":"Ruben","last":"Emerita"},"age":47}`

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "You request has been received")
	})

	router.GET("/list", func(c *gin.Context) {
		cmd := exec.Command("ls", "-lah")
		out, _ := cmd.CombinedOutput()
		c.String(200, string(out))
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	return router
}

func initLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	// Set the formatter
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func doJsonStuff() {
	value := gjson.Get(jsonTest, "name.last")
	println(value.String())
}

func main() {

	initLogger()

	doJsonStuff()

	r := initRouter()
	r.Run(":8080")
}
