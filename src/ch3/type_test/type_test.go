package type_test

import (
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt = 1
	b = int64(c)
	t.Log(a, b, c)
	if math.MaxInt64 > b {
		t.Log(b)
	}
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*"+s+"*", "len", len(s))
}
