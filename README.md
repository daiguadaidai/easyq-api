# easyq-api

数据库查询服务平台api服务

### 技术栈

语言: `Golang`

API框架: `gin`

数据库框架: `gorm`

命令行框架: `cobra`

数据库: `MySQL`

Go版本: `go version go1.15.15 darwin/amd64`

### 启动方法

#### 方法一: 配置文件启动(推荐)

```
# 编译
cd $GOPATH/src/github.com/daiguadaidai/easyq-api
go build

# 通过配置文件启动
./easyq-api --config=./easyq_api.toml
```
