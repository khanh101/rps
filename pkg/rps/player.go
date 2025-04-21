package rps

import "rps/pkg/game"

func MakeConstantPlayer(move game.Move) func() game.Player {
	return func() game.Player {
		return game.MakePlayer("constant_"+GetMoveName(move), func(history [][2]game.Move) game.Move {
			return move
		})
	}
}

func MakeRandomPlayer() func() game.Player {
	return func() game.Player {
		return game.MakePlayer("random", func(history [][2]game.Move) game.Move {
			return randMove()
		})
	}
}

func MakeWinSelfPlayer() func() game.Player {
	return func() game.Player {
		return game.MakePlayer("win_self", func(history [][2]game.Move) game.Move {
			if len(history) < 1 {
				return randMove()
			}
			return winTo(history[len(history)-1][0])
		})
	}
}
func MakeLoseSelfPlayer() func() game.Player {
	return func() game.Player {
		return game.MakePlayer("lose_self", func(history [][2]game.Move) game.Move {
			if len(history) < 1 {
				return randMove()
			}
			return loseTo(history[len(history)-1][0])
		})
	}
}
func MakeWinOppoPlayer() func() game.Player {
	return func() game.Player {
		return game.MakePlayer("win_oppo", func(history [][2]game.Move) game.Move {
			if len(history) < 1 {
				return randMove()
			}
			return winTo(history[len(history)-1][1])
		})
	}
}
func MakeLoseOppoPlayer() func() game.Player {
	return func() game.Player {
		return game.MakePlayer("lose_oppo", func(history [][2]game.Move) game.Move {
			if len(history) < 1 {
				return randMove()
			}
			return loseTo(history[len(history)-1][1])
		})
	}
}

func MakeGeneric1Player(moveList [9]game.Move) func() game.Player {
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
		return game.MakePlayer("generic1", func(history [][2]game.Move) game.Move {
			if len(history) < 2 {
				return randMove()
			}
			return moveMap[history[len(history)-1]]
		})
	}
}
