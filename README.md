# Albedo: Gin Api Project Skeleton
![s-2a.png](https://i.loli.net/2018/09/17/5b9f7028c8e77.png)
<br>

基于 [Gin](https://github.com/gin-gonic/gin) 的 API 开发框架

## 使用的包
* web 框架：[Gin](https://github.com/gin-gonic/gin)
* db：[gorm](https://github.com/jinzhu/gorm)
* json 序列化库：[json-iterator/go](https://github.com/json-iterator/go)
* 配置管理：[viper](https://github.com/spf13/viper)
* 日志记录：[onelog](https://github.com/francoispqt/onelog)
* 命令行：[cobra](https://github.com/spf13/cobra)
* 平滑重启：[endless](https://github.com/fvbock/endless)
* Redis：[go-redis/redis](github.com/go-redis/redis)

## 辅助库/组件
* 包依赖管理：[govendor](https://github.com/kardianos/govendor)
* Makefile：内置快速使用的命令

## 安装
```bash
$ cd $GOPATH/src/github.com/your_name

# 打开 [albedo-installer](https://gitee.com/zqhong/albedo-installer)，下载对应操作系统的安装器

# 以 macOS 系统为例，这里的项目名为 demo
$ chmod +x installer-darwin
$ ./installer-darwin -project-name=demo

$ cd demo

# 安装 golang 的第三方依赖包
# 请确保 $GOPATH/bin 目录在环境变量 PATH 下
$ make deps

$ make build

# 复制并修改配置文件
$ cp conf/config.yaml.example conf/config.yaml

# 测试
$ ./build/web
$ ./build/cli

# 如果启动失败，检查 runtime/albedo.log 日志文件
```

## Makefile 内置命令
```bash
# 同时编译 web 和 cli
$ make build

$ make build-web
$ make build-cli

$ make run-web
$ make -run-cli

$ make test

$ make clean

# 安装依赖
$ make deps
```

## admin.sh 内置命令
```bash
$ bash admin.sh start

$ bash admin.sh stop

# 利用 endless 实现平滑重启
$ bash admin restart

$ bash admin status
```

## 目录结构
```
├── admin.sh                     # 进程管理脚本（start/stop/restart/status）
├── app                          # 进程启动 && 服务初始化代码
│   ├── boot.go                  # 进程启动代码
│   ├── config.go                # 配置服务初始化
│   ├── db.go                    # db 服务初始化
│   ├── gin.go                   # gin 服务初始化
│   ├── logger.go                # logger 服务初始化
├── build                        # 存放编译后的二进制文件
├── conf                         # 配置文件存储目录
│   ├── config.yaml              # 配置文件
├── handler                      # MVC 架构中的 C
│   ├── handler.go
│   ├── sd                       # 健康检查handler
│   │   └── check.go
├── model                        # MVC 架构中的 M
├── pkg                          # 引用的包
├── router                       # 路由相关处理
├── runtime                      # 存放临时文件，如日志、pid 文件
├── util                         # 工具类函数存放目录
└── vendor                       # vendor目录用来管理依赖包
```

## 参考
* [lexkong/apiserver](https://github.com/lexkong/apiserver)
* [chenhg5/morningo](https://github.com/chenhg5/morningo)
* [cg33/morningo-installer](https://gitee.com/cg33/morningo-installer)
* [inhere/go-gin-skeleton](https://github.com/inhere/go-gin-skeleton)
