package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return fmt.Sprintf("Hello World\n")
}

func TestClient(t *testing.T) {
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
