// Package sha256 implements the SHA224 and SHA256 hash algorithms as defined
// in FIPS 180-4.
package sha256

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
	crypto.RegisterHash(crypto.SHA224, New224)
	crypto.RegisterHash(crypto.SHA256, New)
}

// The size of a SHA256 checksum in bytes.
const Size = C.SHA2_256_HASH_LEN

// The size of a SHA224 checksum in bytes.
const Size224 = C.SHA2_224_HASH_LEN

// The blocksize of SHA256 and SHA224 in bytes.
const BlockSize = 64

// New returns a new hash.Hash computing the SHA256 checksum.
func New() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA2_256)
}

// New224 returns a new hash.Hash computing the SHA224 checksum.
func New224() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA2_224)
}

// Sum256 returns the SHA256 checksum of the data.
func Sum256(data []byte) [Size]byte {
	var res [Size]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA2_256, data, res[:])
	return res
}

// Sum224 returns the SHA224 checksum of the data.
func Sum224(data []byte) [Size224]byte {
	var res [Size224]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA2_224, data, res[:])
	return res
}
