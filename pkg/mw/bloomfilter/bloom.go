package bloomfilter

import (
	"github.com/bits-and-blooms/bloom/v3"
)

type BloomFilter struct {
	Filter *bloom.BloomFilter
}

func NewBloomFilter(n uint, fp float64) *BloomFilter {
	return &BloomFilter{
		Filter: bloom.NewWithEstimates(n, fp), //预期元素数量误判率
	}
}

func (f *BloomFilter) AddToBloomFilter(data string) {
	f.Filter.Add([]byte(data))
}

func (f *BloomFilter) TestBloom(data string) bool {
	return f.Filter.Test([]byte(data))
}
