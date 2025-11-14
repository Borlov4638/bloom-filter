package filter

import (
	"math"
	"math/big"
)

type Filter struct {
	BitMap        []bool
	HashFunctions []HashFunction
	Capacity      int
}

func (f *Filter) GetReadebleBitmap() string {
	res := ""

	for _, val := range f.BitMap {
		if val {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}

func NewFilterByProb(p float64, n int) Filter {
	k, m := getOptimalParams(p, n)
	hashFunctions := getKHashFunctionos(k)
	return NewFilter(m, hashFunctions)
}

func NewFilter(size int, h []HashFunction) Filter {
	return Filter{
		make([]bool, size),
		h,
		size,
	}
}

func (f *Filter) AddValue(value string) error {
	for _, hFunc := range f.HashFunctions {
		hRes, err := hFunc(value)
		if err != nil {
			return err
		}

		bitMapIdx := modBytesByCapacity(hRes, f.Capacity)
		f.BitMap[bitMapIdx] = true
	}
	return nil
}

func (f *Filter) IsValueExists(value string) (bool, error) {
	for _, hFunc := range f.HashFunctions {
		hRes, err := hFunc(value)
		if err != nil {
			return false, err
		}

		bitMapIdx := modBytesByCapacity(hRes, f.Capacity)
		if !f.BitMap[bitMapIdx] {
			return false, nil
		}
	}
	return true, nil
}

func modBytesByCapacity(b []byte, capacity int) int {
	intager := new(big.Int).SetBytes(b)

	bigCap := big.NewInt(int64(capacity))

	z := new(big.Int).Mod(intager, bigCap)
	return int(z.Int64())
}

func getOptimalParams(p float64, n int) (k, m int) {
	m = -int(math.Floor((float64(n) * math.Log(p)) / (math.Pow(math.Log(2.0), 2))))
	log2 := float64(math.Log(2.0))
	k = int(math.Floor((float64(m) / float64(n)) * log2))
	return k, m
}
