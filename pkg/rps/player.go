package rps

import (
	"math/rand"
	"rps/pkg/adversarial_game"
)

type ConstantPlayer struct {
	ConstantMove adversarial_game.Move
}

func (p *ConstantPlayer) MakeMove(history adversarial_game.History) adversarial_game.Move {
	return p.ConstantMove
}

func (p *ConstantPlayer) String() string {
	return "constant_player_" + MoveName(p.ConstantMove)
}

type WannaWinOppoPlayer struct{}

func (p *WannaWinOppoPlayer) MakeMove(history adversarial_game.History) adversarial_game.Move {
	if len(history) == 0 {
		return randMove()
	}
	return winTo(history[len(history)-1].OpponentMove)
}

func (p *WannaWinOppoPlayer) String() string {
	return "wanna_win_oppo_player"
}

type WannaWinSelfPlayer struct{}

func (p *WannaWinSelfPlayer) MakeMove(history adversarial_game.History) adversarial_game.Move {
	if len(history) == 0 {
		return randMove()
	}
	return winTo(history[len(history)-1].MyMove)
}

func (p *WannaWinSelfPlayer) String() string {
	return "wanna_win_self_player"
}

type WannaLoseOppoPlayer struct{}

func (p *WannaLoseOppoPlayer) MakeMove(history adversarial_game.History) adversarial_game.Move {
	if len(history) == 0 {
		return randMove()
	}
	return loseTo(history[len(history)-1].OpponentMove)
}

func (p *WannaLoseOppoPlayer) String() string {
	return "wanna_lose_oppo_player"
}

type WannaLoseSelfPlayer struct{}

func (p *WannaLoseSelfPlayer) MakeMove(history adversarial_game.History) adversarial_game.Move {
	if len(history) == 0 {
		return randMove()
	}
	return loseTo(history[len(history)-1].MyMove)
}

func (p *WannaLoseSelfPlayer) String() string {
	return "wanna_lose_self_player"
}

type RandomPlayer struct{}

func (p *RandomPlayer) MakeMove(history adversarial_game.History) adversarial_game.Move {
	return randMove()
}

func (p *RandomPlayer) String() string {
	return "random_player"
}

type RandomHumanPlayer struct {
	RepeatChance float64
}

func (p *RandomHumanPlayer) MakeMove(history adversarial_game.History) adversarial_game.Move {
	if len(history) == 0 {
		return randMove()
	}
	lastMove := history[len(history)-1].MyMove
	if rand.Float64() < p.RepeatChance {
		return lastMove
	}
	for {
		move := randMove()
		if move != lastMove {
			return move
		}
	}
}

func (p *RandomHumanPlayer) String() string {
	return "random_human_player"
}
