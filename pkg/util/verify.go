package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// 计算规则参考“中国国家标准化管理委员会”官方文档：http://www.gb688.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E
// 身份证号码校验
func VerifyIDCard(idCard string) bool {
	if len([]rune(idCard)) != 18 {
		return false
	}
	// a1与对应的校验码对照表，其中key表示a1，value表示校验码，value中的10表示校验码X
	var a1Map = map[int]int{
		0:  1,
		1:  0,
		2:  10,
		3:  9,
		4:  8,
		5:  7,
		6:  6,
		7:  5,
		8:  4,
		9:  3,
		10: 2,
	}

	var idStr = strings.ToUpper(idCard)
	var reg, err = regexp.Compile(`^[0-9]{17}[0-9X]$`)
	if err != nil {
		return false
	}
	if !reg.Match([]byte(idStr)) {
		return false
	}
	var sum int
	var signChar = ""
	for index, c := range idStr {
		var i = 18 - index
		if i != 1 {
			if v, err := strconv.Atoi(string(c)); err == nil {
				// 计算加权因子
				var weight = int(math.Pow(2, float64(i-1))) % 11
				sum += v * weight
			} else {
				return false
			}
		} else {
			signChar = string(c)
		}
	}
	var a1 = a1Map[sum%11]
	var a1Str = fmt.Sprintf("%d", a1)
	if a1 == 10 {
		a1Str = "X"
	}
	return a1Str == signChar
}

// 手机号码校验
func VerifyPhoneNumber(phone string) bool {
	if len([]rune(phone)) != 11 {
		return false
	}
	reg, err := regexp.Compile(`^1[0-9]{10}$`)
	if err != nil {
		return false
	}
	return reg.Match([]byte(phone))
}

// 邮箱格式验证
func ValidateEmail(value string) bool {
	flag, err := regexp.MatchString(`^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$`, value)
	if err != nil {
		return false
	}
	return flag
}

// GeneratePassword pwd 为明文密码，salt 为用户的openid
func GeneratePassword(pwd string, salt string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	p := hex.EncodeToString(h.Sum(nil))
	h.Reset()
	h.Write([]byte(p))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateOpenid 生成Openid
func GenerateOpenid(prefix, appid, uname string) string {
	h := md5.New()
	h.Write([]byte(appid))
	h.Write([]byte(uname))
	h.Write([]byte(time.Now().String()))

	return prefix + hex.EncodeToString(h.Sum(nil))[2:]
}

// GenerateToken 生成token
func GenerateToken() string {
	uu := uuid.New().String()
	return strings.ReplaceAll(uu, "-", "")
}
