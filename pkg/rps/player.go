package rps

import (
	"math/rand"
	"rps/pkg/game"
)

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
	return p.moveMap[p.lastMove]
}

func (p *generic1SymmetricPlayer) RecvMove(move game.Move) {
	selfLastMove := p.lastMove[0]
	p.lastMove = [2]game.Move{selfLastMove, move}
}

func makeGeneric1SymmetricPlayer(moveList [3]game.Move) func() game.Player {
	shift := func(mod int, move game.Move) game.Move {
		return game.Move((int(move) + mod) % 3)
	}
	originalMoveMap := map[[2]game.Move]game.Move{
		{Rock, Rock}:     moveList[0],
		{Rock, Paper}:    moveList[2],
		{Rock, Scissors}: moveList[1],
	}
	moveMap := map[[2]game.Move]game.Move{}
	for mod := 0; mod < 3; mod++ {
		for k, v := range originalMoveMap {
			moveMap[[2]game.Move{shift(mod, k[0]), shift(mod, k[1])}] = shift(mod, v)
		}
	}

	return func() game.Player {
		return &generic1SymmetricPlayer{
			isSecondMove: false,
			lastMove:     [2]game.Move{},
			moveMap:      moveMap,
		}
	}
}

func MakeMirrorPlayer() func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("mirror", func(history [][2]game.Move) game.Move {
			if len(history) == 0 {
				return randMove()
			}
			return history[len(history)-1][1]
		})
	}
}
func MakeCyclePlayer() func() game.Player {
	return func() game.Player {
		last := rand.Intn(3)
		return game.MakeLongPlayer("cycle", func(history [][2]game.Move) game.Move {
			last = (last + 1) % 3
			return game.Move(last)
		})
	}
}
func MakeAntiMirrorPlayer() func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("anti_mirror", func(history [][2]game.Move) game.Move {
			if len(history) == 0 {
				return randMove()
			}
			lastMove := history[len(history)-1][0]
			return winTo(lastMove) // Beat what mirror would do
		})
	}
}
func MakeBiasRockPlayer() func() game.Player {
	return func() game.Player {
		return game.MakeLongPlayer("bias_rock", func(history [][2]game.Move) game.Move {
			r := rand.Float64()
			switch {
			case r < 0.7:
				return Rock
			case r < 0.85:
				return Paper
			default:
				return Scissors
			}
		})
	}
}
func MakeReactiveSwitchPlayer() func() game.Player {
	return func() game.Player {
		var lastMove game.Move
		var lastResult int = 0 // 0 = unknown/tie, 1 = win, -1 = lose

		return game.MakeLongPlayer("reactive_switch", func(history [][2]game.Move) game.Move {
			if len(history) == 0 || lastResult <= 0 {
				lastMove = randMove()
			}
			return lastMove
		})
	}
}
func AllFancyPlayers() []func() game.Player {
	return []func() game.Player{
		MakeMirrorPlayer(),
		MakeCyclePlayer(),
		MakeAntiMirrorPlayer(),
		MakeBiasRockPlayer(),
		MakeReactiveSwitchPlayer(),
	}
}
