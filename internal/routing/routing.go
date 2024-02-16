package routing

import (
	"net/http"

	"RatingSystem.Go/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoute(r *gin.RouterGroup) {
	router := r.Group("api")
	router.GET("healthy", healthy)

}
func RegisterSwagger(r *gin.RouterGroup) {
	docs.SwaggerInfo.Title = "Rating System Api"
	docs.SwaggerInfo.Description = "Swagger Interface for Rating system Api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// healthy godoc
// @Summary      Check if application is healthy
// @Success      200
// @Router       /healthy [get]
func healthy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"healthy": true,
	})
}
