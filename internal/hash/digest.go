package hash

// #cgo LDFLAGS: -levercrypt
// #define HACL_CAN_COMPILE_VEC128
// #define HACL_CAN_COMPILE_VEC256
// #include <EverCrypt_Hash.h>
import "C"
import (
	"hash"
	"runtime"
	"unsafe"
)

type Digest struct {
	inner *C.struct_EverCrypt_Hash_Incremental_state_t_s
}

func New(a C.Spec_Hash_Definitions_hash_alg) hash.Hash {
	res := &Digest{inner: C.EverCrypt_Hash_Incremental_malloc(a)}
	runtime.SetFinalizer(res, func(d *Digest) {
		C.EverCrypt_Hash_Incremental_free(d.inner)
	})
	return res
}

func (d *Digest) Reset() { C.EverCrypt_Hash_Incremental_reset(d.inner) }

func (d *Digest) Size() int {
	a := C.EverCrypt_Hash_Incremental_alg_of_state(d.inner)
	return int(C.EverCrypt_Hash_Incremental_hash_len(a))
}

const BlockSize = 64

func (d *Digest) BlockSize() int { return BlockSize }

func (d *Digest) Write(p []byte) (n int, err error) {
	res := C.EverCrypt_Hash_Incremental_update(d.inner, (*C.uchar)(unsafe.SliceData(p)), C.uint32_t(len(p)))
	switch res {
	case C.EverCrypt_Error_Success:
		return len(p), nil

	case C.EverCrypt_Error_MaximumLengthExceeded:
		panic("EverCrypt_Hash_Incremental_update returned EverCrypt_Error_MaximumLengthExceeded")

	default:
		panic("EverCrypt_Hash_Incremental_update returned unexpected error code")
	}
}

func (d *Digest) Sum(b []byte) []byte {
	res := make([]byte, d.Size())
	C.EverCrypt_Hash_Incremental_digest(d.inner, (*C.uchar)(unsafe.SliceData(res[:])))
	return append(b, res...)
}

func Sum(a C.Spec_Hash_Definitions_hash_alg, data []byte, dst []byte) {
	C.EverCrypt_Hash_Incremental_hash(a, (*C.uchar)(unsafe.SliceData(dst)), (*C.uchar)(unsafe.SliceData(data)), C.uint32_t(len(data)))
}
