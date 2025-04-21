package game

type Move int

type Player interface {
	SendMove() Move
	RecvMove(move Move)
	String() string
}
type PlayerTemplate struct {
	History [][2]Move
}

func (p *PlayerTemplate) WrapSendMove(move Move) Move {
	p.History = append(p.History, [2]Move{move, -1})
	return move
}

func (p *PlayerTemplate) LastOppoMove() Move {
	return p.History[len(p.History)-1][1]
}

func (p *PlayerTemplate) LastSelfMove() Move {
	return p.History[len(p.History)-1][0]
}

func (p *PlayerTemplate) RecvMove(move Move) {
	selfLastMove := p.History[len(p.History)-1][0]
	p.History[len(p.History)-1] = [2]Move{selfLastMove, move}
}
