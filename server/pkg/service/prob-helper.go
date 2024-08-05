package service

import "math"

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
