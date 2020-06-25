package pipefilter

import (
	"errors"
	"strings"
)

// ErrSplitFilterWrongFormat is
var ErrSplitFilterWrongFormat = errors.New("input data should be string")

// SplitFilter is
type SplitFilter struct {
	delimiter string
}

// NewSplitFilter is
func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

// Process is
func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, ErrSplitFilterWrongFormat
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
