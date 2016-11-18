package main

import (
	"html/template"
	"log"
	"os"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/american-dream-league/handlers"
	"github.com/jessemillar/american-dream-league/helpers"
	"github.com/jessemillar/american-dream-league/views"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
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
	router.SetRenderer(templateEngine)

	router.Get("/health", health.Check)

	router.Get("/api/league", handlerGroup.GetAllLeagues)
	router.Get("/api/league/:id", handlerGroup.GetLeagueByID)
	router.Post("/api/league", handlerGroup.MakeLeague)
	router.Put("/api/league", handlerGroup.UpdateLeague)
	router.Delete("/api/league/:id", handlerGroup.DeleteLeagueByID)

	router.Get("/api/hitter", handlerGroup.GetHitters)
	router.Get("/api/pitcher", handlerGroup.GetPitchers)
	// router.Get("/api/players", handlerGroup.GeneratePlayers)

	router.Put("/api/hitter", handlerGroup.GenerateHitter)
	router.Put("/api/pitcher", handlerGroup.GeneratePitcher)

	router.Static("/*", "public")
	router.Get("/", views.Frontend)

	log.Println("American Dream League is listening on " + port)
	server := fasthttp.New(port)
	router.Run(server)
}
