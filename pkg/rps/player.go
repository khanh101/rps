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
	game.playerTemplate
}

func (p *WannaWinOppoPlayer) SendMove() game.Move {
	if len(p.history) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		return p.WrapSendMove(winTo(p.LastOppoMove()))
	}
}

func (p *WannaWinOppoPlayer) String() string {
	return "wanna_win_oppo_player"
}

type WannaWinSelfPlayer struct {
	game.playerTemplate
}

func (p *WannaWinSelfPlayer) SendMove() game.Move {
	if len(p.history) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		return p.WrapSendMove(winTo(p.LastSelfMove()))
	}
}

func (p *WannaWinSelfPlayer) String() string {
	return "wanna_win_self_player"
}

type WannaLoseOppoPlayer struct {
	game.playerTemplate
}

func (p *WannaLoseOppoPlayer) SendMove() game.Move {
	if len(p.history) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		return p.WrapSendMove(loseTo(p.LastOppoMove()))
	}
}

func (p *WannaLoseOppoPlayer) String() string {
	return "wanna_lose_oppo_player"
}

type WannaLoseSelfPlayer struct {
	game.playerTemplate
}

func (p *WannaLoseSelfPlayer) SendMove() game.Move {
	if len(p.history) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		return p.WrapSendMove(loseTo(p.LastSelfMove()))
	}
}

func (p *WannaLoseSelfPlayer) String() string {
	return "wanna_lose_self_player"
}

type GenericPlayer struct {
	game.playerTemplate
	moveMap map[[2]game.Move]game.Move
}

func (p *GenericPlayer) SendMove() game.Move {
	if len(p.history) == 0 {
		return p.WrapSendMove(randMove())
	} else {
		movePair := [2]game.Move{
			p.LastSelfMove(),
			p.LastOppoMove(),
		}
		return p.WrapSendMove()
	}
}

func (p *GenericPlayer) String() string {
	return "generic_player"
}
