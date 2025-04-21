package game

type Move int

type Player interface {
	SendMove() Move
	RecvMove(move Move)
	String() string
}

func MakePlayer(name string, sendMoveFunc func(history [][2]Move) Move) Player {
	return &playerTemplate{
		history:      nil,
		sendMoveFunc: sendMoveFunc,
		name:         name,
	}
}

type playerTemplate struct {
	history      [][2]Move
	sendMoveFunc func(history [][2]Move) Move
	name         string
}

func (p *playerTemplate) SendMove() Move {
	move := p.sendMoveFunc(p.history)
	p.history = append(p.history, [2]Move{move, -1})
	return move
}

func (p *playerTemplate) RecvMove(move Move) {
	selfLastMove := p.history[len(p.history)-1][0]
	p.history[len(p.history)-1] = [2]Move{selfLastMove, move}
}

func (p *playerTemplate) String() string {
	return p.name
}
