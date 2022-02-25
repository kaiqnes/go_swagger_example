package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go_swagger_example/docs"
	"log"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   	API Support
// @contact.url    	http://www.swagger.io/support
// @contact.email  	support@swagger.io

// @license.name  	Apache 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html

// @host			localhost:8080
// @BasePath  		/

func main() {
	router := gin.Default()
	router.GET("/", helloWorld)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// helloWorld godoc
// @Summary      Returns a 'Hello World' string
// @Description  Get a 'Hello World' message
// @Tags         Hello_World
// @Accept       json
// @Produce      json
func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World")
}
