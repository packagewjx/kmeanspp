package kmeanspp

import (
	"bufio"
	"fmt"
	"github.com/packagewjx/kmeanspp/evaluate"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestKmeansppEmotionDataset(t *testing.T) {
	dataset := readArffData("test/dataset/emotions.arff")

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
	data := readArffData("test/dataset/emotions.arff")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMeansPP(6, 30, data)
	}
}

func BenchmarkKMeansPPSceneDataset(b *testing.B) {
	data := readArffData("test/dataset/scene.arff")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMeansPP(6, 30, data)
	}
}

func readArffData(fileName string) [][]float32 {
	const Data = "@data"

	fin, _ := os.Open(fileName)
	file := bufio.NewReader(fin)

	buf, prefix, err := file.ReadLine()
	isReadingData := false

	data := make([][]float32, 0, 1<<8)

	for err == nil {
		line := string(buf)
		for prefix && err == nil {
			buf, prefix, err = file.ReadLine()
			line += string(buf)
		}

		if len(line) == 0 || line[0] == '%' {

		} else if isReadingData {
			split := strings.Split(line, ",")
			lineData := make([]float32, len(split))
			for i, s := range split {
				f, _ := strconv.ParseFloat(s, 32)
				lineData[i] = float32(f)
			}
			data = append(data, lineData)
		} else {
			if strings.EqualFold(line[:len(Data)], Data) {
				isReadingData = true
			}
		}

		buf, prefix, err = file.ReadLine()
	}

	return data
}
