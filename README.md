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
BenchmarkConcat-8   	 1000000	     47336 ns/op
BenchmarkBuffer-8   	100000000	        15.9 ns/op
BenchmarkAppend-8   	100000000	        12.5 ns/op
BenchmarkCopy-8     	500000000	         3.68 ns/op
PASS
ok  	github.com/pcherednichenko/goSpeedTests	53.750s
```