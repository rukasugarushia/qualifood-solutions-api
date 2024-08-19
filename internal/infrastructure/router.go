package infrastructure

import (
	"qualifood-solutions-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Grouping all book related routes
	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("/", handlers.GetAllBooks)
		bookRoutes.GET("/:id", handlers.GetBookByID)
		bookRoutes.POST("/", handlers.CreateBook)
		bookRoutes.PUT("/:id", handlers.UpdateBook)
		bookRoutes.DELETE("/:id", handlers.DeleteBook)
	}

	return router
}
