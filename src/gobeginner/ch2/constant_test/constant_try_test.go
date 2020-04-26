package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Firday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func Test1(t *testing.T) {
	t.Log(Monday, Tuesday)
	t.Log(7&Readable == Readable,
		7&Writeable == Writeable,
		7&Executable == Executable)
}
