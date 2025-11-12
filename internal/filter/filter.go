package filter

import (
	"math/big"
)

type Filter struct {
	bitMap        []bool
	hashFunctions []HashFunction
	capacity      int
}

func (f *Filter) GetReadebleBitmap() string {
	res := ""

	for _, val := range f.bitMap {
		if val {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}

func NewFilter(size int, h []HashFunction) Filter {
	return Filter{
		make([]bool, size),
		h,
		size,
	}
}

func (f *Filter) AddValue(value string) error {
	for _, hFunc := range f.hashFunctions {
		hRes, err := hFunc(value)
		if err != nil {
			return err
		}

		bitMapIdx := modBytesByCapacity(hRes, f.capacity)
		f.bitMap[bitMapIdx] = true
	}
	return nil
}

func (f *Filter) IsValueExists(value string) (bool, error) {
	for _, hFunc := range f.hashFunctions {
		hRes, err := hFunc(value)
		if err != nil {
			return false, err
		}

		bitMapIdx := modBytesByCapacity(hRes, f.capacity)
		if !f.bitMap[bitMapIdx] {
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
