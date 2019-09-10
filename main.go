package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-my-manager-api/config/jwt_config"
	"log"
)

const (
	dataSource             = "admin:123456@tcp(127.0.0.1:3306)/my-manager?charset=utf8&parseTime=true"
	malAnimeServiceAddress = "localhost:10097"
)

// 数据库连接池
var DB *gorm.DB

func main() {
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		log.Println(err)
		panic("连接数据库失败")
	}
	defer db.Close()
	DB = db

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/my-manager-api/user/login", Login)

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwt_config.JwtCustomClaims{},
		SigningKey: []byte(jwt_config.SECRET),
	}

	// 路径组 路径前缀
	r := e.Group("/my-manager-api")
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/user/info", UserInfo)
	r.POST("/mal-anime/search", Search)

	// 启动
	e.Logger.Fatal(e.Start(":8895"))
}
