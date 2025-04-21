package main

import (
	"fmt"
	"rps/pkg/adversarial_game"
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
	var playerList = []adversarial_game.Player{
		&rps.WannaWinPlayer{},
		&rps.WannaLosePlayer{},
		&rps.ConstantPlayer{ConstantMove: rps.Rock},
		&rps.ConstantPlayer{ConstantMove: rps.Paper},
		&rps.ConstantPlayer{ConstantMove: rps.Scissors},
		&rps.RandomPlayer{},
		&rps.RandomHumanPlayer{},
	}
	rounds := 10000
	pointList := adversarial_game.Simulate(playerList, rounds, rps.Cmp, func(j int, p1 adversarial_game.Player, m1 adversarial_game.Move, p2 adversarial_game.Player, m2 adversarial_game.Move, ret int) {
		retSMap := map[int]string{
			+1: ">",
			-1: "<",
			0:  "=",
		}
		_ = retSMap
		// fmt.Printf("round %d (%s %s) %s (%s %s)\n", j+1, p1.String(), rps.MoveName(m1), retSMap[ret], p2.String(), rps.MoveName(m2))
	})
	argSorted := argsort(pointList)
	for i := len(argSorted) - 1; i >= 0; i-- {
		j := argSorted[i]
		fmt.Println(pointList[j], playerList[j])
	}
}
