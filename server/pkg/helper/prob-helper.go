package helper

import (
	"math"
	"sort"
)

type ProbHelper struct{}

func NewProbHelper() *ProbHelper {
	return &ProbHelper{}
}

func (p *ProbHelper) Mean(data []int) float64 {
	sum := 0

	for _, value := range data {
		sum += value
	}

	return float64(sum) / float64(len(data))
}

func (p *ProbHelper) Variance(data []int) float64 {
	mean := int(p.Mean(data))
	sum := 0

	for _, value := range data {
		diff := value - mean
		sum += diff * diff
	}
	return float64(sum) / float64(len(data)-1)
}

func (p *ProbHelper) StdDeviation(data []int) float64 {
	return math.Sqrt(p.Variance(data))
}

func (p *ProbHelper) Median(data []int) int {
	length := float64(len(data))

	sort.Ints(data)

	if length == 0 {
		return 0
	}

	middle := int(length / 2)

	if len(data)%2 == 1 {
		return data[middle]
	}

	return int((float64(data[middle-1] + data[middle])) / 2)
}
