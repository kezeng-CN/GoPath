package pipefilter

import (
	"errors"
	"strconv"
)

// ErrToIntFilterWrongFormat is
var ErrToIntFilterWrongFormat = errors.New("input data should be []string")

// ToIntFilter is
type ToIntFilter struct {
}

// NewToIntFilter is
func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

// Process is
func (pf *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ErrToIntFilterWrongFormat
	}
	ret := []int{}
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}
