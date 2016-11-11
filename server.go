package main

import (
	"log"
	"os"

	"github.com/jessemillar/american-dream-league/accessors"
	"github.com/jessemillar/american-dream-league/handlers"
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

	port := ":9999"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	router.Get("/health", health.Check)

	router.Get("/api/generate/hitter", handlerGroup.GenerateHitter)
	router.Get("/api/generate/pitcher", handlerGroup.GeneratePitcher)
	router.Get("/api/league/id/:id", handlerGroup.GetLeagueById)

	log.Println("American Dream League is listening on " + port)
	server := fasthttp.New(port)
	router.Run(server)
}
