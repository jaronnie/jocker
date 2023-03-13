# jocker
A mini Docker written in Go used to learn docker.

## 技术栈

* cobra(实现 jocker cli)
* go-zero(实现 jockerd)
* grpc-gateway(提供 grpc 的 http 接口)
* gin
* protoc-gen-go-httpsdk(自动生成 jockerd 的 go sdk 库)
* vue(jocker 内置 vue 实现的 jocker ui)

## jocker jockerd

jockerd 即 jocker 的服务端, 支持的协议:
* unix-socket
* tcp
* ssh

```shell
# 启动 jockerd, 默认采用 unix-socket 进行通信
jocker jockerd
```

## jocker cli

```shell
jocker ps
jocker run
jocker exec
```
