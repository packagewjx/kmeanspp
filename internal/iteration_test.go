package internal

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestCalculateCenter(t *testing.T) {
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

	center := [][]float32{
		{4, 0, 0},
		{7, 0, 0},
	}

	expectCenter := [][]float32{
		{3, 3, 3},
		{7.5, 7.5, 7.5},
	}
	expectClass := []int{0, 0, 0, 0, 0, 1, 1, 1, 1}

	actualCenter, actualClass := CalculateNewCentersPureGo(data, center)
	assert.Equal(t, expectCenter, actualCenter)
	assert.Equal(t, expectClass, actualClass)
}

func TestCalculateClassSumAndCount(t *testing.T) {
	const Size = 100
	const Max = 100
	const Dimension = 6
	const K = 6
	rand.Seed(time.Now().Unix())

	data := make([][]float32, Size)
	class := make([]int, Size)
	sum := make([][]float32, K)
	count := make([]int, K)
	for i := 0; i < len(sum); i++ {
		sum[i] = make([]float32, Dimension)
	}

	for i := 0; i < Size; i++ {
		class[i] = rand.Intn(K)
		count[class[i]]++
		p := make([]float32, Dimension)
		for j := 0; j < len(p); j++ {
			p[j] = float32(rand.Intn(Max))
			sum[class[i]][j] += p[j]
		}
		data[i] = p
	}

	actualSum, actualCount := DivideAndConquerClassSumAndCount(data, class, K)
	assert.Equal(t, sum, actualSum)
	assert.Equal(t, count, actualCount)
}

func TestDoCalculateClassSumAndCount(t *testing.T) {
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

	class := []int{0, 0, 0, 1, 1, 2, 2, 2, 2}

	classSum, count := doCalculateClassSumAndCount(data, class, 3)
	assert.Equal(t, []int{3, 2, 4}, count)
	assert.Equal(t, [][]float32{
		{6, 6, 6},
		{9, 9, 9},
		{30, 30, 30},
	}, classSum)
}

func TestDetermineCenter(t *testing.T) {
	data := [][]float32{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
	}

	center := [][]float32{
		{1, 0},
		{5, 0},
	}

	result := determineClosestCenter(data, center)
	assert.Equal(t, []int{0, 0, 0, 1, 1}, result)

	data = [][]float32{
		{1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5},
	}

	center = [][]float32{
		{1, 1, 1, 1, 0, 0},
		{5, 5, 5, 5, 0, 0},
	}

	result = determineClosestCenter(data, center)

	assert.Equal(t, []int{0, 0, 0, 1, 1}, result)
}
