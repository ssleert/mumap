package mumap

import (
	"sync"

	"golang.org/x/exp/constraints"
)

const (
	defaultMapSize = 2048
)

type Map[T constraints.Ordered, Y any] struct {
	mu sync.RWMutex
	hashMap map[T]Y
}

func New[T constraints.Ordered, Y any](mapSize int) Map {
	if mapSize <= 0 {
		mapSize = defaultMapSize
	}

	return Map{
		hashMap: make(map[T]Y, mapSize)
	}
}

func (m *Map[T, Y]) Get(T key) Y {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.hashMap[key]
}

func (m *Map[T, Y]) Set(T key, Y val) {	
	m.mu.Lock()
	m.hashMap[key] = val
	m.mu.Unlock()
}

func (m *Map[T, Y]) Del(T key) {
	m.mu.Lock()
	delete(m.hashMap, key)
	m.mu.Unlock()
}

func (m *Map[T, Y]) Len() int {
	return len(m.hashMap)
}
