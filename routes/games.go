package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarcea/go-ch-test/DB/data"
)

func GetGames(list *data.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := list.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

func GetGameById(list *data.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		var game data.Game
		id := c.Param("id")

		for _, item := range list.Games {
			if item.ID == id {
				game = item
			}

		}

		if game.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "no item found with this ID"})
			return
		}

		c.JSON(http.StatusOK, game)
	}
}

func EditGame(list *data.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		var game data.Game
		id := c.Param("id")

		var body struct {
			Board [8][8]string
			Turn  bool
		}

		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error reading req.body"})
			return
		}

		for i, item := range list.Games {
			if item.ID == id {
				item.Board = body.Board
				item.Turn = !body.Turn
				game = item
				list.Update(item, i)
			}

		}

		if game.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "no game found with this ID"})
			return
		}

		c.JSON(http.StatusOK, game)
	}
}

func AddGame(list *data.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var game data.Game
		// var body struct {
		// 	Board [8][8]string
		// 	Turn  bool
		// }

		// if c.Bind(&body) != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"message": "error reading req.body"})
		// 	return
		// }

		// game = data.Game{
		// 	ID:    uuid.New().String(),
		// 	Board: body.Board,
		// 	Turn:  body.Turn,
		// }
		game := data.InitiateGame()
		list.Add(game)

		c.JSON(200, game)
	}
}
