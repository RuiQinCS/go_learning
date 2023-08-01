package implementation01

import "errors"

// DelElem by tree 2023-07-28
func DelElem(data []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(data) {
		return nil, errors.New("wrong index")
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
