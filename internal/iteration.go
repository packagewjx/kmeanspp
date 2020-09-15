package internal

import (
	"github.com/packagewjx/kmeanspp/internal/utils"
	"math"
	"sync"
)

func CalculateNewCentersPureGo(data [][]float32, centers [][]float32) (newCenter [][]float32, class []int) {
	class = determineClosestCenter(data, centers)
	newCenter = CalculateClassCenter(data, class, len(centers))
	return
}

// Calculate the closest center for all points in data, using centers in centers array.
// Return the center indices array. arr[i] is the center index for data[i].
func determineClosestCenter(data [][]float32, centers [][]float32) []int {
	wg := sync.WaitGroup{}
	res := make([]int, len(data))
	for di, d := range data {
		wg.Add(1)
		go func(pi int, p []float32) {
			defer wg.Done()
			res[pi] = DetermineClosestCenterForPoint(centers, p)
		}(di, d)
	}

	wg.Wait()
	return res
}

// Calculate the closest center in centers for point p. Return the index of the closest center.
// The length of center[i] must be equal to that of p
func DetermineClosestCenterForPoint(centers [][]float32, p []float32) int {
	minC := -1
	minDist := float32(math.MaxFloat32)

	for ci, c := range centers {
		dist := utils.DistanceSquare(p, c)

		if dist < minDist {
			minDist = dist
			minC = ci
		}
	}
	return minC
}

// Calculate the center of all class. data array contains all instances. class[i] is the class number of data[i].
// k is the number of classes.
// Return k centers coordination.
func CalculateClassCenter(data [][]float32, class []int, k int) [][]float32 {
	classSum, classCount := DivideAndConquerClassSumAndCount(data, class, k)

	for i := 0; i < len(classSum); i++ {
		for j := 0; j < len(classSum[i]); j++ {
			classSum[i][j] /= float32(classCount[i])
		}
	}

	return classSum
}

func DivideAndConquerClassSumAndCount(data [][]float32, class []int, k int) (classSum [][]float32, classCount []int) {
	if len(data) < 10 {
		return doCalculateClassSumAndCount(data, class, k)
	}

	split := len(data) / 2
	var leftSum, rightSum [][]float32
	var leftCount, rightCount []int

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		leftSum, leftCount = DivideAndConquerClassSumAndCount(data[:split], class[:split], k)
		wg.Done()
	}()
	go func() {
		rightSum, rightCount = DivideAndConquerClassSumAndCount(data[split:], class[split:], k)
		wg.Done()
	}()
	wg.Wait()

	utils.SumToA2DimensionFloat32(leftSum, rightSum)
	utils.SumToAInt(leftCount, rightCount)

	return leftSum, leftCount
}

func doCalculateClassSumAndCount(data [][]float32, class []int, k int) (classSum [][]float32, classCount []int) {
	classSum = utils.Make2DimensionFloat32Array(k, len(data[0]))
	classCount = make([]int, k)
	for pi, p := range data {
		pc := class[pi]
		classCount[pc]++
		utils.SumToA(classSum[pc], p)
	}
	return
}
