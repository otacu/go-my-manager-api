生成go文件命令
```
protoc -I ./ anime_proto.proto --go_out=plugins=grpc:./
```