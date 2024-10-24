// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blake2b

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
		h := New512()

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
		Sum512(data)
	}
}

func benchmarkWrite(b *testing.B, size int) {
	data := make([]byte, size)
	h := New512()
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Write(data)
	}
}

func BenchmarkWrite128(b *testing.B) { benchmarkWrite(b, 128) }
func BenchmarkWrite1K(b *testing.B)  { benchmarkWrite(b, 1024) }

func BenchmarkSum128(b *testing.B) { benchmarkSum(b, 128) }
func BenchmarkSum1K(b *testing.B)  { benchmarkSum(b, 1024) }

var hashes = []string{
	"786a02f742015903c6c6fd852552d272912f4740e15847618a86e217f71f5419d25e1031afee585313896444934eb04b903a685b1448b755d56f701afe9be2ce",
	"2fa3f686df876995167e7c2e5d74c4c7b6e48f8068fe0e44208344d480f7904c36963e44115fe3eb2a3ac8694c28bcb4f5a0f3276f2e79487d8219057a506e4b",
	"1c08798dc641aba9dee435e22519a4729a09b2bfe0ff00ef2dcd8ed6f8a07d15eaf4aee52bbf18ab5608a6190f70b90486c8a7d4873710b1115d3debbb4327b5",
	"40a374727302d9a4769c17b5f409ff32f58aa24ff122d7603e4fda1509e919d4107a52c57570a6d94e50967aea573b11f86f473f537565c66f7039830a85d186",
	"77ddf4b14425eb3d053c1e84e3469d92c4cd910ed20f92035e0c99d8a7a86cecaf69f9663c20a7aa230bc82f60d22fb4a00b09d3eb8fc65ef547fe63c8d3ddce",
}
