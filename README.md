# GoAPI

基于 Golang Gin 框架开的 API 服务

# 当前进度

1. 音乐解析类 API
   1. 网易云音乐

# 相关命令

### windows 编译 linux 可执行文件

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

### 工作方式

1. 开发模式

```bash
gin.SetMode(gin.DebugMode)
```

2. 生产环境模式

```bash
gin.SetMode(gin.ReleaseMode)
```

3. 测试环境模式

```bash
gin.SetMode(gin.TestMode)
```
