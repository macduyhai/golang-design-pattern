package main

import (
	"log"

	"github.com/macduyhai/golang-design-pattern/Structural-pattern/adapter"
	"github.com/macduyhai/golang-design-pattern/Structural-pattern/proxy"
)

func main() {
	TestAdapter()
	TestProxy()
}
func TestProxy() {
	// proxy pattern Structural pattern
	log.Println("proxy pattern of Structural ")
	highResolutionImage := proxy.NewImageProxy("sample/img1.png")
	// the realImage won't be loaded until user calls display
	// later
	highResolutionImage.Display()
}
func TestAdapter() {
	var animals []adapter.Animal
	animals = append(animals, new(adapter.Cat))
	animals = append(animals, adapter.NewDog())
	for _, entity := range animals {
		entity.Move()
	}
}
