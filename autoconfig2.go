package evercrypt

// #cgo LDFLAGS: -levercrypt
// #define HACL_CAN_COMPILE_VEC128
// #define HACL_CAN_COMPILE_VEC256
// #include <EverCrypt_AutoConfig2.h>
import "C"

func init() {
	C.EverCrypt_AutoConfig2_init()
}
