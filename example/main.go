package main

import (
	"fmt"
	"time"

	fmap "github.com/faster-go/map"
)

func main() {
	// nmap := fmap.NewNMap(64)
	nmap := fmap.NewNMap(1)

	nmap.Set(1, 1)
	nmap.Set(2, 2)
	nmap.Set(3, 3)

	fmt.Println(nmap.Get(1))

	fmt.Println(nmap.Has(1))

	nmap.Range(func(key uint32, value interface{}) {
		fmt.Println("Range ", key, value)
	})

	nmap.RangeClean(func(key uint32, value interface{}) bool {
		if key > 1 {
			fmt.Println("Range clean ", key, value)
			return true
		}
		return false
	})
	fmt.Println(nmap.Pop(1))
	time.Sleep(10 * time.Second)
}
