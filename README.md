# go-zero-boat  

go-zero是一个很好用的微服务框架，在使用的时候，希望能有一个快速启动，新项目拿来就能用的模板工程，于是 go-zero-boat 就诞生了，旨在轻便，容易理解，刚好够用。

### 依赖版本

|dependency|version|
|---|---|
|go-zero|1.5.0|
|goctl|1.5.0|

### 感谢
[go-zero](https://github.com/zeromicro/go-zero)     
[go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook)

### 快速开始
```
# 启动 etcd
$ etcd   
# 使用 modd 进行启动
$ modd
Starting a rpc server at 0.0.0.0:8080...
Starting server at 0.0.0.0:8888...
# 测试服务接口
$ curl --location 'localhost:8888/user/detail/0'
{"id":1,"name":"some lucky user..."}%
$ curl --location 'localhost:8888/user/test/err'
{"code":10001,"msg":"testErr Msg"}%
```

### goctl生成代码
+ **gateway** 编写 gateway/apis/*.api 文件，import到 gateway/index.api，执行 `bash gen-api.sh`.
+ **rpc** 在 rpc 目录下找到服务目录，编写 rpc/**/*.proto 文件，在 `gen-rpc.sh` 中修改 `target` 变量为服务名，执行 `bash gen-api.sh` 


#### 目录
```
.  
├── README.md  
├── common  
│   ├── interceptor  
│   │   └── rpclogger.go  
│   └── xerr  
│       ├── handler.go  
│       ├── xerrCode.go  
│       ├── xerrMsg.go  
│       └── xerrs.go  
├── deploy  
│   ├── goctl      // 自定义goctl模板  
│   │   ├── 1.5.0  
│   │   └── CHANGE.md  
│   └── script     // 初始化sql脚本  
├── gatway   // 网关API  
│   ├── apis  
│   │   └── user.api  
│   ├── gen-api.sh    // goctl生成api目录  
│   ├── index.api  
│   └── logs  
├── go.mod  
├── go.sum  
└── rpc  // rpc服务  
    ├── gen-model.sh   // goctl根据SQL生成model  
    ├── gen-rpc.sh     // goctl根据proto生成rpc目录  
    ├── pet  
    │   └── pet.proto  
    └── user  
         └── user.proto  
```

#### 网关
在考虑rpc网关的时候，参考 [looklook](https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/doc/chinese/02-nginx%E7%BD%91%E5%85%B3.md) 使用nginx做网关是一种很好的思路，但是这里依然选择使用api做网关，在api层可以考虑做统一鉴权，日志收集，限流等等，可以跟rpc服务做一个切割。  
这样的选择，使得业务系统都以rpc服务的形式进行开发，系统主体就分成了 `gateway` 和 `rpc` 两条线。
业务按模块拆分放在 `rpc` 下，统一由 `gateway` 接入来提供http能力。


#### 错误处理
借鉴 looklook 项目的思想，定义rpc和api共用的errCode模块，稍有不同的是looklook项目采用在handler模板代码处修改，本着goctl 原生模板可用尽量不动的前提，结合go-zero官网的错误处理拦截器，最终将错误处理放在了 `httpx.SetErrorHandlerCtx(xerr.Handler)`

#### 热加载
1. 使用 [air](https://github.com/cosmtrek/air) 配置示例如下
   ```
    #!/bin/bash
    # target project
    target=user
    # rpc start
    cd ${target}
    air --build.cmd "go build -o ${target} ${target}.go" --build.bin "./${target}" 
   ```
   缺点，暂不支持同时启动多个项目，需要自己写多个cmd来实现 [#160](https://github.com/cosmtrek/air/issues/160)
2. 使用 [modd](https://github.com/cortesi/modd) 配置示例如下
   ```
   # modd.conf
   rpc/user/**/*.go {
    prep: go build -o build/server/user-rpc  -v rpc/user/user.go
    daemon +sigkill: ./build/server/user-rpc -f rpc/user/etc/user.yaml
    }

    gateway/**/*.go {
        prep: go build -o build/server/gateway  -v gateway/gateway.go
        daemon +sigkill: ./build/server/gateway -f gateway/etc/gateway.yaml
    }
   ```
这里推荐使用 `modd` ，对于多服务支持好一些
