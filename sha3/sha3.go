// Package sha3 implements the SHA-3 fixed-output-length hash functions and
// the SHAKE variable-output-length hash functions defined by FIPS-202.
package sha3

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
	crypto.RegisterHash(crypto.SHA3_224, New224)
	crypto.RegisterHash(crypto.SHA3_256, New256)
	crypto.RegisterHash(crypto.SHA3_384, New384)
	crypto.RegisterHash(crypto.SHA3_512, New512)
}

// New224 creates a new SHA3-224 hash.
// Its generic security strength is 224 bits against preimage attacks,
// and 112 bits against collision attacks.
func New224() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA3_224, 28)
}

// New256 creates a new SHA3-256 hash.
// Its generic security strength is 256 bits against preimage attacks,
// and 128 bits against collision attacks.
func New256() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA3_256, 32)
}

// New384 creates a new SHA3-384 hash.
// Its generic security strength is 384 bits against preimage attacks,
// and 192 bits against collision attacks.
func New384() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA3_384, 48)
}

// New512 creates a new SHA3-512 hash.
// Its generic security strength is 512 bits against preimage attacks,
// and 256 bits against collision attacks.
func New512() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA3_512, 64)
}

// Sum224 returns the SHA3-224 digest of the data.
func Sum224(data []byte) [28]byte {
	var res [28]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA3_224, data, res[:])
	return res
}

// Sum256 returns the SHA3-256 digest of the data.
func Sum256(data []byte) [32]byte {
	var res [32]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA3_256, data, res[:])
	return res
}

// Sum384 returns the SHA3-384 digest of the data.
func Sum384(data []byte) [48]byte {
	var res [48]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA3_384, data, res[:])
	return res
}

// Sum512 returns the SHA3-512 digest of the data.
func Sum512(data []byte) [64]byte {
	var res [64]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA3_512, data, res[:])
	return res
}
