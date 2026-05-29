# HTTP 服务开发模式

due 框架的 HTTP 组件是对 [fiber](https://github.com/gofiber/fiber) 框架的二次封装，保留 fiber 的全部优势，同时提供符合 due 框架规范的调用接口。

## 目录

- [快速开始](#快速开始)
- [服务器配置](#服务器配置)
- [路由系统](#路由系统)
- [中间件](#中间件)
- [请求处理](#请求处理)
- [响应格式](#响应格式)
- [Swagger 集成](#swagger-集成)
- [微服务集成](#微服务集成)
- [最佳实践](#最佳实践)

## 快速开始

### 安装

```bash
go get github.com/dobyte/due/component/http/v2@latest
```

### 最小示例

```go
package main

import (
    "fmt"
    "github.com/dobyte/due/component/http/v2"
    "github.com/dobyte/due/v2"
    "github.com/dobyte/due/v2/codes"
    "github.com/dobyte/due/v2/log"
    "github.com/dobyte/due/v2/utils/xtime"
)

func main() {
    container := due.NewContainer()
    component := http.NewServer()
    initApp(component.Proxy())
    container.Add(component)
    container.Serve()
}

func initApp(proxy *http.Proxy) {
    router := proxy.Router()
    router.Get("/greet", greetHandler)
}

type greetReq struct {
    Message string `json:"message"`
}

type greetRes struct {
    Message string `json:"message"`
}

func greetHandler(ctx http.Context) error {
    req := &greetReq{}
    if err := ctx.Bind().JSON(req); err != nil {
        return ctx.Failure(codes.InvalidArgument)
    }
    log.Info(req.Message)
    return ctx.Success(&greetRes{
        Message: fmt.Sprintf("当前时间: %s", xtime.Now().Format(xtime.DateTime)),
    })
}
```

## 服务器配置

### 创建方式

采用函数选项模式 (Functional Options)：

```go
s := http.NewServer(
    http.WithName("my-api"),
    http.WithAddr(":9090"),
    http.WithConsole(true),
    http.WithBodyLimit(8 * 1024 * 1024),
)
```

### 配置选项

| 配置项 | 默认值 | 说明 |
| --- | --- | --- |
| name | `"http"` | 服务器名称 |
| addr | `":8080"` | 监听地址 |
| bodyLimit | 4MB | 请求体大小限制 |
| concurrency | 262144 | 最大并发连接数 |
| readBufferSize | 4096 | 读缓冲区大小 |
| writeBufferSize | 4096 | 写缓冲区大小 |

### TLS 配置

```go
s := http.NewServer(
    http.WithCredentials("cert.pem", "key.pem"),
)
```

### CORS 跨域配置

```go
s := http.NewServer(
    http.WithCorsOptions(http.CorsOptions{
        Enable:           true,
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: false,
        MaxAge:           3600,
    }),
)
```

### YAML 配置文件

```yaml
http:
  name: "http"
  addr: ":8080"
  console: true
  bodyLimit: "4M"
  concurrency: 262144

  cors:
    enable: true
    allowOrigins: ["*"]
    allowMethods: ["GET", "POST", "HEAD", "PUT", "DELETE", "PATCH"]
    allowHeaders: []
    allowCredentials: false
    maxAge: 0

  swagger:
    enable: true
    title: "API文档"
    basePath: "/swagger"
    filePath: "./docs/swagger.json"
```

## 路由系统

### Router 接口

```go
type Router interface {
    Get(path string, handlers ...any) Router
    Post(path string, handlers ...any) Router
    Head(path string, handlers ...any) Router
    Put(path string, handlers ...any) Router
    Delete(path string, handlers ...any) Router
    Connect(path string, handlers ...any) Router
    Options(path string, handlers ...any) Router
    Trace(path string, handlers ...any) Router
    Patch(path string, handlers ...any) Router
    All(path string, handlers ...any) Router
    Add(methods []string, path string, handlers ...any) Router
    Group(prefix string, middlewares ...any) Router
}
```

### 获取路由器

```go
proxy := server.Proxy()
router := proxy.Router()
```

### 多风格 Handler 支持

due HTTP 组件支持多种风格的处理器：

**due 风格（推荐）**:

```go
func handler(ctx http.Context) error {
    return ctx.Success("Hello")
}
router.Get("/api", handler)
```

**fiber 原生风格**:

```go
func handler(ctx fiber.Ctx) error {
    return ctx.SendString("Hello")
}
router.Get("/api", handler)
```

**net/http 风格**:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
}
router.Get("/api", handler)
```

**fasthttp 风格**:

```go
func handler(ctx *fasthttp.RequestCtx) {
    ctx.SetBody([]byte("Hello"))
}
router.Get("/api", handler)
```

### 路由组

```go
router := proxy.Router()

// API 路由组
api := router.Group("/api")

// V1 版本
v1 := api.Group("/v1")
v1.Get("/users", getUsers)
v1.Post("/users", createUser)

// V2 版本
v2 := api.Group("/v2")
v2.Get("/users", getUsersV2)
```

### 路由组中间件

```go
authMiddleware := func(ctx http.Context) error {
    token := ctx.Get("Authorization")
    if token == "" {
        return ctx.Failure(codes.Unauthorized)
    }
    return ctx.Next()
}

api := router.Group("/api", authMiddleware)
```

## 中间件

### 全局中间件

```go
s := http.NewServer(
    http.WithMiddlewares(
        // due 风格中间件
        func(ctx http.Context) error {
            start := time.Now()
            err := ctx.Next()
            log.Info("请求耗时: %v", time.Since(start))
            return err
        },
        // fiber 原生中间件
        someFiberMiddleware,
    ),
)
```

### 内置中间件

1. **Logger**: 仅当 `WithConsole(true)` 时启用
2. **Recover**: 始终启用，带堆栈跟踪
3. **CORS**: 当 `CorsOptions.Enable == true` 时启用
4. **Swagger UI**: 当 `SwagOptions.Enable == true` 时启用

### 自定义中间件

```go
// 请求ID中间件
func RequestID() http.Handler {
    return func(ctx http.Context) error {
        id := ctx.Get("X-Request-ID")
        if id == "" {
            id = uuid.New().String()
        }
        ctx.Set("X-Request-ID", id)
        return ctx.Next()
    }
}

// 日志中间件
func Logger() http.Handler {
    return func(ctx http.Context) error {
        log.Info("请求: %s %s", ctx.Method(), ctx.Path())
        return ctx.Next()
    }
}
```

## 请求处理

### Context 接口

```go
type Context interface {
    fiber.Ctx                              // 继承 fiber 全部能力
    CTX() fiber.Ctx                        // 获取底层 fiber.Ctx
    Proxy() *Proxy                         // 获取代理 API
    Failure(rst any) error                 // 统一失败响应
    Success(data ...any) error             // 统一成功响应
    StdRequest() *http.Request             // 获取标准 net/http 请求
}
```

### 请求绑定

```go
func createUser(ctx http.Context) error {
    // JSON 绑定
    user := &User{}
    if err := ctx.Bind().JSON(user); err != nil {
        return ctx.Failure(codes.InvalidArgument)
    }

    // 查询参数绑定
    page := ctx.Query("page", "1")
    limit := ctx.Query("limit", "10")

    // 路径参数
    id := ctx.Params("id")

    // 表单数据
    name := ctx.FormValue("name")

    return ctx.Success(user)
}
```

### 获取请求信息

```go
func handler(ctx http.Context) error {
    // 请求方法
    method := ctx.Method()

    // 请求路径
    path := ctx.Path()

    // 请求头
    contentType := ctx.Get("Content-Type")

    // Cookie
    token := ctx.Cookies("token")

    // 客户端IP
    ip := ctx.IP()

    return ctx.Success(nil)
}
```

## 响应格式

### 统一响应结构

```go
type Resp struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`  // 仅开发模式
    Data    any    `json:"data,omitempty"`
}
```

### 成功响应

```go
// 无数据
ctx.Success()
// {"code": 0, "message": "OK"}

// 有数据
ctx.Success(userData)
// {"code": 0, "message": "OK", "data": {...}}
```

### 失败响应

```go
// 使用错误码
ctx.Failure(codes.InvalidArgument)
// {"code": 400, "message": "Invalid argument"}

// 使用 error
ctx.Failure(errors.New("用户不存在"))
// {"code": 500, "message": "用户不存在"}
```

### 自定义响应

```go
// 设置状态码
ctx.Status(201).JSON(data)

// 设置响应头
ctx.Set("X-Custom-Header", "value")

// 重定向
ctx.Redirect("/new-location", 301)
```

## Swagger 集成

### 配置

```go
s := http.NewServer(
    http.WithSwagOptions(http.SwagOptions{
        Enable:   true,
        Title:    "API文档",
        BasePath: "/swagger",
        FilePath: "./docs/swagger.json",
    }),
)
```

### 注解

```go
// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body CreateUserReq true "请求参数"
// @Response 200 {object} http.Resp{Data=User} "响应参数"
// @Router /api/v1/users [post]
func createUser(ctx http.Context) error {
    // ...
}
```

### 生成文档

```bash
# 安装 swag
go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
swag init --parseDependency

# 删除自动生成的 docs.go（可选）
rm -rf ./docs/docs.go
```

## 微服务集成

### 创建 Mesh 客户端

```go
func handler(ctx http.Context) error {
    // 直连（IP）
    client, err := ctx.Proxy().NewMeshClient("direct://127.0.0.1:8011")

    // 直连（实例ID）
    client, err := ctx.Proxy().NewMeshClient("direct://711baf8d-8a06-11ef-b7df-f4f19e1f0070")

    // 服务发现
    client, err := ctx.Proxy().NewMeshClient("discovery://user-service")

    if err != nil {
        return ctx.Failure(err)
    }

    // 使用 client 调用其他微服务...
    return ctx.Success(nil)
}
```

### 配置传输器

```go
s := http.NewServer(
    http.WithRegistry(myRegistry),
    http.WithTransporter(myTransporter),
)
```

## 最佳实践

### 1. 项目结构

```text
project/
├── cmd/
│   └── http/
│       └── main.go
├── internal/
│   ├── handler/        # HTTP 处理器
│   ├── middleware/      # 中间件
│   ├── model/          # 数据模型
│   └── service/        # 业务逻辑
├── docs/               # Swagger 文档
└── etc/
    └── etc.toml        # 配置文件
```

### 2. 错误处理

```go
func handler(ctx http.Context) error {
    user, err := service.GetUser(id)
    if err != nil {
        log.Error("获取用户失败: %v", err)
        return ctx.Failure(codes.InternalError)
    }
    return ctx.Success(user)
}
```

### 3. 参数验证

```go
type CreateUserReq struct {
    Name  string `json:"name" validate:"required,min=2,max=50"`
    Email string `json:"email" validate:"required,email"`
}

func createUser(ctx http.Context) error {
    req := &CreateUserReq{}
    if err := ctx.Bind().JSON(req); err != nil {
        return ctx.Failure(codes.InvalidArgument)
    }

    // 使用验证库
    if err := validate.Struct(req); err != nil {
        return ctx.Failure(codes.InvalidArgument)
    }

    // ...
}
```

### 4. 分层架构

```go
// handler 层 - 处理 HTTP 请求
func getUserHandler(ctx http.Context) error {
    id := ctx.Params("id")
    user, err := userService.GetUser(ctx.Context(), id)
    if err != nil {
        return ctx.Failure(err)
    }
    return ctx.Success(user)
}

// service 层 - 业务逻辑
type UserService struct {
    repo UserRepository
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    return s.repo.FindByID(ctx, id)
}
```

### 5. 中间件链

```go
s := http.NewServer(
    http.WithMiddlewares(
        RequestID(),
        Logger(),
        Recovery(),
        Auth(),
        RateLimit(100),
    ),
)
```

### 6. 静态文件服务

```go
router := proxy.Router()

// 静态文件
router.Static("/static", "./public")

// 单个文件
router.Get("/favicon", func(ctx http.Context) error {
    return ctx.SendFile("./public/favicon.ico")
})
```

### 7. WebSocket 支持

```go
router := proxy.Router()

router.Get("/ws", func(ctx http.Context) error {
    return ctx.Websocket(func(conn *websocket.Conn) {
        for {
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                break
            }
            conn.WriteMessage(msgType, msg)
        }
    })
})
```

## 常见问题

### Q: 如何获取原生 fiber.App？

```go
app := proxy.App()
```

### Q: 如何自定义上下文？

due 已通过 `fiber.NewWithCustomCtx` 注入自定义上下文，每个请求的 `ctx` 都是 `http.Context` 类型。

### Q: 如何处理文件上传？

```go
func upload(ctx http.Context) error {
    file, err := ctx.FormFile("file")
    if err != nil {
        return ctx.Failure(codes.InvalidArgument)
    }

    // 保存文件
    err = ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
    if err != nil {
        return ctx.Failure(codes.InternalError)
    }

    return ctx.Success(nil)
}
```

### Q: 如何设置超时？

```go
s := http.NewServer(
    http.WithReadTimeout(5 * time.Second),
    http.WithWriteTimeout(10 * time.Second),
    http.WithIdleTimeout(120 * time.Second),
)
```

## 参考资料

- [fiber 官方文档](https://docs.gofiber.io/)
- [due GitHub](https://github.com/dobyte/due)
- [due HTTP 组件源码](https://github.com/dobyte/due/tree/main/component/http)
