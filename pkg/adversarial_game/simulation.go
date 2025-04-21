package adversarial_game

import (
	"math/rand"
)

func Simulate(playerList []Player, rounds int, cmp func(m1 Move, m2 Move) int, log func(j int, p1 Player, m1 Move, p2 Player, m2 Move, ret int)) []int {
	pointList := make([]int, len(playerList))

	playerHistoryMat := func(n int) [][]History {
		playerHistoryMat := make([][]History, n)
		for i := 0; i < n; i++ {
			playerHistoryMat[i] = make([]History, n)
		}
		return playerHistoryMat
	}(len(playerList))

	roundList := func(n int, r int) [][2]int {
		playerPairList := [][2]int{}
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				playerPairList = append(playerPairList, [2]int{i, j})
			}
		}

		roundList := make([][2]int, r*len(playerPairList))
		for i := 0; i < len(roundList); i++ {
			roundList[i] = playerPairList[i%len(playerPairList)]
		}

		rand.Shuffle(len(roundList), func(i int, j int) {
			roundList[i], roundList[j] = roundList[j], roundList[i]
		})
		return roundList
	}(len(playerList), rounds)

	for j, round := range roundList {
		i1, i2 := round[0], round[1]
		p1, p2 := playerList[i1], playerList[i2]
		h1, h2 := playerHistoryMat[i1][i2], playerHistoryMat[i2][i1]

		m1, m2 := p1.MakeMove(h1), p2.MakeMove(h2)
		playerHistoryMat[i1][i2] = append(playerHistoryMat[i1][i2], HistoryPoint{
			MyMove:       m1,
			OpponentMove: m2,
		})
		playerHistoryMat[i2][i1] = append(playerHistoryMat[i2][i1], HistoryPoint{
			MyMove:       m2,
			OpponentMove: m1,
		})
		ret := cmp(m1, m2)

		log(j+1, p1, m1, p2, m2, ret)
		switch ret {
		case +1:
			pointList[i1] += 1
			pointList[i2] -= 1
		case -1:
			pointList[i1] -= 1
			pointList[i2] += 1
		default:
		}
	}
	return pointList
}
