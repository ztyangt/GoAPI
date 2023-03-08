# goAPI

基于 Golang Gin 框架开的 API 服务

## windows 编译 linux 可执行文件

在 cmd 或者 powershell 执行以下命令

```bash
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
```

然后正常 go build 即可

再切换回 windows

```bash
go env -w GOARCH=amd64
go env -w GOOS=windows
```
