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

## 辅助库/组件
* 包依赖管理：[govendor](https://github.com/kardianos/govendor)
* Makefile：内置快速使用的命令

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
* [inhere/go-gin-skeleton](https://github.com/inhere/go-gin-skeleton)
