package evaluate

import (
	"github.com/packagewjx/kmeanspp/internal"
	"github.com/packagewjx/kmeanspp/internal/utils"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Calculate the total of the distance between all instance data and its class center.
func ClassCenterDistanceSquareTotal(data [][]float32, class []int, k int) float32 {
	classCenters := internal.CalculateClassCenter(data, class, k)
	sum := float32(0)
	for i := 0; i < len(data); i++ {
		sum += utils.DistanceSquare(data[i], classCenters[class[i]])
	}

	return sum
}

func HopkinsStatistic(data [][]float32) float32 {
	rand.Seed(time.Now().Unix())

	minMaxForDimension := func(d int) (min float32, max float32) {
		min = float32(math.MaxFloat32)
		max = float32(-math.MaxFloat32)
		for i := 0; i < len(data); i++ {
			if data[i][d] < min {
				min = data[i][d]
			} else if data[i][d] > max {
				max = data[i][d]
			}
		}
		return
	}

	generatePoint := func(minMax [][]float32) []float32 {
		p := make([]float32, len(minMax))
		i := 0
		for ; i+3 < len(p); i += 4 {
			p[i] = rand.Float32()*(minMax[i][1]-minMax[i][0]) + minMax[i][0]
			p[i+1] = rand.Float32()*(minMax[i+1][1]-minMax[i+1][0]) + minMax[i+1][0]
			p[i+2] = rand.Float32()*(minMax[i+2][1]-minMax[i+2][0]) + minMax[i+2][0]
			p[i+3] = rand.Float32()*(minMax[i+3][1]-minMax[i+3][0]) + minMax[i+3][0]
		}
		for ; i < len(p); i++ {
			p[i] = rand.Float32()*(minMax[i][1]-minMax[i][0]) + minMax[i][0]
		}
		return p
	}

	// n is sample size
	n := int(math.Max(float64(len(data)/10), 1))
	// x is the sum of distance between points pick from data and its closest neighbor.
	x := float32(0)
	// y is the sum of distance between points pick randomly choose from the data space and the closest point in data.
	y := float32(0)

	picked := sync.Map{}
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			idx := rand.Intn(len(data))
			for _, exist := picked.Load(idx); exist; _, exist = picked.Load(idx) {
				idx = rand.Intn(len(data))
			}
			picked.Store(idx, struct{}{})

			distance := utils.ClosestPointDistance(data[idx], data)
			lock.Lock()
			x += distance
			lock.Unlock()
		}()
	}
	wg.Wait()

	minMax := make([][]float32, len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		min, max := minMaxForDimension(i)
		minMax[i] = []float32{min, max}
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			point := generatePoint(minMax)
			distance := utils.ClosestPointDistance(point, data)
			lock.Lock()
			y += distance
			lock.Unlock()
		}()
	}
	wg.Wait()

	return y / (x + y)
}
