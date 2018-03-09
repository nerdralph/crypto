// Copyright 2018 Ralph Doncaster. All rights reserved.

package sha3

// This file provides functions for creating instances of the Keccak
// functions.  These are the same as the sha3 functions except that
// a domain separation byte (dsbyte) of 0x01 is used instead of 0x06

import (
	"hash"
	"unsafe"
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

// for repetitive hashing using the digest output of one hash as the input
// for the next hash
type ReHash interface {
	Hash()
	Data() []byte
}

func (s* state) Data() []byte {
	// statte array is first struct member and has the same address
	arr := (*[64]byte)(unsafe.Pointer(s))
	return arr[:s.outputLen]
}

// keccakF1600
func (s* state) Hash() {
	//st := (*[25]uint64)(unsafe.Pointer(h.State))
	st := (*[25]uint64)(unsafe.Pointer(s))
	pos := s.outputLen/8
	st[pos] = uint64(s.dsbyte)
	// zero the rest of the state
	for i := pos+1; i < 25; i++ {
		st[i] = 0
	}
	// add pad bit
	st[25 - 2*pos -1] ^= 1<<63
	keccakF1600(st)
}

// Creates a new Keccak-256 ReHash.
func ReHashK256() ReHash { return &state{outputLen: 32, dsbyte: 0x01} }

// Creates a new Keccak-512 ReHash.
func ReHashK512() ReHash { return &state{outputLen: 64, dsbyte: 0x01} }

