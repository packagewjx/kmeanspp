package internal

import (
	"github.com/packagewjx/kmeanspp/internal/utils"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Select initial seed using K-Means++ methodology. Return seed index in data array.
func InitialCenterPureGo(k int, data [][]float32) []int {
	c := make([]int, 1, k)
	rand.Seed(time.Now().Unix())
	c[0] = rand.Intn(len(data))

	for i := 1; i < k; i++ {
		c = append(c, chooseCenter(c, data))
	}

	return c
}

func chooseCenter(centers []int, data [][]float32) int {
	wg := sync.WaitGroup{}
	dist := make([]float32, len(data))
	for pointIndex, point := range data {
		wg.Add(1)
		go func(pi int, p []float32) {
			defer wg.Done()
			min := float32(math.MaxFloat32)
			for _, c := range centers {
				square := utils.DistanceSquare(p, data[c])
				if square < min {
					min = square
				}
			}
			dist[pi] = min
		}(pointIndex, point)
	}
	wg.Wait()

	// Sum here due to no atomic float
	sum := utils.SumFloat32(dist)
	dist[0] = dist[0] / sum
	for i := 1; i < len(dist); i++ {
		dist[i] = dist[i-1] + dist[i]/sum
	}
	// Due to floating point addition error, set to 1 explicitly.
	for i := len(dist) - 2; i >= 0; i-- {
		if dist[i] == dist[len(dist)-1] {
			dist[i] = 1
		} else {
			break
		}
	}
	dist[len(dist)-1] = 1

	// We don't need to check duplicity here because this binary search algorithm always find the first one equal
	// or larger than p. The probability to choose the previously chosen center, say ci, is 0, meaning that
	// dist[ci] == dist[ci - 1], so the algorithm return ci - 1 (ci - 2 if dist[ci - 1] == dist[ci - 2] and so on),
	// in other word, it always return the point which has not been chosen before.
	p := rand.Float32()
	return utils.BinarySearchFloat32(dist, p)
}
