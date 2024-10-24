// Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
//
// MD5 is cryptographically broken and should not be used for secure
// applications.
package md5

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

func init() { crypto.RegisterHash(crypto.MD5, New) }

// The size of an MD5 checksum in bytes.
const Size = C.MD5_HASH_LEN

// The blocksize of MD5 in bytes.
const BlockSize = internal_hash.BlockSize

// New returns a new hash.Hash computing the MD5 checksum.
func New() hash.Hash { return internal_hash.New(C.Spec_Hash_Definitions_MD5) }

// Sum returns the MD5 checksum of the data.
func Sum(data []byte) [Size]byte {
	var res [Size]byte
	internal_hash.Sum(C.Spec_Hash_Definitions_MD5, data, res[:])
	return res
}
