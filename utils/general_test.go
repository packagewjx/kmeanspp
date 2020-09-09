package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		arr    []float32
		k      float32
		expect int
	}{
		{arr: []float32{1, 2, 3, 4}, k: 2.5, expect: 2},
		{arr: []float32{1, 2, 3, 4}, k: 2, expect: 1},
		{arr: []float32{1, 2, 3, 4}, k: 5, expect: 4},
		{arr: []float32{1, 2, 3, 4}, k: 0, expect: 0},
		{arr: []float32{1, 2, 3, 4, 5}, k: 1, expect: 0},
		{arr: []float32{1, 2, 3, 4, 5}, k: 2, expect: 1},
		{arr: []float32{1, 2, 3, 4, 5}, k: 3, expect: 2},
		{arr: []float32{1, 2, 3, 4, 5}, k: 3.5, expect: 3},
		{arr: []float32{1, 2, 2, 3, 4}, k: 2, expect: 1},
		{arr: []float32{1, 2, 2, 3, 4}, k: 1, expect: 0},
		{arr: []float32{1, 2, 2, 3, 4}, k: 2.5, expect: 3},
		{arr: []float32{1, 2, 2, 3, 4}, k: 1.5, expect: 1},
		{arr: []float32{1}, k: 0, expect: 0},
		{arr: []float32{1}, k: 1, expect: 0},
		{arr: []float32{1}, k: 2, expect: 1},
		{arr: []float32{}, k: 2, expect: 0},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, BinarySearchFloat32(c.arr, c.k))
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		arr    []float32
		expect float32
	}{
		{arr: []float32{1, 2, 3, 4}, expect: 10},
		{arr: []float32{1, 2, 3}, expect: 6},
		{arr: []float32{1}, expect: 1},
		{arr: []float32{0}, expect: 0},
		{arr: []float32{}, expect: 0},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, SumFloat32(c.arr))
	}
}
