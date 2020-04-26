package pipefilter

import (
	"reflect"
	"testing"
)

func TestStraigtPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	convert := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("pl", spliter, convert, sum)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}

func TestStraigtSplitFilter(t *testing.T) {
	spliter := NewSplitFilter(",")
	sp := NewStraightPipeline("pl", spliter)
	ret, err := sp.Process("1,2,3")
	if !reflect.DeepEqual([]string{"1", "2", "3"}, ret) {
		t.Fatal(err)
	}
}

func TestSumFilter(t *testing.T) {
	sum := NewSumFilter()
	sp := NewStraightPipeline("pl", sum)
	ret, err := sp.Process([]int{1, 2, 3})
	if 6 != ret {
		t.Fatal(err)
	}
}

func TestToIntFilter(t *testing.T) {
	convert := NewToIntFilter()
	sp := NewStraightPipeline("pl", convert)
	ret, err := sp.Process([]string{"1", "2", "3"})
	if !reflect.DeepEqual([]int{1, 2, 3}, ret) {
		t.Fatal(err)
	}
}
