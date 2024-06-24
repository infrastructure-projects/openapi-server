# openapi-server
openapi聚合服务器，用于聚合各微服务的openapi文档

#### 一、什么是聚合文档服务?
在当前的微服务架构中，所有的`OpenAPI`文档全都维护在不同的服务中，如果想要查看某个服务的文档，非常的繁琐。
`openapi-server`将所有微服务的文档全都收集起来，使得开发工程师/测试工程师等只需要通过一个入口就能查询到所有微服务的文档信息。

#### 二、部署方式
假设当前工作路径为`/opt`,现需要将`service-demo1`和`service-demo2`的`OpenAPI`文档收集起来.
该目录的结构应为:
```text
.
├── app
└── conf
     └── application.yaml
```
在`application.yaml`中定义服务配置
```yaml
application:
    global-endpoint: /v3/api-docs
    urls:
      - name: service-demo1
        url: "/v1/service-demo1"
      - name: service-demo1
        url: "/v1/service-demo2"
        endpoint: /v2/api-docs # 使用单独的endpoint
```

## 三、编译命令

### 3.1 Windows

```shell
GOOS=windows GOARCH=amd64 go build app.ext cmd/main.go
```

### 3.2 Linux

```shell
GOOS=linux GOARCH=amd64 go build -o app cmd/main.go
```

### 3.2 MacOS

x86:

```shell
GOOS=darwin GOARCH=amd64 go build -o app cmd/main.go
```

ARM:

```shell
GOOS=darwin GOARCH=arm64 go build -o app cmd/main.go
```
