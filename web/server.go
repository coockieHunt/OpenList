package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	Port string
}

func (a *WebServer) RunWebServer() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, fmt.Errorf("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
	})

	router.Static("/assets", "./web/www/assets")
	router.Static("/api", "./web/www/api")
	router.Static("/append", "./web/www/append")
	router.StaticFile("/pico.min.css", "./web/www/pico.min.css")

	router.LoadHTMLGlob("web/www/*.html")

	// home
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "OpenList",
		})
	})

	//settings
	router.GET("/settings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "settings.html", gin.H{
			"PageName": "Settings",
		})
	})

	//list
	router.GET("/list/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"PageName": "List",
			"ListID":   c.Param("id"),
		})
	})

	router.GET("/addList", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addList.html", gin.H{
			"PageName": "Add List",
		})
	})

	//user
	router.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{
			"PageName": "User Management",
		})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "userLogin.html", gin.H{
			"PageName": "Login",
		})
	})

	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting Web UI on %s", addr)
	log.Fatal(router.Run(addr))
}
