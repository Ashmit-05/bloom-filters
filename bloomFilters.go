package main

import "errors"

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
