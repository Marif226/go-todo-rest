package router

// import (
// 	"github.com/labstack/echo/v4"
// )

// func InitRoutes() *echo.Echo {
// 	router := echo.New()

// 	auth := router.Group("/auth")
// 	{
// 		auth.POST("/sign-up", h.signUp)
// 		auth.POST("/sign-in", h.signIn)
// 	}

// 	api := router.Group("/api")
// 	{
// 		lists := api.Group("/lists")
// 		{
// 			lists.POST("/", h.createList)
// 			lists.GET("/", h.getAllLists)
// 			lists.GET("/:id", h.getListById)
// 			lists.PUT("/:id", h.updateList)
// 			lists.DELETE("/:id", h.deleteList)

// 			items := lists.Group(":id/items")
// 			{
// 				items.POST("/", h.createItem)
// 				items.GET("/", h.getAllItems)
// 				items.GET("/:item_id", h.getItemById)
// 				items.PUT("/:item_id", h.updateItem)
// 				items.DELETE("/:item_id", h.deleteItem)
// 			}
// 		}
// 	}

// 	return router
// }