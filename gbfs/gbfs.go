package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/sha3"
	"hash"
	"math/rand"
	"time"
)

func main() {
	GBFs:= genBloom(4, []string{"12", "121"}, 1024)
	fmt.Println(GBFs)
	result := queryGarbled("12", GBFs, 1024)
	fmt.Println(result)
}

func genBloom(lambda int, inputArray []string, m uint64) []int {
	var GBFs = make([]int, m)
	for element := range inputArray {
		var emptySlot int64 = -1
		var finalShare = element
		var hashes = []hash.Hash{sha1.New(), sha256.New(), sha512.New(), sha3.New224(), sha3.New256()}
		for i := 0; i < len(hashes); i++ {
			hashes[i].Write([]byte(inputArray[element]))
			hashData := hashes[i].Sum(nil)
			j := BytesToInt64(hashData) % m
			if GBFs[j] == 0 {
				if emptySlot == -1 {
					emptySlot = int64(j)
				} else {
					newShare := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(lambda))
					GBFs[j] = int(newShare)
					finalShare = finalShare ^ GBFs[j]
				}
			} else {
				finalShare = finalShare ^ GBFs[j]
			}
		}
		for i := 0; uint64(i) < m; i++ {
			if GBFs[i] == 0 {
				GBFs[i] = int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(lambda)))
			}
		}
	}
	return GBFs
}

func queryGarbled(x string, GBFs []int, m uint64) bool {
	var hashes = []hash.Hash{sha1.New(), sha256.New(), sha512.New(), sha3.New224(), sha3.New256()}
	for i := 0; i < len(hashes); i++ {
		hashes[i].Write([]byte(x))
		hashData := hashes[i].Sum(nil)
		j := BytesToInt64(hashData) % m
		if GBFs[j] == 0 {
			return false
		}
	}
	return true
}

func BytesToInt64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}
