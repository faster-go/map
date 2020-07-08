package fmap

import (
	"sync"
	"testing"
)

const (
	getTotal = 1024
	setTotal = 1024
)

func GetSet(m IntMap, finished chan struct{}) (set func(key, value int), get func(key, value int)) {
	return func(key, value int) {
			// r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < getTotal; i++ {
				// m.Get(r.Intn(getTotal))
				m.Get(i)
			}
			finished <- struct{}{}
		}, func(key, value int) {
			// r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < setTotal; i++ {
				// n := r.Intn(setTotal)
				// m.Set(n, n)
				m.Set(i, i)
			}
			finished <- struct{}{}
		}
}

func benchmarkMultiGetSetDifferent(bucket int, b *testing.B) {
	m := NewIntMap(bucket)
	finished := make(chan struct{}, 2*b.N)
	get, set := GetSet(m, finished)
	m.Set(-1, -1)
	b.ResetTimer()
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
		// r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < getTotal; i++ {
			// m.Load(r.Intn(getTotal))
			m.Load(i)
		}
		finished <- struct{}{}
	}
	set = func(key, value int) {
		// r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < setTotal; i++ {
			// n := r.Intn(setTotal)
			// m.Store(n, n)
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
