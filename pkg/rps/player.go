package rps

import (
	"rps/pkg/game"
)

type ConstantPlayer struct {
	ConstantMove game.Move
}

func (p *ConstantPlayer) SendMove() game.Move {
	return p.ConstantMove
}

func (p *ConstantPlayer) RecvMove(m game.Move) {
}

func (p *ConstantPlayer) String() string {
	return "constant_player_" + MoveName(p.ConstantMove)
}

type RandomPlayer struct{}

func (p *RandomPlayer) SendMove() game.Move {
	return randMove()
}

func (p *RandomPlayer) RecvMove(m game.Move) {
}

func (p *RandomPlayer) String() string {
	return "random_player"
}

type WannaWinOppoPlayer struct {
	game.PlayerTemplate
}

func (p *WannaWinOppoPlayer) SendMove() game.Move {
	if len(p.History) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		return p.WrapSendMove(winTo(p.LastOppoMove()))
	}
}

func (p *WannaWinOppoPlayer) String() string {
	return "wanna_win_oppo_player"
}

type WannaWinSelfPlayer struct {
	game.PlayerTemplate
}

func (p *WannaWinSelfPlayer) SendMove() game.Move {
	if len(p.History) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		return p.WrapSendMove(winTo(p.LastSelfMove()))
	}
}

func (p *WannaWinSelfPlayer) String() string {
	return "wanna_win_self_player"
}

type WannaLoseOppoPlayer struct {
	game.PlayerTemplate
}

func (p *WannaLoseOppoPlayer) SendMove() game.Move {
	if len(p.History) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		return p.WrapSendMove(loseTo(p.LastOppoMove()))
	}
}

func (p *WannaLoseOppoPlayer) String() string {
	return "wanna_lose_oppo_player"
}

type WannaLoseSelfPlayer struct {
	game.PlayerTemplate
}

func (p *WannaLoseSelfPlayer) SendMove() game.Move {
	if len(p.History) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		return p.WrapSendMove(loseTo(p.LastSelfMove()))
	}
}

func (p *WannaLoseSelfPlayer) String() string {
	return "wanna_lose_self_player"
}
