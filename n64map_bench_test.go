package fmap

import (
	"sync"
	"testing"
)

const (
	getTotal64 = 1024
	setTotal64 = 1024
)

func GetSet64(m N64Map, finished chan struct{}) (set func(key, value int), get func(key, value int)) {
	return func(key, value int) {
			for i := uint64(0); i < getTotal; i++ {
				m.Get(i)
			}
			finished <- struct{}{}
		}, func(key, value int) {
			for i := uint64(0); i < setTotal; i++ {
				m.Set(i, i)
			}
			finished <- struct{}{}
		}
}

func benchmarkMultiGetSet64Different(bucket int, b *testing.B) {
	m := NewN64Map(bucket)
	finished := make(chan struct{}, 2*b.N)
	get, set := GetSet64(m, finished)
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

func GetSet64SyncMap(m *sync.Map, finished chan struct{}) (get func(key, value int), set func(key, value int)) {
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

func BenchmarkMultiGetSet64DifferentSyncMap(b *testing.B) {
	var m sync.Map
	finished := make(chan struct{}, 2*b.N)
	get, set := GetSet64SyncMap(&m, finished)
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

func BenchmarkMultiGetSet64Different_1(b *testing.B) {
	benchmarkMultiGetSet64Different(1, b)
}

func BenchmarkMultiGetSet64Different_16(b *testing.B) {
	benchmarkMultiGetSet64Different(16, b)
}

func BenchmarkMultiGetSet64Different_32(b *testing.B) {
	benchmarkMultiGetSet64Different(32, b)
}

func BenchmarkMultiGetSet64Different_64(b *testing.B) {
	benchmarkMultiGetSet64Different(64, b)
}

func BenchmarkMultiGetSet64Different_128(b *testing.B) {
	benchmarkMultiGetSet64Different(128, b)
}

func BenchmarkMultiGetSet64Different_256(b *testing.B) {
	benchmarkMultiGetSet64Different(256, b)
}
