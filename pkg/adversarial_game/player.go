package adversarial_game

type Move int

type HistoryPoint struct {
	MyMove       Move
	OpponentMove Move
}

type History = []HistoryPoint

type Player interface {
	MakeMove(history History) Move
	String() string
}
