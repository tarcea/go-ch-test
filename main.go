package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tarcea/go-ch-test/DB/data"
	"github.com/tarcea/go-ch-test/middlewares"
	"github.com/tarcea/go-ch-test/routes"
)

var list = data.New()

func main() {

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	r.GET("/ws", routes.Websocket)
	r.GET("/games", routes.GetGames(list))
	r.POST("/games", routes.AddGame(list))
	r.GET("/games/:id", routes.GetGameById(list))
	r.PUT("/games/:id", routes.EditGame(list))

	r.Run()
}
