package web

import (
	"crypto/md5"
	"encoding/hex"
	"github/leave8080/go_package/pkg/eno"
	"github/leave8080/go_package/pkg/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	sign  = "sign"
	ts    = "ts"
	appid = "appid"
)

// VerifySign 验证Sign，签名规则，base64(md5(appid+path+ts))
func VerifySign(c *gin.Context) {
	ts := c.GetHeader(ts)
	tsTime := time.Unix(util.String2Int64(ts), 0)
	if time.Now().Sub(tsTime).Seconds() > 60 || tsTime.Sub(time.Now()).Seconds() > 60 {
		JSON(c, nil, eno.New(10401, "签名错误，请检查系统时间"))
		c.Abort()
		return
	}
	sign := c.GetHeader(sign)
	appid := c.GetHeader(appid)
	path := c.Request.RequestURI
	split := strings.Split(path, "?")
	if !CheckSign(sign, appid, split[0], ts) {
		JSON(c, nil, eno.InvalidSignErr)
		c.Abort()
		return
	}
}

// VerifySign 验证Sign，签名规则，base64(md5(appid+path+ts))
func VerifySignV2(sec float64) gin.HandlerFunc {
	return func(c *gin.Context) {
		ts := c.GetHeader(ts)
		tsTime := time.Unix(util.String2Int64(ts), 0)
		// 如果是小于 0 秒不进行时间校验
		if sec > 0 {
			if time.Now().Sub(tsTime).Seconds() > sec || tsTime.Sub(time.Now()).Seconds() > sec {
				JSON(c, nil, eno.New(10401, "签名错误，请检查系统时间"))
				c.Abort()
				return
			}
		}

		sign := c.GetHeader(sign)
		appid := c.GetHeader(appid)
		path := c.Request.RequestURI
		split := strings.Split(path, "?")
		if !CheckSign(sign, appid, split[0], ts) {
			JSON(c, nil, eno.InvalidSignErr)
			c.Abort()
			return
		}
	}
}

func CheckSign(sign, appid, path, ts string) (ok bool) {
	hash := md5.New()
	hash.Write([]byte(appid))
	hash.Write([]byte(path))
	hash.Write([]byte(ts))
	calSign := hex.EncodeToString(hash.Sum(nil))
	return sign == calSign
}
