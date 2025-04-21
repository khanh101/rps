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

				v1, v2 := 0, 0
				for k := 0; k < rounds; k++ {
					m1, m2 := p1.SendMove(), p2.SendMove()
					p1.RecvMove(m2)
					p2.RecvMove(m1)

					ret := cmp(m1, m2)

					switch ret {
					case +1:
						v1++
					case -1:
						v2++
					default:
					}
				}
				diff := func() int {
					if v1 > v2 {
						return +1
					}
					if v2 > v1 {
						return -1
					}
					return 0
				}()

				{
					pointVecMu.Lock()
					defer pointVecMu.Unlock()
					pointVec[i] += v1
					pointVec[j] += v2
					log(i, j, diff)
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
