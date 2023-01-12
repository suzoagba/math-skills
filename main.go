package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	line := string(data)
	str := strings.Split(line, "\n")

	var f []float64
	for _, data1 := range str {
		num, err := strconv.ParseFloat(data1, 64)
		if err != nil {
			fmt.Println(err)
		}
		f = append(f, num)
	}
	fmt.Println(count(f))
}
func count(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)
	sort.Float64s(dataCopy)
	var sum float64 = 0
	for i := 0; i < len(dataCopy); i++ {
		sum += dataCopy[i]
	}
	var average float64 = sum / float64(len(dataCopy))
	fmt.Println("Average:\n", math.Round(average))
	var median float64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}
	fmt.Println("Median:\n", math.Round(median))
	var sd float64
	for j := 0; j < l; j++ {
		sd += math.Pow(dataCopy[j]-average, 2)
		fmt.Println(sd)
	}
	fmt.Println("Variance:\n", math.Round(sd))
	sd = math.Sqrt(sd / 10)
	fmt.Println("Standard Deviation:")
	return math.Round(sd)
}
