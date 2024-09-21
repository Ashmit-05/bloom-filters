package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	fnvOffset uint32 = 2166136261
	seed      uint32 = 2
	fnvPrime         = 37
	c1               = 0xcc9e2d51
	c2               = 0x1b873593
)

type BloomFilter struct {
	arr  []int
	size uint32
}

func CreateBloomFilter(size uint32) (*BloomFilter, error) {
	if size <= 0 {
		return nil, errors.New("Size of the bloom filter must be more than 0")
	}
	return &BloomFilter{arr: make([]int, size), size: size}, nil
}

func (b *BloomFilter) Add(key string) {
	h1 := FnvHash(key)
	h2 := MurmurHash3(key)
	h3 := h1 + 2*h2
	pos1 := h1 % b.size
	pos2 := h2 % b.size
	pos3 := h3 % b.size
	b.arr[pos1] = 1
	b.arr[pos2] = 1
	b.arr[pos3] = 1
}

func (b *BloomFilter) Contains(key string) bool {
	h1 := FnvHash(key)
	h2 := MurmurHash3(key)
	h3 := h1 + 2*h2
	pos1 := h1 % b.size
	pos2 := h2 % b.size
	pos3 := h3 % b.size
	if b.arr[pos1] == 1 && b.arr[pos2] == 1 && b.arr[pos3] == 1 {
		return true
	}
	return false
}

func (b *BloomFilter) Clear() {
	for i := range b.arr {
		b.arr[i] = 0
	}
}

func FnvHash(key string) uint32 {
	hashValue := fnvOffset
	for _, char := range key {
		r := hashValue * fnvPrime
		hashValue = r ^ uint32(char)
	}
	return hashValue
}

// not providing seed as an argument
func MurmurHash3(key string) uint32 {
	bytes := []byte(key)
	length := len(bytes)
	hash := seed

	// Process the input in 4-byte chunks
	for i := 0; i < length; i += 4 {
		var k uint32

		if i+4 <= length {
			k = binary.BigEndian.Uint32(bytes[i : i+4])
		} else {
			for j := 0; j < length-i; j++ {
				k |= uint32(bytes[i+j]) << (8 * (3 - j))
			}
		}

		k *= c1
		k = (k << 15) | (k >> (32 - 15))
		k *= c2

		hash ^= k
		hash = (hash << 13) | (hash >> (32 - 13))
		hash = hash*5 + 0xe6546b64
	}

	// Finalization
	hash ^= uint32(length)
	hash ^= hash >> 16
	hash *= 0x85ebca6b
	hash ^= hash >> 13
	hash *= 0xc2b2ae35
	hash ^= hash >> 16

	return hash
}

func main() {
	b, err := CreateBloomFilter(16)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		keys := []string{"bat", "rat", "cats"}
		for _, key := range keys {
			b.Add(key)
		}
		for _, key := range keys {
			fmt.Println(key, b.Contains(key))
		}
		fmt.Println("cat", b.Contains("cat"))
		b.Clear()
		fmt.Println("After clearing..")
		fmt.Println("bat", b.Contains("bat"))
	}
}
