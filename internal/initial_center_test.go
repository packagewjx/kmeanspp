package internal

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestInitialCenter(t *testing.T) {
	data := [][]float32{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
		{5, 5, 5},
		{6, 6, 6},
		{7, 7, 7},
		{8, 8, 8},
		{9, 9, 9},
	}

	center := InitialCenterPureGo(3, data)
	assert.True(t, checkNoDuplicate(center))
}

func TestInitialCenterRandom(t *testing.T) {
	size := 100
	dimension := 4
	k := 5
	round := 10000

	data := make([][]float32, size)
	for i := 0; i < size; i++ {
		data[i] = make([]float32, dimension)
	}

	rand.Seed(time.Now().Unix())
	for r := 0; r < round; r++ {
		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data[i]); j++ {
				data[i][j] = rand.Float32()
			}
		}

		centers := InitialCenterPureGo(k, data)
		assert.True(t, checkNoDuplicate(centers))
		for _, center := range centers {
			assert.Condition(t, func() (success bool) {
				return center >= 0 && center < len(data)
			})
		}
	}
}

func TestChooseCenterNoDuplicate(t *testing.T) {
	data := [][]float32{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
		{4, 4, 4},
		{5, 5, 5},
		{6, 6, 6},
		{7, 7, 7},
		{8, 8, 8},
		{9, 9, 9},
	}

	for i := 0; i < 10000; i++ {
		center := chooseCenter([]int{7, 8}, data)
		assert.NotEqual(t, center, 7)
		assert.NotEqual(t, center, 8)
	}
}

func checkNoDuplicate(arr []int) bool {
	set := make(map[int]struct{})
	for _, i := range arr {
		set[i] = struct{}{}
	}
	return len(arr) == len(set)
}
