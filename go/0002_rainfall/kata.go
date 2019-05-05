// https://www.codewars.com/kata/rainfall/train/go
package kata

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

func Mean(town string, strng string) float64 {
	rainfalls := getRainfalls(town, strng)
	if len(rainfalls) == 0 {
		return -1.0
	}
	sum := 0.0
	for _, rainfall := range rainfalls {
		sum += rainfall
	}
	return sum / float64(len(rainfalls))

}
func Variance(town string, strng string) float64 {
	rainfalls := getRainfalls(town, strng)
	if len(rainfalls) == 0 {
		return -1.0
	}
	mean := Mean(town, strng)
	if mean == -1.0 {
		return -1.0
	}
	sum := 0.0
	for _, rainfall := range rainfalls {
		diff := rainfall - mean
		sum += diff * diff
	}
	return sum / float64(len(rainfalls))
}

func getRainfalls(town string, str string) []float64 {
	rainfalls := make([]float64, 0)
	r := csv.NewReader(strings.NewReader(str))
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		strTown := strings.Split(records[0], ":")[0]
		if strTown != town {
			continue
		}

		// handle case when we found searched town
		for _, record := range records {
			rainfall, err := strconv.ParseFloat(strings.Split(record, " ")[1], 64)
			if err != nil {
				panic(err)
			}
			rainfalls = append(rainfalls, rainfall)
		}
		break
	}

	return rainfalls
}
