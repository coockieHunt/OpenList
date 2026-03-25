package main

import (
	"OpenList/Go/cli"
	"OpenList/Go/routes"
	"OpenList/Go/sqlite"
	"OpenList/web"
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
			os.Setenv(key, val)
		}
	}
	return scanner.Err()
}

func main() {
	_ = godotenv.Load()

	sqlite.InitDB(&sqlite.List{}, &sqlite.Item{}, &sqlite.AuthToken{})

	if len(os.Args) > 1 {
		if !cli.HandleCLI() {
			return
		}
	}

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

	// Auth
	router.POST("/api/auth/token", routes.GenrateAuthToken)

	go func() {
		port := os.Getenv("API_PORT")
		if port == "" {
			port = "8080"
		}
		log.Println("Api run at : http://localhost:" + port)
		if err := router.Run("localhost:" + port); err != nil {
			log.Fatalf("Erreur serveur API: %v", err)
		}
	}()

	server := web.WebServer{Port: os.Getenv("WEB_PORT")}
	server.RunWebServer()
}
