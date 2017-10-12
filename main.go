package main

import (
	"bytes"
	"fmt"
	_ "io"
	_ "os"
	"time"
)

func main() {
	start := time.Now()

	const count = 1000000

	for i := 1; i < count; i++ {
		_ = "test" + "test 2" + "test 3" + "test 4"
	}

	elapsed := time.Since(start)
	fmt.Printf("\n Program time: %s", elapsed)

	start = time.Now()
	var buffer bytes.Buffer
	for i := 1; i < count; i++ {
		buffer.WriteString("test")
		buffer.WriteString("test 2")
		buffer.WriteString("test 3")
		buffer.WriteString("test 4")
		_ = buffer.String()
		buffer.Reset()
	}

	elapsed = time.Since(start)
	fmt.Printf("\n Program time: %s", elapsed)
}
