```
go test -bench . *.go
```

```
php sort.php
```

```
// Golang quickSort

goos: darwin
goarch: amd64
BenchmarkQuickSort1000-4       	   19274	     73380 ns/op
BenchmarkQuickSort10000-4      	    2436	    478411 ns/op
BenchmarkQuickSort100000-4     	     222	   5187577 ns/op
BenchmarkQuickSort1000000-4    	      22	  53431253 ns/op
BenchmarkQuickSort10000000-4   	       2	 574170556 ns/op
```

```
// PHP quickSort

./nums_1000:     total = 0.01116s;   average = 0.00558s
./nums_10000:    total = 0.11796s;   average = 0.05898s
./nums_100000:   total = 1.23729s;   average = 0.61865s
./nums_1000000:  total = 16.03092s;  average = 8.01546s
./nums_10000000: total = 169.25291s; average = 84.62646s
```