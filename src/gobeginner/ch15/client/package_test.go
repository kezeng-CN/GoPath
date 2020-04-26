package client_test

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFabonacciSerie(8))
	t.Log(series.Square(1))
}
