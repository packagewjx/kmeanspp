package evaluate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestClassCenterDistanceSquareTotal(t *testing.T) {
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

	class := []int{0, 0, 0, 0, 0, 1, 1, 1, 1}

	total := ClassCenterDistanceSquareTotal(data, class, 2)
	assert.Equal(t, float32(45), total)
}

func TestHopkinsStatistic(t *testing.T) {
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

	statistic := HopkinsStatistic(data)
	fmt.Println(statistic)
	assert.Condition(t, func() (success bool) {
		return statistic >= 0 && statistic <= 1
	})

	const Size = 10000
	const Dimension = 30
	const Max = 100
	data = make([][]float32, Size)
	for i := 0; i < Size; i++ {
		data[i] = make([]float32, Dimension)
		for j := 0; j < Dimension; j++ {
			data[i][j] = Max * rand.Float32()
		}
	}
	statistic = HopkinsStatistic(data)
	fmt.Println(statistic)
	assert.Condition(t, func() (success bool) {
		return statistic >= 0 && statistic <= 1
	})
}
