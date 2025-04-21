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

	armList := rps.AllGeneric1Player()
	thompsonPlayerMaker := func() game.Player {
		return game.NewMABPlayer(armList, rps.Cmp)
	}

	var playerMakerList = []func() game.Player{
		rps.MakeConstantPlayer(rps.Rock),
		rps.MakeConstantPlayer(rps.Paper),
		rps.MakeConstantPlayer(rps.Scissors),
		rps.MakeRandomPlayer(),
		rps.MakeWinSelfPlayer(),
		rps.MakeLoseSelfPlayer(),
		rps.MakeWinOppoPlayer(),
		rps.MakeLoseOppoPlayer(),
		thompsonPlayerMaker,
	}
	n := len(playerMakerList)
	playerList := make([]game.Player, n)
	for i := 0; i < n; i++ {
		playerList[i] = playerMakerList[i]()
	}

	rounds := 100000
	pointList := game.Simulate(playerMakerList, rounds, rps.Cmp, func(i int, j int, diff int) {
		diffMap := map[int]string{
			+1: ">",
			-1: "<",
			0:  "=",
		}
		fmt.Printf("%s %s %s\n", playerList[i].String(), diffMap[diff], playerList[j].String())
		_ = diffMap
		// fmt.Printf("round %d (%s %s) %s (%s %s)\n", j+1, p1.String(), rps.GetMoveName(m1), retSMap[ret], p2.String(), rps.GetMoveName(m2))
	})
	argSorted := argsort(pointList)
	for i := len(argSorted) - 1; i >= 0; i-- {
		j := argSorted[i]
		fmt.Println(pointList[j], n-1, playerMakerList[j]())
	}
}
