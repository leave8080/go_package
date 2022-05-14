package util

// IF 三元运算
func IF(expr bool, turePart, falsePart interface{}) interface{} {
	if expr {
		return turePart
	}
	return falsePart
}

// Any 同python any逻辑, 任何一个为输入true，结果即为true
func Any(boolIter ...bool) bool {
	for _, iter := range boolIter {
		if iter {
			return true
		}
	}
	return false
}

// All 同python all逻辑, 所有输入为ture结构为true
func All(boolIter ...bool) bool {
	for _, iter := range boolIter {
		if !iter {
			return false
		}
	}
	return true
}
