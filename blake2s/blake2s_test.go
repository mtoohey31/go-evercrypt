// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blake2s

import (
	"fmt"
	"testing"
)

func TestHashes(t *testing.T) {
	input := make([]byte, 255)
	for i := range input {
		input[i] = byte(i)
	}

	for i, expectedHex := range hashes {
		h := New256()

		h.Write(input[:i])
		sum := h.Sum(nil)

		if gotHex := fmt.Sprintf("%x", sum); gotHex != expectedHex {
			t.Fatalf("#%d (single write): got %s, wanted %s", i, gotHex, expectedHex)
		}

		h.Reset()
		for j := 0; j < i; j++ {
			h.Write(input[j : j+1])
		}

		sum = h.Sum(sum[:0])
		if gotHex := fmt.Sprintf("%x", sum); gotHex != expectedHex {
			t.Fatalf("#%d (byte-by-byte): got %s, wanted %s", i, gotHex, expectedHex)
		}
	}
}

// Benchmarks

func benchmarkSum(b *testing.B, size int) {
	data := make([]byte, size)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum256(data)
	}
}

func benchmarkWrite(b *testing.B, size int) {
	data := make([]byte, size)
	h := New256()
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Write(data)
	}
}

func BenchmarkWrite64(b *testing.B) { benchmarkWrite(b, 64) }
func BenchmarkWrite1K(b *testing.B) { benchmarkWrite(b, 1024) }

func BenchmarkSum64(b *testing.B) { benchmarkSum(b, 64) }
func BenchmarkSum1K(b *testing.B) { benchmarkSum(b, 1024) }

var hashes = []string{
	"69217a3079908094e11121d042354a7c1f55b6482ca1a51e1b250dfd1ed0eef9",
	"e34d74dbaf4ff4c6abd871cc220451d2ea2648846c7757fbaac82fe51ad64bea",
	"ddad9ab15dac4549ba42f49d262496bef6c0bae1dd342a8808f8ea267c6e210c",
	"e8f91c6ef232a041452ab0e149070cdd7dd1769e75b3a5921be37876c45c9900",
	"0cc70e00348b86ba2944d0c32038b25c55584f90df2304f55fa332af5fb01e20",
}
