package data

import (
	"strings"

	"github.com/google/uuid"
)

func InitiateGame() Game {
	var new_game Game
	pieces := [8]string{"R", "N", "B", "Q", "K", "B", "N", "R"}
	for i := 0; i < len(new_game.Board); i++ {
		for j := 0; j < len(new_game.Board[i]); j++ {
			new_game.Board[i][j] = "0"
		}
	}
	for k := 0; k < 8; k++ {
		new_game.Board[0][k] = strings.ToLower(pieces[k])
		new_game.Board[7][k] = pieces[k]
		new_game.Board[1][k] = "p"
		new_game.Board[6][k] = "P"
	}

	new_game.ID = uuid.New().String()
	// new_game.ID = "1"
	new_game.Turn = true

	return new_game
}
