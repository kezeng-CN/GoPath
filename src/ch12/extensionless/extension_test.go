package extension_test

import "fmt"

import "testing"

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Print("...")
}

func (p *Pet) speakto(host string) {
	p.speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	var dog *Dog = new(Dog)
	dog.speakto("Chao")
}
