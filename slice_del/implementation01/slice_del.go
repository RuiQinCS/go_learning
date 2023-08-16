package implementation01

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标范围错误")

// DelElem by tree 2023-07-28
func DelElem(data []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(data) {
		return nil, fmt.Errorf("err: %w, 长度 %d, 下标 %d", ErrIndexOutOfRange, len(data), idx)
	}

	res := make([]int, len(data)-1)
	for i := 0; i < idx; i++ {
		res[i] = data[i]
	}

	for i := idx + 1; i < len(data); i++ {
		res[i-1] = data[i]
	}

	return res, nil
}
