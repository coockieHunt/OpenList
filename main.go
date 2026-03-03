package main

import (
	"OpenList/Go/routes"
	"OpenList/Go/sqlite"
	"OpenList/web"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//init db end router
	sqlite.InitDB(&sqlite.List{}, &sqlite.Item{})
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:" + os.Getenv("WEB_PORT")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//list Routes
	router.GET("/api/list", routes.GetAllLists)
	router.GET("/api/list/:idList", routes.GetListByID)

	// List
	router.POST("/api/list", routes.NewList)
	router.DELETE("/api/list/:idList", routes.DeleteList)

	// Item
	router.POST("/api/item/:idList", routes.AddItem)
	router.PUT("/api/item/:idList/:idItem", routes.ValidateItemID)
	router.DELETE("/api/item/:idList/:idItem", routes.DeleteItem)

	go func() {
		log.Println("Api run at :http://localhost:" + os.Getenv("API_PORT"))
		if err := router.Run("localhost:" + os.Getenv("API_PORT")); err != nil {
			log.Fatalf("Erreur serveur API: %v", err)
		}
	}()

	server := web.WebServer{Port: os.Getenv("WEB_PORT")}
	server.RunWebServer()
}
