package fmap

import (
	"testing"
)

func TestSetGet(t *testing.T) {
	nmap := NewNMap(64)

	for i := uint32(0); i < 10; i++ {
		nmap.Set(i, i)
	}

	for i := uint32(0); i < 10; i++ {
		if v, ok := nmap.Get(i); !ok || v.(uint32) != i {
			t.Errorf("should have k=%v v=%v\n", i, v)
		}
	}
}

func TestHas(t *testing.T) {
	nmap := NewNMap(64)

	for i := uint32(0); i < 10; i++ {
		nmap.Set(i, i)
	}

	for i := uint32(0); i < 10; i++ {
		if ok := nmap.Has(i); !ok {
			t.Errorf("should have k=%v ", i)
		}
	}
}

func TestPop(t *testing.T) {
	nmap := NewNMap(64)

	for i := uint32(0); i < 10; i++ {
		nmap.Set(i, i)
	}
	// 9 is in map
	if n, ok := nmap.Pop(9); !ok || n.(uint32) != 9 {
		t.Errorf("should pop 9 ok")
	}

	// 10 is not in map
	if n, ok := nmap.Pop(10); ok || n == 10 {
		t.Errorf("should pop 9 ok")
	}

	for i := uint32(0); i < uint32(nmap.Count()); i++ {
		if v, ok := nmap.Get(i); !ok || v.(uint32) != i {
			t.Errorf("should have k=%v v=%v\n", i, v)
		}
	}
}

func TestRange(t *testing.T) {
	nmap := NewNMap(64)

	n, total := uint32(0), uint32(0)
	for i := uint32(0); i < 10; i++ {
		nmap.Set(i, i)
		total += i
	}

	nmap.Range(func(key uint32, value interface{}) {
		n += value.(uint32)
	})

	if n != total {
		t.Errorf("n should equal total")
	}

	nmap.RangeClean(func(key uint32, value interface{}) bool {
		return true
	})

	if nmap.Count() != 0 {
		t.Errorf("map should be cleaned")
	}
}
