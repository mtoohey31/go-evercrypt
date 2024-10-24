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
	"errors"
	"hash"
	"runtime"
	"unsafe"

	_ "mtoohey.com/go-evercrypt/internal/autoconfig2"
)

func init() { crypto.RegisterHash(crypto.MD5, New) }

// The size of an MD5 checksum in bytes.
const Size = C.MD5_HASH_LEN

// The blocksize of MD5 in bytes.
const BlockSize = 64

var MaximumLengthExceeded = errors.New("maximum length exceeded")

type digest struct {
	inner *C.struct_EverCrypt_Hash_Incremental_state_t_s
}

// New returns a new hash.Hash computing the MD5 checksum.
func New() hash.Hash {
	res := &digest{inner: C.EverCrypt_Hash_Incremental_malloc(C.Spec_Hash_Definitions_MD5)}
	runtime.SetFinalizer(res, func(d *digest) {
		C.EverCrypt_Hash_Incremental_free(d.inner)
	})
	return res
}

func (d *digest) Reset() { C.EverCrypt_Hash_Incremental_reset(d.inner) }

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Write(p []byte) (n int, err error) {
	res := C.EverCrypt_Hash_Incremental_update(d.inner, (*C.uchar)(unsafe.SliceData(p)), C.uint32_t(len(p)))
	switch res {
	case C.EverCrypt_Error_Success:
		return len(p), nil

	case C.EverCrypt_Error_MaximumLengthExceeded:
		return 0, MaximumLengthExceeded

	default:
		panic("EverCrypt_Hash_Incremental_update returned unexpected error code")
	}
}

func (d *digest) Sum(b []byte) []byte {
	var res [Size]byte
	C.EverCrypt_Hash_Incremental_digest(d.inner, (*C.uchar)(unsafe.SliceData(res[:])))
	return res[:]
}

// Sum returns the MD5 checksum of the data.
func Sum(data []byte) [Size]byte {
	var res [Size]byte
	C.EverCrypt_Hash_Incremental_hash(C.Spec_Hash_Definitions_MD5, (*C.uchar)(unsafe.SliceData(res[:])), (*C.uchar)(unsafe.SliceData(data)), C.uint32_t(len(data)))
	return res
}
