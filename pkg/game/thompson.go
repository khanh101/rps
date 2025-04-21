package game

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
)

type ThompsonPlayer struct {
	armList    []Player
	winList    []int
	loseList   []int
	cmp        func(m1 Move, m2 Move) int
	lastMove   Move
	lastPlayer int
}

func (p *ThompsonPlayer) String() string {
	return fmt.Sprintf("thompson_%d_arms", len(p.armList))
}

func NewThompsonPlayer(armMakerList []func() Player, cmp func(m1 Move, m2 Move) int) *ThompsonPlayer {
	armList := make([]Player, len(armMakerList))
	for i, playerMaker := range armMakerList {
		armList[i] = playerMaker()
	}
	return &ThompsonPlayer{
		armList:  armList,
		winList:  make([]int, len(armList)),
		loseList: make([]int, len(armList)),
		cmp:      cmp,
	}
}

func (p *ThompsonPlayer) SendMove() Move {
	// choose player
	probList := make([]float64, len(p.armList))
	for i := range probList {
		prob := randBeta(float64(p.winList[i]+1), float64(p.loseList[i]+1))
		probList[i] = prob
	}

	i := argmax(probList)
	// use player
	move := p.armList[i].SendMove()
	p.lastMove = move
	p.lastPlayer = i
	return move
}

func (p *ThompsonPlayer) RecvMove(move Move) {
	ret := p.cmp(p.lastMove, move)
	switch ret {
	case +1:
		p.winList[p.lastPlayer]++
	case -1:
		p.loseList[p.lastPlayer]++
	default:
	}
	p.armList[p.lastPlayer].RecvMove(move)
}

func randBeta(alpha float64, beta float64) float64 {
	return distuv.Beta{
		Alpha: alpha,
		Beta:  beta,
	}.Rand()
}

// argmax samples from a discrete distribution given by weights.
// It returns the selected weight.
func argmax(weights []float64) int {
	if len(weights) == 0 {
		panic("weights is empty")
	}
	i := 0
	for j := 1; j < len(weights); j++ {
		if weights[j] > weights[i] {
			i = j
		}
	}
	return i
}
