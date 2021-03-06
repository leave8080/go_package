package eno

import "sync"

var (
	ErrorMap = new(sync.Map)
	// base error
	OK              = New(0, "SUCCESS")
	RequestErr      = New(10400, "Request Error")
	InvalidSignErr  = New(10401, "Invalid Sign")
	InvalidAppidErr = New(10402, "Invalid Appid")
	InvalidTokenErr = New(10403, "Invalid Token")
	NothingFound    = New(10404, "Nothing Found")
	ServerErr       = New(10500, "Server Error")
	ServerBusyErr   = New(10503, "服务器忙，请稍后重试")
)
