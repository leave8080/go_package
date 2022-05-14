package util

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	NULL       = ""
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
)

// 获取随机字符串
//    length：字符串长度
func RandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

// 获取随机数组
//    length：字符串长度
func RandomInt(length int) string {
	str := "0123456789"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

// 解析时间
func ParseDateTime(timeStr string) (datetime time.Time) {
	datetime, _ = time.ParseInLocation(TimeLayout, timeStr, time.Local)
	return
}

// 字符串转Float64
func String2Float64(floatStr string) (floatNum float64) {
	floatNum, _ = strconv.ParseFloat(floatStr, 64)
	return
}

// 字符串转Float32
func String2Float32(floatStr string) (floatNum float32) {
	float, _ := strconv.ParseFloat(floatStr, 32)
	floatNum = float32(float)
	return
}

// Float64转字符串
//    floatNum：float64数字
//    prec：精度位数（不传则默认float数字精度）
func Float64ToString(floatNum float64, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(floatNum, 'f', prec[0], 64)
		return
	}
	floatStr = strconv.FormatFloat(floatNum, 'f', -1, 64)
	return
}

// Float32转字符串
//    floatNum：float32数字
//    prec：精度位数（不传则默认float数字精度）
func Float32ToString(floatNum float32, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(float64(floatNum), 'f', prec[0], 32)
		return
	}
	floatStr = strconv.FormatFloat(float64(floatNum), 'f', -1, 32)
	return
}

// 字符串转Int
func String2Int(intStr string) (intNum int) {
	intNum, _ = strconv.Atoi(intStr)
	return
}

// 字符串转Int32
func String2Int32(intStr string) (int32Num int32) {
	intNum, _ := strconv.Atoi(intStr)
	int32Num = int32(intNum)
	return
}

// 字符串转Int64
func String2Int64(intStr string) (int64Num int64) {
	intNum, _ := strconv.Atoi(intStr)
	int64Num = int64(intNum)
	return
}

// Int转字符串
func Int2String(intNum int) (intStr string) {
	intStr = strconv.Itoa(intNum)
	return
}

// Int32转字符串
func Int322String(intNum int32) (int32Str string) {
	// 10, 代表10进制
	int32Str = strconv.FormatInt(int64(intNum), 10)
	return
}

// Int64转字符串
func Int642String(intNum int64) (int64Str string) {
	// 10, 代表10进制
	int64Str = strconv.FormatInt(intNum, 10)
	return
}
