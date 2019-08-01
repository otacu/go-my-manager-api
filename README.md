# go-my-manager-api
[我的后台管理系统](https://github.com/otacu/vue-my-manager)的后台Api。
用echo框架做api，用gorm框架操作数据库。
后台用户角色权限等系统模块由后台api直连数据库，后续其它业务采用rpc调用其它微服务的方式。
## 编译
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```
编译后会得到一个叫main的无后缀文件。
## 赋予权限
```
chmod +x main
```
## 启动
```
nohup ./main &
```
之后用netstat -tunlp |grep 端口号 查询进程，用ps -ef|grep go是查不到的。
