// Package blake2s implements the BLAKE2s hash algorithm defined by RFC 7693.
package blake2s

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

func init() { crypto.RegisterHash(crypto.BLAKE2s_256, New256) }

// The blocksize of BLAKE2s in bytes.
const BlockSize = 64

// The hash size of BLAKE2s-256 in bytes.
const Size = C.BLAKE2S_HASH_LEN

// New256 returns a new hash.Hash computing the BLAKE2s-256 checksum.
func New256() hash.Hash {
	return internal_hash.New(C.Spec_Hash_Definitions_Blake2S, BlockSize)
}

// Sum256 returns the BLAKE2b-512 checksum of the data.
func Sum256(data []byte) [Size]byte {
	var res [Size]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_Blake2S, data, res[:])
	return res
}
