package game

import "sync"

func Simulate(playerMakerList []func() Player, rounds int, cmp func(m1 Move, m2 Move) int, log func(i int, j int, diff int)) []int {
	n := len(playerMakerList)

	pointVec := make([]int, n)
	pointVecMu := sync.Mutex{}

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				p1, p2 := playerMakerList[i](), playerMakerList[j]()

				diff := 0
				for k := 0; k < rounds; k++ {
					m1, m2 := p1.SendMove(), p2.SendMove()
					p1.RecvMove(m2)
					p2.RecvMove(m1)

					ret := cmp(m1, m2)

					switch ret {
					case +1:
						diff++
					case -1:
						diff--
					default:
					}
				}
				diff = sign(diff)

				{
					pointVecMu.Lock()
					defer pointVecMu.Unlock()
					log(i, j, diff)
					switch diff {
					case +1:
						pointVec[i]++
						pointVec[j]--
					case -1:
						pointVec[i]--
						pointVec[j]++
					default:
					}
				}
			}(i, j)
		}
	}
	wg.Wait()
	return pointVec
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}
