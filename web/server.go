package web

import (
	auth "OpenList/Go/service/auth"
	"OpenList/Go/service/sqlite"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed www/static www/templates
var webFiles embed.FS

type WebServer struct {
	Port string
}

func mustSub(f embed.FS, dir string) fs.FS {
	sub, err := fs.Sub(f, dir)
	if err != nil {
		log.Fatalf("embed sub error: %v", err)
	}
	return sub
}

func (a *WebServer) render(c *gin.Context, templateName string, data gin.H) {
	pageID := strings.TrimSuffix(templateName, ".html")
	data["PageID"] = pageID
	data["ApiURL"] = os.Getenv("API_URL")
	data["WebURL"] = os.Getenv("WEB_URL")
	data["TokenExpirationHours"] = os.Getenv("TOKEN_EXPIRATION_HOURS")

	c.HTML(http.StatusOK, templateName, data)
}

func (a *WebServer) RunWebServer() {
	router := gin.Default()

	funcMap := template.FuncMap{
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
		"fileExists": func(path string) bool {
			_, err := webFiles.Open("www/static/" + path)
			return err == nil
		},
	}

	router.SetFuncMap(funcMap)

	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFS(webFiles, "www/templates/*.html"))
	router.SetHTMLTemplate(tmpl)

	router.StaticFS("/assets", http.FS(mustSub(webFiles, "www/static/assets")))
	router.StaticFS("/core", http.FS(mustSub(webFiles, "www/static/core")))
	router.StaticFS("/components", http.FS(mustSub(webFiles, "www/static/components")))
	router.StaticFS("/pages", http.FS(mustSub(webFiles, "www/static/pages")))
	router.StaticFileFS("/pico.min.css", "assets/css/pico.min.css", http.FS(mustSub(webFiles, "www/static")))

	readSessionUser := func(c *gin.Context) (*sqlite.User, bool) {
		token, err := c.Cookie("openlist_session")
		if err != nil || token == "" {
			return nil, false
		}
		user, err := auth.ValidateSession(token)
		if err != nil {
			return nil, false
		}
		return user, true
	}

	requireWebAuth := func(c *gin.Context) {
		user, ok := readSessionUser(c)
		if !ok {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		if user.FirstLogin && c.Request.URL.Path != "/change-password" {
			c.Redirect(http.StatusFound, "/change-password")
			c.Abort()
			return
		}
		c.Set("webUser", user)
		c.Next()
	}

	router.GET("/login", func(c *gin.Context) {
		if user, ok := readSessionUser(c); ok {
			if user.FirstLogin {
				c.Redirect(http.StatusFound, "/change-password")
			} else {
				c.Redirect(http.StatusFound, "/")
			}
			return
		}
		a.render(c, "login.html", gin.H{"PageName": "Connexion"})
	})

	protected := router.Group("/")
	protected.Use(requireWebAuth)
	{
		protected.GET("/change-password", func(c *gin.Context) {
			a.render(c, "changePassword.html", gin.H{"PageName": "Changer le mot de passe"})
		})

		protected.GET("/", func(c *gin.Context) {
			a.render(c, "index.html", gin.H{"PageName": "Accueil"})
		})

		protected.GET("/settings", func(c *gin.Context) {
			a.render(c, "settings.html", gin.H{"PageName": "Paramètres"})
		})

		protected.GET("/list/:id", func(c *gin.Context) {
			a.render(c, "list.html", gin.H{"PageName": "Liste", "ListID": c.Param("id")})
		})

		protected.GET("/addList", func(c *gin.Context) {
			a.render(c, "addList.html", gin.H{"PageName": "Ajouter une liste"})
		})
	}

	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting Web UI on %s", addr)
	log.Fatal(router.Run(addr))
}
