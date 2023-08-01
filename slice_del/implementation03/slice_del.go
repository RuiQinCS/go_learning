package implementation03

import (
	"errors"
)

// DelElem by tree 2023-07-28
func DelElem[T any](data []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(data) {
		return nil, errors.New("wrong index")
	}

	length := len(data)

	/*
		直接操作原切片不利于底层数组中无用数据的回收
	*/
	copy(data[idx:length-1], data[idx+1:])

	return data[0 : length-1], nil
}
