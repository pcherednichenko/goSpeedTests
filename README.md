## Go speed tests for combining strings

How to run the test:
```
go test -bench=. -benchtime=1000ms -v
```

Result:
```
goos: darwin
goarch: amd64
pkg: github.com/pcherednichenko/goSpeedTests
BenchmarkConcat-8   	 1000000	     52710 ns/op
BenchmarkBuffer-8   	100000000	        18.2 ns/op
BenchmarkAppend-8   	100000000	        13.3 ns/op
BenchmarkCopy-8     	100000000	        12.2 ns/op
PASS
ok  	github.com/pcherednichenko/goSpeedTests	57.548s
```