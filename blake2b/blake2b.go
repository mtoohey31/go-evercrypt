// Package blake2b implements the BLAKE2b hash algorithm defined by RFC 7693.
package blake2b

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

func init() { crypto.RegisterHash(crypto.BLAKE2b_512, New512) }

// The blocksize of BLAKE2b in bytes.
const BlockSize = 128

// The hash size of BLAKE2b-512 in bytes.
const Size = C.BLAKE2B_HASH_LEN

// New512 returns a new hash.Hash computing the BLAKE2b-512 checksum.
func New512() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_Blake2B, BlockSize)
}

// Sum512 returns the BLAKE2b-512 checksum of the data.
func Sum512(data []byte) [Size]byte {
	var res [Size]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_Blake2B, data, res[:])
	return res
}
