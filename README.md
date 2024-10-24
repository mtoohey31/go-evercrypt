# go-evercrypt

A Go wrapper for [EverCrypt](https://github.com/hacl-star/hacl-star#evercrypt). APIs are intended to be as close as possible to the existing equivalent Go packages to make drop-in replacement possible. To build this package, you must have the EverCrypt C library installed.

This is currently a WIP. Here is a list of the corresponding Go packages and EverCrypt functionality that I'm aware of, in a checklist tracking what's been implemented in this module:
- [ ] [`crypto/aes`](https://pkg.go.dev/crypto/aes)/[`EverCrypt_AEAD`](https://hacl-star.github.io/EverCryptAEAD.html)
- [ ] [`crypto/cipher`](https://pkg.go.dev/crypto/cipher)/[`EverCrypt_CTR`](https://hacl-star.github.io/EverCryptCTR.html)
- [ ] [`crypto/ecdh`](https://pkg.go.dev/crypto/ecdh)/[`EverCrypt_Curve25519`](https://hacl-star.github.io/EverCryptNonAgile.html#curve25519)
- [ ] [`crypto/ed25519`](https://pkg.go.dev/crypto/ed25519)/[`EverCrypt_Ed25519`](https://hacl-star.github.io/EverCryptNonAgile.html#ed25519)
- [ ] [`crypto/hmac`](https://pkg.go.dev/crypto/hmac)/[`EverCrypt_HMAC`](https://hacl-star.github.io/EverCryptHMAC.html)
- [x] [`crypto/md5`](https://pkg.go.dev/crypto/md5)/[`EverCrypt_Hash`](https://hacl-star.github.io/EverCryptHash.html)
- [x] [`crypto/sha1`](https://pkg.go.dev/crypto/sha1)/[`EverCrypt_Hash`](https://hacl-star.github.io/EverCryptHash.html)
- [x] [`crypto/sha256`](https://pkg.go.dev/crypto/sha256)/[`EverCrypt_Hash`](https://hacl-star.github.io/EverCryptHash.html)
- [x] [`crypto/sha512`](https://pkg.go.dev/crypto/sha512)/[`EverCrypt_Hash`](https://hacl-star.github.io/EverCryptHash.html)
- [x] [`golang.org/x/crypto/blake2b`](https://pkg.go.dev/golang.org/x/crypto/blake2b)/[`EverCrypt_Hash`](https://hacl-star.github.io/EverCryptHash.html)
- [x] [`golang.org/x/crypto/blake2s`](https://pkg.go.dev/golang.org/x/crypto/blake2s)/[`EverCrypt_Hash`](https://hacl-star.github.io/EverCryptHash.html)
- [ ] [`golang.org/x/crypto/chacha20poly1305`](https://pkg.go.dev/golang.org/x/crypto/chacha20poly1305)/[`EverCrypt_AEAD`](https://hacl-star.github.io/EverCryptAEAD.html)
- [ ] [`golang.org/x/crypto/hkdf`](https://pkg.go.dev/golang.org/x/crypto/hkdf)/[`EverCrypt_HKDF`](https://hacl-star.github.io/EverCryptHKDF.html)
- [ ] [`golang.org/x/crypto/poly1305`](https://pkg.go.dev/golang.org/x/crypto/poly1305)/[`EverCrypt_Poly1305`](https://hacl-star.github.io/EverCryptNonAgile.html#poly1305)
- [x] [`golang.org/x/crypto/sha3`](https://pkg.go.dev/golang.org/x/crypto/sha3)/[`EverCrypt_Hash`](https://hacl-star.github.io/EverCryptHash.html)
