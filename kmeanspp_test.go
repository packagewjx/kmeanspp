package kmeanspp

import (
	"fmt"
	"github.com/packagewjx/kmeanspp/evaluate"
	"github.com/packagewjx/kmeanspp/internal/utils"
	"testing"
)

func TestKmeansppEmotionDataset(t *testing.T) {
	dataset := utils.ReadArffData("test/dataset/emotions.arff")

	_, class := KMeansPP(1, 30, dataset)
	last := evaluate.ClassCenterDistanceSquareTotal(dataset, class, 1)
	fmt.Printf("K=%d dist=%.2f\n", 1, last)

	for i := 2; i <= 30; i++ {
		_, class = KMeansPP(i, 30, dataset)
		dist := evaluate.ClassCenterDistanceSquareTotal(dataset, class, i)
		fmt.Printf("K=%d dist=%.2f diff%%=%.2f%%\n", i, dist, (last-dist)/last*100)
		last = dist
	}
}

func BenchmarkKMeansPPEmotionsDataset(b *testing.B) {
	data := utils.ReadArffData("test/dataset/emotions.arff")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMeansPP(6, 30, data)
	}
}

func BenchmarkKMeansPPSceneDataset(b *testing.B) {
	data := utils.ReadArffData("test/dataset/scene.arff")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMeansPP(6, 30, data)
	}
}
