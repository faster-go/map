package fmap

import (
	"testing"
)

func TestSetGet64(t *testing.T) {
	nmap := NewN64Map(64)

	for i := uint64(0); i < 10; i++ {
		nmap.Set(i, i)
	}

	for i := uint64(0); i < 10; i++ {
		if v, ok := nmap.Get(i); !ok || v.(uint64) != i {
			t.Errorf("should have k=%v v=%v\n", i, v)
		}
	}
}

func TestHas64(t *testing.T) {
	nmap := NewN64Map(64)

	for i := uint64(0); i < 10; i++ {
		nmap.Set(i, i)
	}

	for i := uint64(0); i < 10; i++ {
		if ok := nmap.Has(i); !ok {
			t.Errorf("should have k=%v ", i)
		}
	}
}

func TestPop64(t *testing.T) {
	nmap := NewN64Map(64)

	for i := uint64(0); i < 10; i++ {
		nmap.Set(i, i)
	}
	// 9 is in map
	if n, ok := nmap.Pop(9); !ok || n.(uint64) != 9 {
		t.Errorf("should pop 9 ok")
	}

	// 10 is not in map
	if n, ok := nmap.Pop(10); ok || n == 10 {
		t.Errorf("should pop 9 ok")
	}

	for i := uint64(0); i < uint64(nmap.Count()); i++ {
		if v, ok := nmap.Get(i); !ok || v.(uint64) != i {
			t.Errorf("should have k=%v v=%v\n", i, v)
		}
	}
}

func TestRange64(t *testing.T) {
	nmap := NewN64Map(64)

	n, total := uint64(0), uint64(0)
	for i := uint64(0); i < 10; i++ {
		nmap.Set(i, i)
		total += i
	}

	nmap.Range(func(key uint64, value interface{}) {
		n += value.(uint64)
	})

	if n != total {
		t.Errorf("n should equal total")
	}

	nmap.RangeClean(func(key uint64, value interface{}) bool {
		return true
	})

	if nmap.Count() != 0 {
		t.Errorf("map should be cleaned")
	}
}
