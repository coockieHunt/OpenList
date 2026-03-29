package main

import (
	"OpenList/Go/cli"
	routes "OpenList/Go/handler"
	auth "OpenList/Go/service/auth"
	"OpenList/Go/service/sqlite"
	"OpenList/web"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func printBanner(username, password string) {
	line := "─────────────────────────────────────────────────────────"

	fmt.Println()
	fmt.Println(line)

	fmt.Println(`
   ____  ____  _______   ____    _______________
  / __ \/ __ \/ ____/ | / / /   /  _/ ___/_  __/
 / / / / /_/ / __/ /  |/ / /    / / \__ \ / /   
/ /_/ / ____/ /___/ /|  / /____/ / ___/ // /    
\____/_/   /_____/_/ |_/_____/___//____//_/     
`)

	fmt.Println(line)
	fmt.Println("  FIRST TIME LOGIN DETECTED")
	fmt.Println(line)
	fmt.Printf("  ►  Username :  %s\n", username)
	fmt.Printf("  ►  Password :  %s\n", password)
	fmt.Println(line)
	fmt.Println(" /!\\ Change this password immediately upon first login!")
	fmt.Println(line)
	fmt.Println()
}
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

	sqlite.InitDB(&sqlite.List{}, &sqlite.Item{}, &sqlite.User{}, &sqlite.Session{})

	created, username, password, err := auth.EnsureDefaultUser()
	if err != nil {
		log.Fatalf("failed to initialize default user: %v", err)
	}
	if created {
		printBanner(username, password)
	}

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

	router.POST("/api/auth/login", routes.Login)

	authGroup := router.Group("/api")
	authGroup.Use(routes.AuthRequired())
	{
		authGroup.GET("/auth/status", routes.AuthStatus)
		authGroup.POST("/auth/logout", routes.Logout)
		authGroup.POST("/auth/change-password", routes.ChangePassword)

		protected := authGroup.Group("/")
		protected.Use(routes.MustChangePasswordGuard())
		{
			protected.GET("list", routes.GetAllLists)
			protected.GET("list/:idList", routes.GetListByID)
			protected.POST("list", routes.NewList)
			protected.DELETE("list/:idList", routes.DeleteList)
			protected.POST("item/:idList", routes.AddItem)
			protected.PUT("item/:idList/:idItem", routes.ValidateItemID)
			protected.DELETE("item/:idList/:idItem", routes.DeleteItem)
		}
	}

	go func() {
		port := os.Getenv("API_PORT")
		if port == "" {
			port = "8080"
		}
		log.Println("Api run at : http://localhost:" + port)
		if err := router.Run("localhost:" + port); err != nil {
			log.Fatalf("API server error: %v", err)
		}
	}()

	server := web.WebServer{Port: os.Getenv("WEB_PORT")}
	server.RunWebServer()
}
