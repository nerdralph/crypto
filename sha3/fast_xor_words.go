// code copied from crypto/cipher/xor.go
// for aligned use only

package sha3

import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))

// XORs multiples of 4 or 8 bytes (depending on architecture.)
// The arguments are assumed to be of equal length.
func FastXORWords(dst, a, b []byte) {
    dw := *(*[]uintptr)(unsafe.Pointer(&dst))
    aw := *(*[]uintptr)(unsafe.Pointer(&a))
    bw := *(*[]uintptr)(unsafe.Pointer(&b))
    n := len(b) / wordSize
    for i := 0; i < n; i++ {
        dw[i] = aw[i] ^ bw[i]
    }
}
