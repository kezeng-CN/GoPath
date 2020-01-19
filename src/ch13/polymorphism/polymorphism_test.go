package polymorphism_test

import "testing"

import "fmt"

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgrammer(goProg)
	writeFirstProgrammer(javaProg)
}
