package implementation02

import (
	"errors"
	"fmt"
)

/*
	 DelElem input :
		data : 切片数据
		idx : 删除该位置（含义为数组下标）的数据
	 output :
		返回删除后的切片
		错误
*/
var ErrIndexOutOfRange = errors.New("下标范围错误")

// DelElem by tree 2023-07-28
func DelElem(data []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(data) {
		return nil, fmt.Errorf("err: %w, 长度 %d, 下标 %d", ErrIndexOutOfRange, len(data), idx)
	}

	//idx >= len(data) could handle
	//if data == nil || len(data) == 0 {
	//	return nil, errors.New("wrong data")
	//}

	/*
		新建切片利于原切片对应底层数组中无用数据的回收
	*/
	res := make([]int, len(data)-1)

	copy(res[0:idx], data[0:idx])
	copy(res[idx:], data[idx+1:])

	return res, nil
}
