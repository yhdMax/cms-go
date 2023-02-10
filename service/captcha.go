package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

// 开辟验证码存储空间	DefaultMemStore(大小，失效时间)
var store = base64Captcha.DefaultMemStore
var Id string
var B64s string

func InitCode() (id string, b64s string, err error) {
	// 使用默认验证码样式
	var driver = base64Captcha.NewDriverDigit(44, 132, 4, 0.1, 2)
	// 生成验证码，同时保存至store
	var cd = base64Captcha.NewCaptcha(driver, store)
	// 生成base64图像验证及id
	id, b64s, err = cd.Generate()

	if err != nil {
		fmt.Println("sessionError:", err)
		return
	}
	return
}

// SetCaptcha 生成验证码
func SetCaptcha(c *gin.Context) {
	//session := sessions.Default(c)
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	Id, B64s, _ = InitCode()
	if Id == "" && B64s == "" {
		fmt.Println("Error: init_code_error")
		return
	}
	c.String(http.StatusOK, B64s)
	return
}

// CodeVerify 校验code
func CodeVerify(code string) bool {
	if Id == "" || B64s == "" {
		return false
	}
	// 校验验证码 Verify(id, code, 是否在内存清除code图片)
	return store.Verify(Id, code, true)
}
