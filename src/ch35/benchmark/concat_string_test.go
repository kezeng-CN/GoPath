package benchmark_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := [...]string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := [...]string{"1", "2", "3", "4", "5"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
