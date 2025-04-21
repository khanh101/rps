package rps

import "rps/pkg/game"

func MakeConstantPlayer(move game.Move) func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("constant_"+GetMoveName(move), func(history [][2]game.Move) game.Move {
			return move
		})
	}
}

func MakeRandomPlayer() func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("random", func(history [][2]game.Move) game.Move {
			return randMove()
		})
	}
}

func MakeWinSelfPlayer() func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("win_self", func(history [][2]game.Move) game.Move {
			if len(history) < 1 {
				return randMove()
			}
			return winTo(history[len(history)-1][0])
		})
	}
}
func MakeLoseSelfPlayer() func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("lose_self", func(history [][2]game.Move) game.Move {
			if len(history) < 1 {
				return randMove()
			}
			return loseTo(history[len(history)-1][0])
		})
	}
}
func MakeWinOppoPlayer() func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("win_oppo", func(history [][2]game.Move) game.Move {
			if len(history) < 1 {
				return randMove()
			}
			return winTo(history[len(history)-1][1])
		})
	}
}
func MakeLoseOppoPlayer() func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("lose_oppo", func(history [][2]game.Move) game.Move {
			if len(history) < 1 {
				return randMove()
			}
			return loseTo(history[len(history)-1][1])
		})
	}
}

func AllGeneric1Player() []func() game.Player {
	// a_0 + a_1 3 + a_2 3^2 + ... + a_8 3^8 in [0, 3^9)
	var playerMakerList []func() game.Player
	for i := 0; i < 19683; i++ {
		moveList := [9]game.Move{}
		n := i
		for j := 0; j < 9; j++ {
			moveList[j] = game.Move(n % 3)
			n /= 3
		}
		playerMakerList = append(playerMakerList, makeGeneric1Player(moveList))
	}
	return playerMakerList
}

type generic1Player struct {
	isSecondMove bool
	lastMove     [2]game.Move
	moveMap      map[[2]game.Move]game.Move
}

func (p *generic1Player) String() string {
	return "generic"
}

func (p *generic1Player) SendMove() game.Move {
	if !p.isSecondMove {
		move := randMove()
		p.lastMove = [2]game.Move{move, -1}
		p.isSecondMove = true
		return move
	}
	return p.moveMap[p.lastMove]
}

func (p *generic1Player) RecvMove(move game.Move) {
	selfLastMove := p.lastMove[0]
	p.lastMove = [2]game.Move{selfLastMove, move}
}

func makeGeneric1Player(moveList [9]game.Move) func() game.Player {
	moveMap := map[[2]game.Move]game.Move{
		{Rock, Paper}:        moveList[0],
		{Paper, Scissors}:    moveList[1],
		{Scissors, Rock}:     moveList[2],
		{Rock, Rock}:         moveList[3],
		{Paper, Paper}:       moveList[4],
		{Scissors, Scissors}: moveList[5],
		{Paper, Rock}:        moveList[6],
		{Scissors, Paper}:    moveList[7],
		{Rock, Scissors}:     moveList[8],
	}
	return func() game.Player {
		return &generic1Player{
			isSecondMove: false,
			lastMove:     [2]game.Move{},
			moveMap:      moveMap,
		}
	}
}
