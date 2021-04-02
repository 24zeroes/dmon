package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		dmons := api.Group("/dmons")
		{
			dmons.POST("/", h.createDmon)
			dmons.GET("/", h.getAllDmons)
			dmons.GET("/:id", h.getDmonById)
			dmons.PUT("/:id", h.updateDmon)
			dmons.DELETE("/:id", h.deleteDmon)

			items := dmons.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
