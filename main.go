package main

import (
	"fmt"
	"rps/pkg/game"
	"rps/pkg/rps"
	"sort"
)

func argsort(slice []int) []int {
	// Create a slice of indices
	indices := make([]int, len(slice))
	for i := range slice {
		indices[i] = i
	}

	// Sort the indices based on the values in the original slice
	sort.Slice(indices, func(i, j int) bool {
		return slice[indices[i]] < slice[indices[j]]
	})

	return indices
}

func main() {
	var armMakerList = []func() game.Player{
		rps.MakeConstantPlayer(rps.Rock),
		rps.MakeConstantPlayer(rps.Paper),
		rps.MakeConstantPlayer(rps.Scissors),
		rps.MakeRandomPlayer(),
		rps.MakeWinSelfPlayer(),
		rps.MakeLoseSelfPlayer(),
		rps.MakeWinOppoPlayer(),
		rps.MakeLoseOppoPlayer(),
	}
	thompsonPlayer := func() game.Player {
		return game.NewThompsonPlayer(armMakerList, rps.Cmp)
	}
	playerMakerList := append(armMakerList, thompsonPlayer)
	n := len(playerMakerList)

	rounds := 100000
	pointList := game.Simulate(playerMakerList, rounds, rps.Cmp, func(p1 game.Player, p2 game.Player, m1 game.Move, m2 game.Move, ret int) {
		retSMap := map[int]string{
			+1: ">",
			-1: "<",
			0:  "=",
		}
		_ = retSMap
		// fmt.Printf("round %d (%s %s) %s (%s %s)\n", j+1, p1.String(), rps.GetMoveName(m1), retSMap[ret], p2.String(), rps.GetMoveName(m2))
	})
	argSorted := argsort(pointList)
	for i := len(argSorted) - 1; i >= 0; i-- {
		j := argSorted[i]
		fmt.Println(pointList[j], n-1, playerMakerList[j]())
	}
}
