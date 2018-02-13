// Copyright 2018 Ralph Doncaster. All rights reserved.

package sha3

// This file provides functions for creating instances of the Keccak
// functions.  These are the same as the sha3 functions except that
// a domain separation byte (dsbyte) of 0x01 is used instead of 0x06

import (
	"hash"
)

// Creates a new Keccak-256 hash.
func K256() hash.Hash { return &state{rate: 136, outputLen: 32, dsbyte: 0x01} }

// Creates a new Keccak-512 hash.
func K512() hash.Hash { return &state{rate: 72, outputLen: 64, dsbyte: 0x01} }

// Calculates the Keccak-256 digest of the data.
func SumK256(data []byte) (digest [32]byte) {
	h := K256()
	h.Write(data)
	h.Sum(digest[:0])
	return
}

// Calculates the Keccak-512 digest of the data.
func SumK512(data []byte) (digest [64]byte) {
	h := K512()
	h.Write(data)
	h.Sum(digest[:0])
	return
}
