package rps

import (
	"math/rand"
	"rps/pkg/adversarial_game"
)

func MoveName(move adversarial_game.Move) string {
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
	Rock     adversarial_game.Move = 0
	Paper    adversarial_game.Move = 1
	Scissors adversarial_game.Move = 2
)

func Cmp(move1 adversarial_game.Move, move2 adversarial_game.Move) int {
	switch [2]adversarial_game.Move{move1, move2} {
	case [2]adversarial_game.Move{Rock, Paper}, [2]adversarial_game.Move{Paper, Scissors}, [2]adversarial_game.Move{Scissors, Rock}:
		return -1
	case [2]adversarial_game.Move{Paper, Rock}, [2]adversarial_game.Move{Scissors, Paper}, [2]adversarial_game.Move{Rock, Scissors}:
		return +1
	default:
		return 0
	}
}

func randMove() adversarial_game.Move {
	return adversarial_game.Move(rand.Intn(3))
}

func loseTo(move adversarial_game.Move) adversarial_game.Move {
	switch move {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		panic("wrong Constantadversarial_game.Move")
	}
}

func winTo(move adversarial_game.Move) adversarial_game.Move {
	switch move {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		panic("wrong Constantadversarial_game.Move")
	}
}
