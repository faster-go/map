package fmap

import (
	"sync"
)

const (
	defaultBucketSize = 128
)

var (
	bucketSize int
)

type IntMap []*internalIntMap
type intKeyMap map[int]interface{}

type internalIntMap struct {
	objs intKeyMap
	sync.RWMutex
}

// New create a new map
// goroutine safe and high performance
// key is int and value is interface
func NewIntMap(bucket int) IntMap {
	if bucket <= 0 {
		bucket = defaultBucketSize
	}
	bucketSize = bucket
	m := make(IntMap, bucketSize)
	for i := 0; i < bucketSize; i++ {
		m[i] = &internalIntMap{objs: make(intKeyMap)}
	}
	return m
}

func (m IntMap) getInternalMap(key int) *internalIntMap {
	return m[uint(key)%uint(bucketSize)]
}

// Set set the key and value to map
func (m IntMap) Set(key int, value interface{}) {
	interMap := m.getInternalMap(key)
	interMap.Lock()
	interMap.objs[key] = value
	interMap.Unlock()
}

// Get get value from map by key
func (m IntMap) Get(key int) (interface{}, bool) {
	interMap := m.getInternalMap(key)
	interMap.RLock()
	val, ok := interMap.objs[key]
	interMap.RUnlock()
	return val, ok
}

// Count returns the count of objs
func (m IntMap) Count() int {
	count := 0
	for i := 0; i < defaultBucketSize; i++ {
		interMap := m[i]
		interMap.RLock()
		count += len(interMap.objs)
		interMap.RUnlock()
	}
	return count
}

// Has check key exist or not
func (m IntMap) Has(key int) bool {
	interMap := m.getInternalMap(key)
	interMap.RLock()
	_, ok := interMap.objs[key]
	interMap.RUnlock()
	return ok
}

// Remove remove an object from the map
func (m IntMap) Remove(key int) {
	interMap := m.getInternalMap(key)
	interMap.Lock()
	delete(interMap.objs, key)
	interMap.Unlock()
}

// Pop pop an object from the map and returns it
func (m IntMap) Pop(key int) (v interface{}, exists bool) {
	// Try to get interMap.
	interMap := m.getInternalMap(key)
	interMap.Lock()
	v, exists = interMap.objs[key]
	delete(interMap.objs, key)
	interMap.Unlock()
	return v, exists
}

// Empty check map is empty or not
func (m IntMap) Empty() bool {
	return m.Count() == 0
}
