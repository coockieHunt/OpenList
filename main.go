package main

import (
	"OpenList/routes"
    "OpenList/sqlite"
	"github.com/gin-gonic/gin"
)

func main() {

    //init db end router
    sqlite.InitDB(&sqlite.List{}, &sqlite.Item{})
    router := gin.Default()
    
    //list Routes
    router.GET("/list", routes.GetAllLists)
    router.GET("/list/:id", routes.GetListByID)

    // List
    router.POST("/list", routes.NewList)
    router.DELETE("/list/:id", routes.DeleteList)

    // Item
	router.POST("/item/:id", routes.AddItem)
    router.DELETE("/item/:id", routes.DeleteItem)
    router.PUT("/item/:id", routes.ValidateItemID)

    // run server
    router.Run("localhost:8080")
}