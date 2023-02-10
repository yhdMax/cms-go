package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"vue-next-admin-go/dao"
	"vue-next-admin-go/modles"
	"vue-next-admin-go/utils"
)

func Login(c *gin.Context) {
	response := modles.Response{}
	response.Code = 204
	// 开启session
	session := sessions.Default(c)

	payload := new(dao.UserInfoModel)
	if err := c.ShouldBindJSON(payload); err != nil {
		response.Error = utils.PramsError
		c.JSON(http.StatusOK, response)
		return
	}
	if payload.UserName == "" || payload.PassWord == "" || payload.Code == "" {
		response.Error = utils.PramsError
		c.JSON(http.StatusOK, response)
		return
	}

	payload.UserName = strings.TrimSpace(payload.UserName)
	payload.PassWord = strings.TrimSpace(payload.PassWord)
	if CodeVerify(payload.Code) {
		info, err := payload.Info()

		if err != nil {
			response.Error = "账号不存在，请注册"
			c.JSON(http.StatusOK, response)
			return
		}

		if info.Status != "active" {
			response.Error = "此账号已被关闭"
			c.JSON(http.StatusOK, response)
			return
		}

		// 初始化生成 token 配置
		claims := Claims
		claims["id"] = info.ID
		claims["user_name"] = info.Name
		claims["standard_claims"] = jwt.StandardClaims{
			ExpiresAt: time.Now().Add(300 * time.Second).Unix(), // 过期时间
			Issuer:    "next_admin",
		}

		token, err := GenerateToken(claims)

		if err != nil {
			response.Error = "生成token失败"
			c.JSON(http.StatusOK, response)
			return
		}

		// token 存入session 并返回
		session.Set("admin_token", token)
		err = session.Save()
		if err != nil {
			return
		}

		response.Data = modles.Token{
			Token: token,
		}
		response.Code = 200
		c.JSON(http.StatusOK, response)
		return
	} else {
		response.Code = 205
		response.Error = utils.CaptchaCodeError
		c.JSON(http.StatusOK, response)
		return
	}
}

func SignIn(c *gin.Context) {
	token, _ := c.Cookie("token")
	isToken := Verification(c, token)

	c.JSON(http.StatusOK, isToken)

	return
}
