package rps

import (
	"math/rand"
	"rps/pkg/game"
)

func GetMoveName(move game.Move) string {
	switch move {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	default:
		return "unknown"
	}
}

const (
	Rock     game.Move = 0
	Paper    game.Move = 1
	Scissors game.Move = 2
)

func Cmp(move1 game.Move, move2 game.Move) int {
	switch [2]game.Move{move1, move2} {
	case [2]game.Move{Rock, Paper}, [2]game.Move{Paper, Scissors}, [2]game.Move{Scissors, Rock}:
		return -1
	case [2]game.Move{Paper, Rock}, [2]game.Move{Scissors, Paper}, [2]game.Move{Rock, Scissors}:
		return +1
	default:
		return 0
	}
}

func randMove() game.Move {
	return game.Move(rand.Intn(3))
}

func loseTo(move game.Move) game.Move {
	switch move {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		panic("wrong move")
	}
}

func winTo(move game.Move) game.Move {
	switch move {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		panic("wrong move")
	}
}
