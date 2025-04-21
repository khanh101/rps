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
	// n = a_0 + a_1 3 + a_2 3^2
	var playerMakerList []func() game.Player
	for i := 0; i < 9; i++ {
		moveList := func(n int) [3]game.Move {
			moveList := [3]game.Move{}
			j := 0
			for n > 0 {
				moveList[j] = game.Move(n % 3)
				n /= 3
				j++
			}
			return moveList
		}(i)

		playerMakerList = append(playerMakerList, makeGeneric1SymmetricPlayer(moveList))
	}
	return playerMakerList
}

type generic1SymmetricPlayer struct {
	isSecondMove bool
	lastMove     [2]game.Move
	moveMap      map[[2]game.Move]game.Move
}

func (p *generic1SymmetricPlayer) String() string {
	return "generic"
}

func (p *generic1SymmetricPlayer) SendMove() game.Move {
	if !p.isSecondMove {
		move := randMove()
		p.lastMove = [2]game.Move{move, -1}
		p.isSecondMove = true
		return move
	}
	return p.getFromMoveMap(p.lastMove)
}

func (p *generic1SymmetricPlayer) getFromMoveMap(lastMove [2]game.Move) game.Move {
	shift2 := func(mod int, move2 [2]game.Move) [2]game.Move {
		shiftedMove2 := move2
		shiftedMove2[0] = game.Move((int(shiftedMove2[0]) + mod) % 3)
		shiftedMove2[1] = game.Move((int(shiftedMove2[1]) + mod) % 3)
		return shiftedMove2
	}
	shift1 := func(mod int, move game.Move) game.Move {
		return game.Move((int(move) + mod) % 3)
	}
	for mod := 0; mod < 3; mod++ {
		if move, ok := p.moveMap[shift2(mod, lastMove)]; ok {
			return shift1(3-mod, move)
		}
	}
	panic("unreachable")
}

func (p *generic1SymmetricPlayer) RecvMove(move game.Move) {
	selfLastMove := p.lastMove[0]
	p.lastMove = [2]game.Move{selfLastMove, move}
}

func makeGeneric1SymmetricPlayer(moveList [3]game.Move) func() game.Player {
	moveMap := map[[2]game.Move]game.Move{
		{0, Paper}:    moveList[0],
		{0, Rock}:     moveList[2],
		{0, Scissors}: moveList[1],
	}
	return func() game.Player {
		return &generic1SymmetricPlayer{
			isSecondMove: false,
			lastMove:     [2]game.Move{},
			moveMap:      moveMap,
		}
	}
}
