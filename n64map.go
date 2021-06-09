package fmap

import (
	"sync"
)

type N64Map []*internalN64Map
type n64mapType map[uint64]interface{}

type internalN64Map struct {
	objs n64mapType
	sync.RWMutex
}

// New create a new map
// goroutine safe and high performance
func NewN64Map(bucket int) N64Map {
	if bucket <= 0 {
		bucket = defaultBucketSize
	}
	m := make(N64Map, bucket)
	for i := 0; i < bucket; i++ {
		m[i] = &internalN64Map{objs: make(n64mapType)}
	}
	return m
}

func (m N64Map) getInternalN64Map(key uint64) *internalN64Map {
	return m[uint64(key)%uint64(len(m))]
}

// Set set the key and value to map
func (m N64Map) Set(key uint64, value interface{}) {
	interMap := m.getInternalN64Map(key)
	interMap.Lock()
	interMap.objs[key] = value
	interMap.Unlock()
}

// Get get value from map by key
func (m N64Map) Get(key uint64) (interface{}, bool) {
	interMap := m.getInternalN64Map(key)
	interMap.RLock()
	val, ok := interMap.objs[key]
	interMap.RUnlock()
	return val, ok
}

// Count returns the count of objs
func (m N64Map) Count() int {
	count := 0
	for i := 0; i < len(m); i++ {
		interMap := m[i]
		interMap.RLock()
		count += len(interMap.objs)
		interMap.RUnlock()
	}
	return count
}

// Has check key exist or not
func (m N64Map) Has(key uint64) bool {
	interMap := m.getInternalN64Map(key)
	interMap.RLock()
	_, ok := interMap.objs[key]
	interMap.RUnlock()
	return ok
}

// Remove remove an object from the map
func (m N64Map) Remove(key uint64) {
	interMap := m.getInternalN64Map(key)
	interMap.Lock()
	delete(interMap.objs, key)
	interMap.Unlock()
}

// Pop pop an object from the map and returns it
func (m N64Map) Pop(key uint64) (v interface{}, exists bool) {
	// Try to get interMap.
	interMap := m.getInternalN64Map(key)
	interMap.Lock()
	v, exists = interMap.objs[key]
	delete(interMap.objs, key)
	interMap.Unlock()
	return v, exists
}

// Empty check map is empty or not
func (m N64Map) Empty() bool {
	return m.Count() == 0
}

func (m N64Map) Range(f func(key uint64, value interface{})) {
	for _, inMap := range m {
		inMap.Lock()
		for k, v := range inMap.objs {
			f(k, v)
		}
		inMap.Unlock()
	}
}

func (m N64Map) RangeClean(check func(key uint64, value interface{}) bool) {
	for _, inMap := range m {
		inMap.Lock()
		for k, v := range inMap.objs {
			if check(k, v) {
				delete(inMap.objs, k)
			}
		}
		inMap.Unlock()
	}
}
