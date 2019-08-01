package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go-my-manager-api/config/jwt_config"
)

func Login(c echo.Context) error {
	r := new(LoginRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	// 查询用户
	var managerUser User
	if DB.Where(&User{UserName: r.UserName, Password: r.Password}).First(&managerUser).RecordNotFound() {
		return echo.ErrUnauthorized
	}

	// 设置jwt自定义参数
	claims := &jwt_config.JwtCustomClaims{
		// 自定义参数 name
		managerUser.Name,
		// 设置父类参数值
		jwt.StandardClaims{
			// 超时时间15分钟
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	// 用上自定义参数创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwt_config.SECRET))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func UserInfo(c echo.Context) error {
	// 从token中取出用户名
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwt_config.JwtCustomClaims)
	name := claims.Name
	// 查询用户
	var managerUser User
	DB.First(&managerUser, "name = ?", name)
	u := new(UserInfoResponse)
	u.Name = managerUser.Name
	u.Avatar = managerUser.Avatar
	u.Roles = []string{}
	if managerUser.Roles != "" {
		u.Roles = strings.Split(managerUser.Roles, ",")
	}
	return c.JSON(http.StatusOK, u)
}
