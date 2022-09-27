package data

type Board [8][8]string

type Game struct {
	ID    string `json:"id"`
	Board Board  `json:"board"`
	Turn  bool   `json:"turn"` // true Withe, false Black
}

type List struct {
	Games []Game
}

var game = InitiateGame()

func New() *List {
	return &List{
		Games: []Game{game},
	}
}

func (r *List) Add(item Game) {
	r.Games = append(r.Games, item)
}

func (r *List) Update(item Game, id int) {
	r.Games[id] = item
}

func (r *List) GetAll() []Game {
	return r.Games
}

func (r *Game) PieceMove(game Game, piece string, from string, to string) Game {
	//TODO: implement logic for a general move
	return game
}
