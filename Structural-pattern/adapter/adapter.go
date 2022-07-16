package adapter

import "log"

// Cho phép các interface ko liên quan đến nhau có thể làm việc cùng nhau
type Animal interface {
	Move()
}

type Cat struct{}

func (c *Cat) Move() {
	log.Println("run run meo meo")
}

type Dog struct{}

func (d *Dog) Slither() {
	log.Println("run gau gau")
}

type DogAdapter struct {
	*Dog
}

func NewDog() *DogAdapter {
	return &DogAdapter{new(Dog)}
}
func (d *DogAdapter) Move() {
	d.Slither()
}
