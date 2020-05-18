package httpgin

import (
	"examples/task"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	//_ "github.com/swaggo/gin-swagger/example/basic/docs"
	_ "examples/docs"
)

type Response struct {
	Code    string
	Message interface{}
	Error   interface{}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for Jiashiwen to test golang componets .
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func HttpServer() {
	r := gin.New()

	r.Use(Logger4gin(), gin.Recovery())
	r.Use(cors.New(cors.Config{AllowAllOrigins: true}))
	r.GET("/hello", hellofunc)
	r.POST("/somePost", posting)
	r.POST("/api/v1/createtask", CreateTask)

	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//disable swagger
	//r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
	r.Run(":9090")

}

func hellofunc(c *gin.Context) {

	response := gin.H{
		"message": 1,
	}
	c.JSON(200, response)

}

func posting(c *gin.Context) {
	var t task.Task
	var a = 0
	c.Bind(&t)

	if t.TaskStatus == 0 {
		a = 9000
	}

	c.JSON(http.StatusOK, gin.H{
		"TaskId":     t.TaskId,
		"TaskStatus": a,
		"TaskType":   t.TaskType,
	})
}
