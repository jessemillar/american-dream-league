package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/american-dream-league/handlers"
	"github.com/jessemillar/american-dream-league/helpers"
	"github.com/jessemillar/american-dream-league/views"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME")

	// Construct a new accessor group and connects it to the database
	accessorGroup := new(accessors.AccessorGroup)
	accessorGroup.Open("mysql", database)

	// Constructs a new controller group and gives it the accessor group
	handlerGroup := new(handlers.HandlerGroup)
	handlerGroup.Accessors = accessorGroup

	templateEngine := &helpers.Template{
		Templates: template.Must(template.ParseGlob("public/*/*.html")),
	}

	port := ":9999"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Renderer = templateEngine

	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	router.GET("/api/league", handlerGroup.GetAllLeagues)
	router.GET("/api/league/:id", handlerGroup.GetLeagueByID)
	router.POST("/api/league", handlerGroup.MakeLeague)
	router.PUT("/api/league", handlerGroup.UpdateLeague)
	router.DELETE("/api/league/:id", handlerGroup.DeleteLeagueByID)

	router.GET("/api/user", handlerGroup.GetAllUsers)
	router.GET("/api/user/:id", handlerGroup.GetUserByID)
	router.POST("/api/user", handlerGroup.MakeUser)
	router.PUT("/api/user", handlerGroup.UpdateUser)
	router.DELETE("/api/user/:id", handlerGroup.DeleteUserByID)

	router.GET("/api/name", handlerGroup.GetAllNames)
	router.GET("/api/name/:id", handlerGroup.GetNameByID)
	router.POST("/api/name", handlerGroup.MakeName)
	router.PUT("/api/name", handlerGroup.UpdateName)
	router.DELETE("/api/name/:id", handlerGroup.DeleteNameByID)

	router.GET("/api/email", handlerGroup.GetAllEmails)
	router.GET("/api/email/:id", handlerGroup.GetEmailByID)
	router.POST("/api/email", handlerGroup.MakeEmail)
	router.PUT("/api/email", handlerGroup.UpdateEmail)
	router.DELETE("/api/email/:id", handlerGroup.DeleteEmailByID)

	router.GET("/api/password", handlerGroup.GetAllPasswords)
	router.GET("/api/password/:id", handlerGroup.GetPasswordByID)
	router.POST("/api/password", handlerGroup.MakePassword)
	router.PUT("/api/password", handlerGroup.UpdatePassword)
	router.DELETE("/api/password/:id", handlerGroup.DeletePasswordByID)

	router.Static("/*", "public")
	router.GET("/", views.Frontend)

	router.Start(port)
}
