package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) ItemRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up")
		auth.POST("sign-in")
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")
		}
	}

	return router
}
