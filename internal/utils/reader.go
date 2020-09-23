package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadArffData(fileName string) [][]float32 {
	const Data = "@data"

	fin, _ := os.Open(fileName)
	file := bufio.NewReader(fin)

	isReadingData := false

	data := make([][]float32, 0, 1<<8)

	for line, err := file.ReadString('\n'); err == nil; line, err = file.ReadString('\n') {
		// Remove last \n
		line = line[:len(line)-1]
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
	}

	return data
}
