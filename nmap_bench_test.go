package fmap

import (
	"sync"
	"testing"
)

const (
	getTotal = 1024
	setTotal = 1024
)

func GetSet(m NMap, finished chan struct{}) (set func(key, value int), get func(key, value int)) {
	return func(key, value int) {
			for i := uint32(0); i < getTotal; i++ {
				m.Get(i)
			}
			finished <- struct{}{}
		}, func(key, value int) {
			for i := uint32(0); i < setTotal; i++ {
				m.Set(i, i)
			}
			finished <- struct{}{}
		}
}

func benchmarkMultiGetSetDifferent(bucket int, b *testing.B) {
	m := NewNMap(bucket)
	finished := make(chan struct{}, 2*b.N)
	get, set := GetSet(m, finished)
	m.Set(1, 1)
	b.ResetTimer()
	//run N*2 goroutine (Get/Set KV(0-1000))
	for i := 0; i < b.N; i++ {
		go set(i, i)
		go get(i, i)
	}
	for i := 0; i < 2*b.N; i++ {
		<-finished
	}
}

func GetSetSyncMap(m *sync.Map, finished chan struct{}) (get func(key, value int), set func(key, value int)) {
	get = func(key, value int) {
		for i := 0; i < getTotal; i++ {
			m.Load(i)
		}
		finished <- struct{}{}
	}
	set = func(key, value int) {
		for i := 0; i < setTotal; i++ {
			m.Store(i, i)
		}
		finished <- struct{}{}
	}
	return
}

func BenchmarkMultiGetSetDifferentSyncMap(b *testing.B) {
	var m sync.Map
	finished := make(chan struct{}, 2*b.N)
	get, set := GetSetSyncMap(&m, finished)
	m.Store(-1, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go set(i, i)
		go get(i, i)
	}
	for i := 0; i < 2*b.N; i++ {
		<-finished
	}
}

func BenchmarkMultiGetSetDifferent_1(b *testing.B) {
	benchmarkMultiGetSetDifferent(1, b)
}

func BenchmarkMultiGetSetDifferent_16(b *testing.B) {
	benchmarkMultiGetSetDifferent(16, b)
}

func BenchmarkMultiGetSetDifferent_32(b *testing.B) {
	benchmarkMultiGetSetDifferent(32, b)
}

func BenchmarkMultiGetSetDifferent_64(b *testing.B) {
	benchmarkMultiGetSetDifferent(64, b)
}

func BenchmarkMultiGetSetDifferent_128(b *testing.B) {
	benchmarkMultiGetSetDifferent(128, b)
}

func BenchmarkMultiGetSetDifferent_256(b *testing.B) {
	benchmarkMultiGetSetDifferent(256, b)
}
