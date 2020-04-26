package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	// c := [...]int{1, 2, 3, 4, 6}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writeable
	Executeable
)

func TestBitClear(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeable == Executeable)
	a = a &^ Readable
	t.Log(a&Readable == Readable)
	a = a &^ Writeable
	t.Log(a&Writeable == Writeable)
	a = a &^ Executeable
	t.Log(a&Executeable == Executeable)
}
