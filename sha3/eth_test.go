package sha3

import (
    "fmt"
    "math"
    "bytes"
    "encoding/hex"
)

const (
    cacheBYTESINIT = 16*1024*1024
    cacheBYTESGROWTH = 128*1024
    hashBYTES = 64
)

func isPrime(n int) bool {
    if (n==2)||(n==3) {return true;}
    if n%2 == 0 { return false }
    if n%3 == 0 { return false }
    sqrt := int(math.Sqrt(float64(n)))
    for i := int(5); i <= sqrt; i+=6 {
        if n%i == 0 { return false }
        if n%(i+2) == 0 { return false; }
    }
    return true
}

func cacheSize(epoch int) int {
    sz := cacheBYTESINIT + cacheBYTESGROWTH * epoch
    sz -= hashBYTES
    for ; !isPrime(int(sz / hashBYTES)); sz -= 2 * hashBYTES {}
    return sz
}

func makeCache(epoch int, seed []byte) []byte {
    sz := cacheSize(epoch)
    cache := make([]byte, sz)
	digest := SumK512(seed)
    copy(cache, digest[:])

    for pos := hashBYTES; pos < sz; pos += hashBYTES {
        digest := SumK512(cache[pos-hashBYTES:pos])
        copy(cache[pos:], digest[:])
    }
    return cache
}

func makeCacheFast(epoch int, seed []byte) []byte {
    sz := cacheSize(epoch)
    cache := make([]byte, sz)
	digestStart := SumK512(seed)
    copy(cache, digestStart[:])
	kf512 := ReHashK512()
	digest := kf512.Data()
	copy(digest, digestStart[:])

    for pos := hashBYTES; pos < sz; pos += hashBYTES {
		kf512.Hash()
        copy(cache[pos:], digest)
    }
    return cache
}

func ExampleEth(){
    var digest [32]byte
	epochSeed, _ := hex.DecodeString("f4c702c8373b1a64")
    epoch := 0
    for !bytes.Equal(digest[:8], epochSeed) {
        digest = SumK256(digest[:])
        epoch++
    }
    fmt.Printf("Epoch %d seed %x\n",epoch, digest)
    cache := makeCache(epoch, digest[:])
    fmt.Println("Cache size: ", len(cache))
    fmt.Printf("%x\n",cache[len(cache)-32:])
	// Output:
	// Epoch 169 seed f4c702c8373b1a6467eb30d4640b62436e55a6c8c9e2f486197f6e0b2dd72f5f
	// Cache size:  38925632
	// 9d1f8e1371dacc5e7169526ce97c2da815e414d6afc1a82cef53755e04bc2627
}

func ExampleEthFast(){
    epoch := 0
	epochSeed, _ := hex.DecodeString("f4c702c8373b1a64")
	kf256 := ReHashK256()
	digest := kf256.Data()
    for !bytes.Equal(digest[:8], epochSeed) {
		kf256.Hash()
        epoch++
    }
    fmt.Printf("Epoch %d seed %x\n",epoch, digest)
    cache := makeCacheFast(epoch, digest[:])
    fmt.Println("Cache size: ", len(cache))
    fmt.Printf("%x\n",cache[len(cache)-32:])
	// Output:
	// Epoch 169 seed f4c702c8373b1a6467eb30d4640b62436e55a6c8c9e2f486197f6e0b2dd72f5f
	// Cache size:  38925632
	// 9d1f8e1371dacc5e7169526ce97c2da815e414d6afc1a82cef53755e04bc2627
}

