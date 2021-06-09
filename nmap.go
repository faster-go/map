package fmap

import (
	"sync"
)

const (
	defaultBucketSize = 128
)

type NMap []*internalNMap
type nmapType map[uint32]interface{}

type internalNMap struct {
	objs nmapType
	sync.RWMutex
}

// New create a new map
// goroutine safe and high performance
func NewNMap(bucket int) NMap {
	if bucket <= 0 {
		bucket = defaultBucketSize
	}
	m := make(NMap, bucket)
	for i := 0; i < bucket; i++ {
		m[i] = &internalNMap{objs: make(nmapType)}
	}
	return m
}

func (m NMap) getInternalNMap(key uint32) *internalNMap {
	return m[uint32(key)%uint32(len(m))]
}

// Set set the key and value to map
func (m NMap) Set(key uint32, value interface{}) {
	interMap := m.getInternalNMap(key)
	interMap.Lock()
	interMap.objs[key] = value
	interMap.Unlock()
}

// Get get value from map by key
func (m NMap) Get(key uint32) (interface{}, bool) {
	interMap := m.getInternalNMap(key)
	interMap.RLock()
	val, ok := interMap.objs[key]
	interMap.RUnlock()
	return val, ok
}

// Count returns the count of objs
func (m NMap) Count() int {
	count := 0
	for _, inMap := range m {
		inMap.RLock()
		count += len(inMap.objs)
		inMap.RUnlock()
	}
	return count
}

// Has check key exist or not
func (m NMap) Has(key uint32) bool {
	interMap := m.getInternalNMap(key)
	interMap.RLock()
	_, ok := interMap.objs[key]
	interMap.RUnlock()
	return ok
}

// Remove remove an object from the map
func (m NMap) Remove(key uint32) {
	interMap := m.getInternalNMap(key)
	interMap.Lock()
	delete(interMap.objs, key)
	interMap.Unlock()
}

// Pop pop an object from the map and returns it
func (m NMap) Pop(key uint32) (v interface{}, exists bool) {
	// Try to get interMap.
	interMap := m.getInternalNMap(key)
	interMap.Lock()
	v, exists = interMap.objs[key]
	delete(interMap.objs, key)
	interMap.Unlock()
	return v, exists
}

// Empty check map is empty or not
func (m NMap) Empty() bool {
	return m.Count() == 0
}

func (m NMap) Range(f func(key uint32, value interface{})) {
	for _, inMap := range m {
		inMap.Lock()
		for k, v := range inMap.objs {
			f(k, v)
		}
		inMap.Unlock()
	}
}

func (m NMap) RangeClean(check func(key uint32, value interface{}) bool) {
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
