# K-Means++ Implementation in Golang

Reference:
- Arthur, D.; Vassilvitskii, S. (2007). "k-means++: the advantages of careful seeding". Proceedings of the eighteenth annual ACM-SIAM symposium on Discrete algorithms. Society for Industrial and Applied Mathematics Philadelphia, PA, USA. pp. 1027–1035.

# User Guide

## Main Entrance

K-Means entrance function is 

```go
func KMeansPP(k, round int, data [][]float32) []int
```

Example can be found in `kmeanspp_test.go`

**Attention**

User should ensure every element in data array has same length.

## Clustering Result Evaluation

Functions for evaluation can be found under evaluate package. Mainly contains:

```go
// Calculate the total of the distance between all instance data and its class center.
func ClassCenterDistanceSquareTotal(data [][]float32, class []int, k int) float32

func HopkinsStatistic(data [][]float32) float32
```

