package routes

import (
	"to-do-list/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Allow all origins (debugging only!)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	r := router.Group("/api")

	r.POST("/todos", controllers.CreateTodo)

	r.GET("/todos", controllers.GetAllTodos)

	r.PUT("/todo/:id", controllers.EditTodo)

	r.DELETE("/todo/:id", controllers.DeleteTodoById)

	return router
}
