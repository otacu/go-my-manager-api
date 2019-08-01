package main

// query是指get请求url中的参数
type LoginRequest struct {
	UserName string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}
