package main

import (
	"github.com/flpp-oliveira/api-go-swagger-docker/docs"
	_ "github.com/flpp-oliveira/api-go-swagger-docker/docs" // substitua <nome do seu projeto> pelo caminho de importação do seu projeto

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerResponse struct {
	Message string `json:"message"`
}

// @Summary Ping example
// @Description It checks if the server is running
// @Produce  json
// @Success 200 {object} SwaggerResponse
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(200, SwaggerResponse{
		Message: "pong",
	})
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/test")
		{
			eg.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Test endpoint",
				})
			})
			eg.GET("/test", ping)
		}
	}
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // A URL apontando para a especificação da API JSON
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", ping)

	r.Run() // listen and serve on 0.0.0.0:8080
}
