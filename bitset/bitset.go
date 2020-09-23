package bitset

import (
	"log"
)

const (
	WordSize = 64
	One = 1
)

type BitSet struct {
	words      []uint64
	wordsInUse int
}

func NewBitSet() *BitSet {
	return &BitSet{
		words: make([]uint64,1),
		wordsInUse:1,
	}
}

func (b *BitSet) Set(bitIndex int)  {
	if bitIndex<0 {
		log.Panicf("bitIndex<0: %d",bitIndex)
	}

	wordIndex := b.wordIndex(bitIndex)
	bitIndex %= WordSize

	needMoreWords := wordIndex +1 - len(b.words)
	if needMoreWords >0 {
		b.words = append(b.words,make([]uint64, needMoreWords)...)
		b.wordsInUse += needMoreWords
	}
	b.words[wordIndex] |= One<<bitIndex
}

func (b *BitSet) Get(bitIndex int) bool {
	if bitIndex<0 {
		log.Panicf("bitIndex<0: %d",bitIndex)
	}

	wordIndex := b.wordIndex(bitIndex)
	bitIndex%= WordSize

	if wordIndex >=b.wordsInUse {
		return false
	}
	return b.words[wordIndex] & (One<<bitIndex) != 0
}

func (b *BitSet) Clean(bitIndex int)  {
	if bitIndex<0 {
		log.Panicf("bitIndex<0: %d",bitIndex)
	}

	wordIndex := b.wordIndex(bitIndex)
	bitIndex %= WordSize
	if wordIndex >= b.wordsInUse {
		return
	}
	b.words[wordIndex] ^= One<<bitIndex
}

func (b *BitSet) wordIndex(bitIndex int) int {
	return bitIndex>>6
}


