package pipefilter

import "errors"

// ErrSumFilterWrongFormat is
var ErrSumFilterWrongFormat = errors.New("input should be")

// SumFilter is
type SumFilter struct {
}

// NewSumFilter is
func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

// Process is
func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, ErrSumFilterWrongFormat
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
