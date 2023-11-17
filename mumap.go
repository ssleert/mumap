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

func New[T constraints.Ordered, Y any](mapSize int) Map[T, Y] {
	if mapSize <= 0 {
		mapSize = defaultMapSize
	}

	return Map[T, Y]{
		hashMap: make(map[T]Y, mapSize),
	}
}

func (m *Map[T, Y]) Get(key T) (Y, bool) {
	m.mu.RLock()
	v, ok := m.hashMap[key]
	m.mu.RUnlock()
	return v, ok
}

func (m *Map[T, Y]) Set(key T, val Y) {	
	m.mu.Lock()
	m.hashMap[key] = val
	m.mu.Unlock()
}

func (m *Map[T, Y]) Del(key T) {
	m.mu.Lock()
	delete(m.hashMap, key)
	m.mu.Unlock()
}

func (m *Map[T, Y]) Len() int {
	return len(m.hashMap)
}
