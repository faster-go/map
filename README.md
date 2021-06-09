# Fmap
Fmap is a faster map, better than sync.Map, easy to use as origin map
Memory usage is similar to origin map

# Why fast
It using a fragment-map arch
Get/Set api use hash(key%N) to find a internal map
This can avoid lots of lock race
```
============fmap=============
      Get(K)/Set(K,V)
        index=K%N
          |
          +------+
                 |
[  0  ][  1  ][...  ][  N  ]
   |      |      |      |
   |      |      |      |
   |      |      |      |
+--+---+--+---+--+--+---+---+
| map0 | map1 |...  | mapN  |
|lock0 |lock1 |...  |lockN  |
+------+------+-----+-------+
```

# benchmark
This show: it's faster than sync.Map when setting bucket >= 16
```
goos: darwin
goarch: amd64
pkg: github.com/faster-go/map
BenchmarkMultiGetSet64DifferentSyncMap
BenchmarkMultiGetSet64DifferentSyncMap-4   	   27722	    123670 ns/op
BenchmarkMultiGetSet64Different_1
BenchmarkMultiGetSet64Different_1-4        	    7926	    428635 ns/op
BenchmarkMultiGetSet64Different_16
BenchmarkMultiGetSet64Different_16-4       	   33333	    106161 ns/op
BenchmarkMultiGetSet64Different_32
BenchmarkMultiGetSet64Different_32-4       	   38814	     92811 ns/op
BenchmarkMultiGetSet64Different_64
BenchmarkMultiGetSet64Different_64-4       	   44708	     80400 ns/op
BenchmarkMultiGetSet64Different_128
BenchmarkMultiGetSet64Different_128-4      	   52779	     74070 ns/op
BenchmarkMultiGetSet64Different_256
BenchmarkMultiGetSet64Different_256-4      	   61125	     61604 ns/op
BenchmarkMultiGetSetDifferentSyncMap
BenchmarkMultiGetSetDifferentSyncMap-4     	   33190	    114794 ns/op
BenchmarkMultiGetSetDifferent_1
BenchmarkMultiGetSetDifferent_1-4          	    9002	    437252 ns/op
BenchmarkMultiGetSetDifferent_16
BenchmarkMultiGetSetDifferent_16-4         	   36232	     96233 ns/op
BenchmarkMultiGetSetDifferent_32
BenchmarkMultiGetSetDifferent_32-4         	   43854	     91238 ns/op
BenchmarkMultiGetSetDifferent_64
BenchmarkMultiGetSetDifferent_64-4         	   47797	     78608 ns/op
BenchmarkMultiGetSetDifferent_128
BenchmarkMultiGetSetDifferent_128-4        	   50012	     69572 ns/op
BenchmarkMultiGetSetDifferent_256
BenchmarkMultiGetSetDifferent_256-4        	   66978	     54394 ns/op
PASS
ok  	github.com/faster-go/map	66.117s
```
