// Package sha512 implements the SHA-384 and SHA-512 hash algorithms as defined
// in FIPS 180-4.
package sha512

// #cgo LDFLAGS: -levercrypt
// #define HACL_CAN_COMPILE_VEC128
// #define HACL_CAN_COMPILE_VEC256
// #include <EverCrypt_Hash.h>
import "C"
import (
	"crypto"
	"hash"

	internal_hash "mtoohey.com/go-evercrypt/internal/hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA384, New384)
	crypto.RegisterHash(crypto.SHA512, New)
}

// Size is the size, in bytes, of a SHA-512 checksum.
const Size = C.SHA2_512_HASH_LEN

// Size384 is the size, in bytes, of a SHA-384 checksum.
const Size384 = C.SHA2_384_HASH_LEN

// BlockSize is the block size, in bytes, of the SHA-384 and SHA-512 hash
// functions.
const BlockSize = 64

// New returns a new hash.Hash computing the SHA-512 checksum.
func New() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA2_512)
}

// New384 returns a new hash.Hash computing the SHA-384 checksum.
func New384() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA2_384)
}

// Sum512 returns the SHA512 checksum of the data.
func Sum512(data []byte) [Size]byte {
	var res [Size]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA2_512, data, res[:])
	return res
}

// Sum384 returns the SHA384 checksum of the data.
func Sum384(data []byte) [Size384]byte {
	var res [Size384]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA2_384, data, res[:])
	return res
}
