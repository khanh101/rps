package game

import "sync"

func Simulate(playerMakerList []func() Player, rounds int, cmp func(m1 Move, m2 Move) int, log func(p1 Player, p2 Player, m1 Move, m2 Move, ret int)) []int {
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
					log(p1, p2, m1, m2, ret)

					switch ret {
					case +1:
						diff++
					case -1:
						diff--
					default:
					}
				}

				{
					pointVecMu.Lock()
					defer pointVecMu.Unlock()
					if diff > 0 {
						pointVec[i]++
						pointVec[j]--
					}
					if diff < 0 {
						pointVec[i]--
						pointVec[j]++
					}
				}
			}(i, j)
		}
	}
	wg.Wait()
	return pointVec
}
