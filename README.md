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

quickSort for ./nums_1000:     total = 0.01058s;   average = 0.00529s
quickSort for ./nums_10000:    total = 0.11735s;   average = 0.05868s
quickSort for ./nums_100000:   total = 1.21976s;   average = 0.60988s
quickSort for ./nums_1000000:  total = 14.05871s;  average = 7.02936s
quickSort for ./nums_10000000: total = 168.24321s; average = 84.12161s

// PHP native sort function
native sort for ./nums_1000:     total = 0.00371s;  average = 0.00186s
native sort for ./nums_10000:    total = 0.03739s;  average = 0.0187s
native sort for ./nums_100000:   total = 0.30923s;  average = 0.15462s
native sort for ./nums_1000000:  total = 3.11071s;  average = 1.55536s
native sort for ./nums_10000000: total = 32.87943s; average = 16.43972s
```