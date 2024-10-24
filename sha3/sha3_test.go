// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3

// Tests include all the ShortMsgKATs provided by the Keccak team at
// https://github.com/gvanas/KeccakCodePackage
//
// They only include the zero-bit case of the bitwise testvectors
// published by NIST in the draft of FIPS-202.

import (
	"bytes"
	"compress/flate"
	"encoding/hex"
	"encoding/json"
	"hash"
	"math/rand"
	"os"
	"strings"
	"testing"
)

const (
	testString  = "brekeccakkeccak koax koax"
	katFilename = "testdata/keccakKats.json.deflate"
)

// testDigests contains functions returning hash.Hash instances
// with output-length equal to the KAT length for SHA-3, Keccak
// and SHAKE instances.
var testDigests = map[string]func() hash.Hash{
	"SHA3-224": New224,
	"SHA3-256": New256,
	"SHA3-384": New384,
	"SHA3-512": New512,
}

// decodeHex converts a hex-encoded string into a raw byte string.
func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// structs used to marshal JSON test-cases.
type KeccakKats struct {
	Kats map[string][]struct {
		Digest  string `json:"digest"`
		Length  int64  `json:"length"`
		Message string `json:"message"`

		// Defined only for cSHAKE
		N string `json:"N"`
		S string `json:"S"`
	}
}

// TestKeccakKats tests the SHA-3 and Shake implementations against all the
// ShortMsgKATs from https://github.com/gvanas/KeccakCodePackage
// (The testvectors are stored in keccakKats.json.deflate due to their length.)
func TestKeccakKats(t *testing.T) {
	// Read the KATs.
	deflated, err := os.Open(katFilename)
	if err != nil {
		t.Errorf("error opening %s: %s", katFilename, err)
	}
	file := flate.NewReader(deflated)
	dec := json.NewDecoder(file)
	var katSet KeccakKats
	err = dec.Decode(&katSet)
	if err != nil {
		t.Errorf("error decoding KATs: %s", err)
	}

	for algo, function := range testDigests {
		d := function()
		for _, kat := range katSet.Kats[algo] {
			d.Reset()
			in, err := hex.DecodeString(kat.Message)
			if err != nil {
				t.Errorf("error decoding KAT: %s", err)
			}
			d.Write(in[:kat.Length/8])
			got := strings.ToUpper(hex.EncodeToString(d.Sum(nil)))
			if got != kat.Digest {
				t.Errorf("function=%s, length=%d\nmessage:\n %s\ngot:\n  %s\nwanted:\n %s",
					algo, kat.Length, kat.Message, got, kat.Digest)
				t.Logf("wanted %+v", kat)
				t.FailNow()
			}
			continue
		}
	}
}

// TestUnalignedWrite tests that writing data in an arbitrary pattern with
// small input buffers.
func TestUnalignedWrite(t *testing.T) {
	buf := sequentialBytes(0x10000)
	for alg, df := range testDigests {
		d := df()
		d.Reset()
		d.Write(buf)
		want := d.Sum(nil)
		d.Reset()
		for i := 0; i < len(buf); {
			// Cycle through offsets which make a 137 byte sequence.
			// Because 137 is prime this sequence should exercise all corner cases.
			offsets := [17]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1}
			for _, j := range offsets {
				if v := len(buf) - i; v < j {
					j = v
				}
				d.Write(buf[i : i+j])
				i += j
			}
		}
		got := d.Sum(nil)
		if !bytes.Equal(got, want) {
			t.Errorf("Unaligned writes, alg=%s\ngot %q, want %q", alg, got, want)
		}
	}
}

// TestAppend checks that appending works when reallocation is necessary.
func TestAppend(t *testing.T) {
	d := New224()

	for capacity := 2; capacity <= 66; capacity += 64 {
		// The first time around the loop, Sum will have to reallocate.
		// The second time, it will not.
		buf := make([]byte, 2, capacity)
		d.Reset()
		d.Write([]byte{0xcc})
		buf = d.Sum(buf)
		expected := "0000DF70ADC49B2E76EEE3A6931B93FA41841C3AF2CDF5B32A18B5478C39"
		if got := strings.ToUpper(hex.EncodeToString(buf)); got != expected {
			t.Errorf("got %s, want %s", got, expected)
		}
	}
}

// TestAppendNoRealloc tests that appending works when no reallocation is necessary.
func TestAppendNoRealloc(t *testing.T) {
	buf := make([]byte, 1, 200)
	d := New224()
	d.Write([]byte{0xcc})
	buf = d.Sum(buf)
	expected := "00DF70ADC49B2E76EEE3A6931B93FA41841C3AF2CDF5B32A18B5478C39"
	if got := strings.ToUpper(hex.EncodeToString(buf)); got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

// sequentialBytes produces a buffer of size consecutive bytes 0x00, 0x01, ..., used for testing.
//
// The alignment of each slice is intentionally randomized to detect alignment
// issues in the implementation. See https://golang.org/issue/37644.
// Ideally, the compiler should fuzz the alignment itself.
// (See https://golang.org/issue/35128.)
func sequentialBytes(size int) []byte {
	alignmentOffset := rand.Intn(8)
	result := make([]byte, size+alignmentOffset)[alignmentOffset:]
	for i := range result {
		result[i] = byte(i)
	}
	return result
}

// benchmarkHash tests the speed to hash num buffers of buflen each.
func benchmarkHash(b *testing.B, h hash.Hash, size, num int) {
	b.StopTimer()
	h.Reset()
	data := sequentialBytes(size)
	b.SetBytes(int64(size * num))
	b.StartTimer()

	var state []byte
	for i := 0; i < b.N; i++ {
		for j := 0; j < num; j++ {
			h.Write(data)
		}
		state = h.Sum(state[:0])
	}
	b.StopTimer()
	h.Reset()
}

func BenchmarkSha3_512_MTU(b *testing.B) { benchmarkHash(b, New512(), 1350, 1) }
func BenchmarkSha3_384_MTU(b *testing.B) { benchmarkHash(b, New384(), 1350, 1) }
func BenchmarkSha3_256_MTU(b *testing.B) { benchmarkHash(b, New256(), 1350, 1) }
func BenchmarkSha3_224_MTU(b *testing.B) { benchmarkHash(b, New224(), 1350, 1) }

func BenchmarkSha3_512_1MiB(b *testing.B) { benchmarkHash(b, New512(), 1024, 1024) }
