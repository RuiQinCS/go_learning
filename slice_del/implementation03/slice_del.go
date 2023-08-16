package implementation03

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标范围错误")

// DelElem by tree 2023-07-28
func DelElem[T any](data []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(data) {
		return nil, fmt.Errorf("err: %w, 长度 %d, 下标 %d", ErrIndexOutOfRange, len(data), idx)
	}

	length := len(data)

	/*
		直接操作原切片不利于底层数组中无用数据的回收
	*/
	copy(data[idx:length-1], data[idx+1:])

	return data[0 : length-1], nil
}
