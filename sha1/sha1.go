// Package sha1 implements the SHA-1 hash algorithm as defined in RFC 3174.
//
// SHA-1 is cryptographically broken and should not be used for secure
// applications.
package sha1

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

func init() { crypto.RegisterHash(crypto.SHA1, New) }

// The size of a SHA-1 checksum in bytes.
const Size = C.SHA1_HASH_LEN

// The blocksize of SHA-1 in bytes.
const BlockSize = 64

// New returns a new hash.Hash computing the SHA1 checksum.
func New() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_SHA1, BlockSize)
}

// Sum returns the SHA-1 checksum of the data.
func Sum(data []byte) [Size]byte {
	var res [Size]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_SHA1, data, res[:])
	return res
}
