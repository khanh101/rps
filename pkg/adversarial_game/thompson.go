package adversarial_game

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
)

type ThompsonPlayer struct {
	playerList []Player
	wList      []int
	lList      []int
	lastPlayer int
}

func (p *ThompsonPlayer) String() string {
	playerListStr := ""
	for _, player := range p.playerList {
		playerListStr += player.String() + ","
	}
	return fmt.Sprintf("thompson_player_[%s]", playerListStr)
}

func NewThompsonPlayer(playerList []Player) *ThompsonPlayer {
	return &ThompsonPlayer{
		playerList: playerList,
		wList:      make([]int, len(playerList)),
		lList:      make([]int, len(playerList)),
	}
}

func randBeta(alpha float64, beta float64) float64 {
	return distuv.Beta{
		Alpha: alpha,
		Beta:  beta,
	}.Rand()
}

// randDiscrete samples from a discrete distribution given by weights.
// It returns the selected weight.
func randDiscrete(weights []float64) int {
	if len(weights) == 0 {
		panic("weights must not be empty")
	}

	// Compute the total weight
	total := 0.0
	for _, w := range weights {
		if w < 0 {
			panic("weights must be non-negative")
		}
		total += w
	}

	if total == 0 {
		panic("sum of weights must be greater than zero")
	}

	// Generate a random number in [0, total)
	r := rand.Float64() * total

	// Find the corresponding weight
	cumulative := 0.0
	for i, w := range weights {
		cumulative += w
		if r < cumulative {
			return i
		}
	}

	// Fallback due to floating point precision issues
	return len(weights) - 1
}

func (p *ThompsonPlayer) MakeMove(history History) Move {
	if len(history) != 0 {
		// update history
		ret := history[len(history)-1].Ret
		switch ret {
		case +1:
			p.wList[p.lastPlayer] += 1
		case -1:
			p.lList[p.lastPlayer] += 1
		default:

		}
	}
	probList := make([]float64, len(p.playerList))
	for i := range probList {
		prob := randBeta(float64(p.wList[i]+1), float64(p.lList[i]+1))
		probList[i] = prob
	}

	i := randDiscrete(probList)
	p.lastPlayer = i
	return p.playerList[p.lastPlayer].MakeMove(history)
}
