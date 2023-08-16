package implementation04

import (
	"errors"
	"fmt"
	"runtime"
)

var ErrIndexOutOfRange = errors.New("下标范围错误")

// DelElem by tree 2023-08-02
func DelElem[T any](data []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(data) {
		return nil, fmt.Errorf("err: %w, 长度 %d, 下标 %d", ErrIndexOutOfRange, len(data), idx)
	}

	oldCap := cap(data)
	length := len(data)

	copy(data[idx:length-1], data[idx+1:])

	return shrinkSlice(data[0:length-1], oldCap), nil
}

func shrinkSlice[T any](data []T, oldCap int) []T {
	newLen := len(data)
	newCap := shrink(newLen, oldCap)

	if newCap == oldCap {
		return data
	}

	newData := make([]T, newLen, newCap) // len 为 newLen，而非0
	copy(newData[0:newLen], data[0:newLen])

	data = nil
	runtime.GC()

	return newData
}

func shrink(newLen, oldCap int) (newCap int) {
	if newLen < oldCap>>1 {
		return oldCap >> 1
	}

	threshold := 256

	if oldCap <= threshold {
		return oldCap
	}

	newCap = oldCap
	factor := 0.9

	for int(float64(newCap)*factor) > newLen {
		newCap = int(float64(newCap) * factor)
	}

	return newCap
}
