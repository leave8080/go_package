package web

import (
	"fmt"
	"github/leave8080/go_package/pkg/eno"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	TypeOctetStream = "application/octet-stream"
	TypeForm        = "application/x-www-form-urlencoded"
	TypeJson        = "application/json"
	TypeXml         = "application/xml"
	TypeJpg         = "image/jpeg"
	TypePng         = "image/png"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, data interface{}, err error) {
	e := eno.AnalyseError(err)

	rsp := BaseResponse{
		Code:    e.Code(),
		Message: e.Message(),
		Data:    data,
	}
	c.JSON(http.StatusOK, rsp)
}

func Redirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
}

func File(c *gin.Context, fileBytes []byte, fileName, fileType string) {
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	var contentType string
	switch fileType {
	case "jpg", "jpeg":
		contentType = TypeJpg
	case "png":
		contentType = TypePng
	default:
		contentType = TypeOctetStream
	}
	c.Data(http.StatusOK, contentType, fileBytes)
}
