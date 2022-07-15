package singleton

import (
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
		instance = &singleton{count: 1}
	})
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
