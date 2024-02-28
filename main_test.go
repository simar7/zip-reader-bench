package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"testing"

	iox "github.com/aquasecurity/trivy/pkg/x/io"
)

func FuncWithReadAll(b *testing.B) {
	f, err := os.Open("test.zip")
	if err != nil {
		b.Fatal(err)
	}

	buf, err := io.ReadAll(f)
	if err != nil {
		b.Fatal(err)
	}

	_, err = zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		b.Fatal(err)
	}
}

func FuncWithDiscard(b *testing.B) {
	f, err := os.Open("test.zip")
	if err != nil {
		b.Fatal(err)
	}

	size, err := io.Copy(io.Discard, f)
	if err != nil {
		b.Fatal(err)
	}

	bb, err := iox.NewReadSeekerAt(f)
	if err != nil {
		b.Fatal(err)
	}

	_, err = zip.NewReader(bb, size)
	if err != nil {
		b.Fatal(err)
	}
}

// old
//func BenchmarkNormal(b *testing.B) {
//	b.ReportAllocs()
//	for n := 0; n < b.N; n++ {
//		FuncWithReadAll(b)
//	}
//}

// new
func BenchmarkNormal(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		FuncWithDiscard(b)
	}
}
