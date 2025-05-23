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
	var armList []func() game.Player
	armList = append(armList, rps.AllGeneric1Player()...)
	armList = append(armList, rps.AllFancyPlayers()...)
	mabPlayer := func() game.Player {
		return game.NewMABPlayer(armList, rps.Cmp)
	}

	var playerMakerList = []func() game.Player{
		mabPlayer,
		rps.MakeRandomPlayer(),
		rps.MakeConstantPlayer(rps.Rock),
		rps.MakeConstantPlayer(rps.Paper),
		rps.MakeConstantPlayer(rps.Scissors),
		rps.MakeWinSelfPlayer(),
		rps.MakeLoseSelfPlayer(),
		rps.MakeWinOppoPlayer(),
		rps.MakeLoseOppoPlayer(),
	}
	playerMakerList = append(playerMakerList, rps.AllFancyPlayers()...)
	n := len(playerMakerList)
	playerList := make([]game.Player, n)
	for i := 0; i < n; i++ {
		playerList[i] = playerMakerList[i]()
	}

	rounds := 200
	pointList := game.Simulate(playerMakerList, rounds, rps.Cmp, func(i int, j int, diff int) {
		diffMap := map[int]string{
			+1: ">",
			-1: "<",
			0:  "=",
		}
		_ = diffMap
		if i == 0 || j == 0 {
			fmt.Printf("%s %s %s\n", playerList[i].String(), diffMap[diff], playerList[j].String())
		}
	})
	argSorted := argsort(pointList)
	for i := len(argSorted) - 1; i >= 0; i-- {
		j := argSorted[i]
		fmt.Printf("player %s won %d out of %d times\n", playerMakerList[j]().String(), pointList[j], rounds*(n-1))
	}
}
