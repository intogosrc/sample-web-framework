/*
 @Title http 接收安全参数
 @Description
 @Author  Leo
 @Update  2020/8/6 2:20 下午
*/

package rpchelper

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func RequestParameterExists(c *gin.Context, key string) (string, bool) {
	v,ok := c.GetQuery(key)

	if !ok {
		v,ok = c.GetPostForm(key)
	}

	if !ok {
		return "",false
	}

	return v, true
}

// 获取字符串值
func RequestParameterString(c *gin.Context, key string) string {
	v := c.Query(key)

	if v == "" {
		v = c.PostForm(key)
	}

	return v
}

// 获取整数值
func RequestParameterInt(c *gin.Context, key string) (int64, bool) {
	v := RequestParameterString(c, key)

	if v == "" {
		return 0, false
	}

	i, e := strconv.ParseInt(v, 10, 64)
	if e != nil {
		return 0, false
	}

	return i, true
}

// 获取浮点值
func RequestParameterFloat(c *gin.Context, key string) (float64, bool) {
	v := RequestParameterString(c, key)

	if v == "" {
		return 0, false
	}

	i, e := strconv.ParseFloat(v, 64)
	if e != nil {
		return 0, false
	}

	return i, true
}