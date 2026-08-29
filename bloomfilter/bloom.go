package main

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	bits   [1000]byte
	length int //8000
	hashes int
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{
		bits:   [1000]byte{},
		length: 8000,
		hashes: 5,
	}
}

func hash(s string) (uint32, uint32) {
	h1 := murmur3.Sum32([]byte(s))
	h2 := murmur3.Sum32([]byte(s + "|salt"))
	return h1, h2
}

func (bf *BloomFilter) add(s string) {
	h1, h2 := hash(s)

	for i := 0; i < bf.hashes; i++ {
		idx := (h1 + uint32(i)*h2) % uint32(bf.length)
		bf.bits[idx/8] |= 1 << (idx % 8)
	}
}

func (bf *BloomFilter) found(s string) bool {
	h1, h2 := hash(s)
	for i := 0; i < bf.hashes; i++ {
		idx := (uint32(h1) + uint32(i)*uint32(h2)) % uint32(bf.length)
		if bf.bits[idx/8]&(1<<(idx%8)) == 0 {
			return false
		}
	}
	return true
}

func main() {
	bf := NewBloomFilter()

	items := []string{
		"apple",
		"banana",
		"grape",
		"mango",
	}

	for _, item := range items {
		bf.add(item)
	}

	tests := []string{
		"apple",
		"banana",
		"orange",
		"mango",
		"pear",
	}

	for _, item := range tests {
		fmt.Printf("%s => %v\n", item, bf.found(item))
	}
}
