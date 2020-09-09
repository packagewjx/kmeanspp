package kmeanspp

import (
	"github.com/packagewjx/kmeanspp/internal"
	"golang.org/x/sys/cpu"
)

var (
	initialCenterFunc       func(k int, data [][]float32) []int
	calculateNewCentersFunc func(data [][]float32, centers [][]float32) ([][]float32, []int)
)

func init() {
	if cpu.X86.HasAVX2 {
		initialCenterFunc = internal.InitialCenterPureGo
		calculateNewCentersFunc = internal.CalculateNewCentersPureGo
	} else if cpu.X86.HasAVX {
		initialCenterFunc = internal.InitialCenterPureGo
		calculateNewCentersFunc = internal.CalculateNewCentersPureGo
	} else if cpu.X86.HasSSE41 {
		initialCenterFunc = internal.InitialCenterPureGo
		calculateNewCentersFunc = internal.CalculateNewCentersPureGo
	} else if cpu.X86.HasSSE3 {
		initialCenterFunc = internal.InitialCenterPureGo
		calculateNewCentersFunc = internal.CalculateNewCentersPureGo
	} else if cpu.X86.HasSSE2 {
		initialCenterFunc = internal.InitialCenterPureGo
		calculateNewCentersFunc = internal.CalculateNewCentersPureGo
	} else {
		initialCenterFunc = internal.InitialCenterPureGo
		calculateNewCentersFunc = internal.CalculateNewCentersPureGo
	}
}

// K-Means++ Entrance function. k is the number of classes. round is the number of iteration. data array contains
// all instances data. data[i][j] is jth feature of ith instance.
func KMeansPP(k, round int, data [][]float32) []int {
	if k <= 0 || round <= 0 {
		return []int{}
	}

	centerIndices := initialCenterFunc(k, data)
	centers := make([][]float32, k)
	for i, ci := range centerIndices {
		centers[i] = data[ci]
	}

	var class []int
	for r := 0; r < round; r++ {
		centers, class = calculateNewCentersFunc(data, centers)
	}

	return class
}
