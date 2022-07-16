package main

import "github.com/macduyhai/golang-design-pattern/Structural-pattern/adapter"

func main() {
	TestAdapter()
}

func TestAdapter() {
	var animals []adapter.Animal
	animals = append(animals, new(adapter.Cat))
	animals = append(animals, adapter.NewDog())
	for _, entity := range animals {
		entity.Move()
	}
}
