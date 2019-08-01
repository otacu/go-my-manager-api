package main

type UserInfoResponse struct {
	Name   string   `json:"name" form:"name" query:"name"`
	Avatar string   `json:"avatar" form:"avatar" query:"avatar"`
	Roles  []string `json:"roles" form:"roles" query:"roles"`
}
