package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	ResponseCodeOk       = 100000
	ResponseCodeTokenErr = 300000
	ResponseCodeErr      = 500000
)

type response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type listData struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

// 返回错误
func ReturnError(c *gin.Context, code int, err error) {
	resp := &response{Data: nil, Code: code, Success: false, Message: err.Error()}
	c.JSON(http.StatusOK, resp)
}

// 返回列表数据的错误
func ReturnListError(c *gin.Context, code int, err error) {
	listData := &listData{List: make([]interface{}, 0), Total: 0}
	resp := &response{Data: listData, Code: code, Success: false, Message: err.Error()}
	c.JSON(http.StatusOK, resp)
}

// 成功并放回数据
func ReturnSuccess(c *gin.Context, obj interface{}) {
	resp := &response{Data: obj, Code: ResponseCodeOk, Success: true, Message: "success"}
	c.JSON(http.StatusOK, resp)
}

// 成功并放回数据
func ReturnList(c *gin.Context, obj interface{}, total int) {
	listData := &listData{List: obj, Total: total}
	resp := &response{Data: listData, Code: ResponseCodeOk, Success: true, Message: "success"}
	c.JSON(http.StatusOK, resp)
}

// 返回文件信息
func ReturnFile(c *gin.Context, fileName string, path string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}

func ForwardRequest(c *gin.Context, targetUrl string, timeout int64) {
	target, err := url.Parse(targetUrl)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, fmt.Errorf("解析需要转发请求的目标URL出错, 目标URL:%v. %v", targetUrl, err.Error()))
		return
	}
	c.Request.URL.Path = target.Path
	c.Request.URL.RawQuery = target.RawQuery
	target.Path = ""
	target.RawQuery = ""
	proxy := httputil.NewSingleHostReverseProxy(target)
	pTransport := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeout))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	proxy.Transport = pTransport
	proxy.ServeHTTP(c.Writer, c.Request)
}
