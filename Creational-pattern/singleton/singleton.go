package singleton

import (
	"fmt"
	"sync"
)

var (
	instance *singleton
	once     sync.Once
)

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

func NewInstance() Singleton {
	once.Do(func() {
		fmt.Println("Init Intance")
		instance = &singleton{count: 1}
	})
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
