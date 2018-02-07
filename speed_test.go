package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.Write([]byte("x"))
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

func BenchmarkAppend(b *testing.B) {
	bs := make([]byte, 0)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bs = append(bs, []byte("x")...)
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, 0)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bs = append(bs, []byte("x")...)
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}

//go test -bench=. -benchtime=100ms -v
