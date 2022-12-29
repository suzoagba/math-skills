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
	args, _ := os.ReadFile(os.Args[1])
	numb := StrToFloat(string(args))
	fmt.Println("Average:", math.Round(average(numb)))
	fmt.Println("Mediana:", math.Round(median(numb)))
	fmt.Println("Variance:", math.Round(variance(numb)))
	fmt.Println("Standard Deviation:", math.Round(Deviation(numb)))
}
func StrToFloat(s string) []float64 {
	arr := strings.Split(s, "\n")
	res := make([]float64, len(arr)-1)
	for i := range res {
		res[i], _ = strconv.ParseFloat(arr[i], 64)
	}
	return res
}
func average(data []float64) float64 {
	var sum float64
	for _, j := range data {
		sum += j
	}
	return sum / float64(len(data))
}
func median(data []float64) float64 {
	sort.Float64s(data)
	var med float64
	if len(data) == 0 {
		return 0
	} else if len(data)%2 == 0 {
		med = (data[len(data)/2-1] + data[len(data)/2]) / 2
	} else {
		med = data[len(data)/2-1]
	}
	return med
}
func variance(data []float64) float64 {
	var sumVar float64
	for _, j := range data {
		sumVar += math.Pow(j, 2)
	}
	result := sumVar/float64(len(data)) - math.Pow(average(data), 2)
	return result
}
func Deviation(data []float64) float64 {
	return math.Sqrt(variance(data))
}
