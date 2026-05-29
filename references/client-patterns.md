# 客户端开发模式

due 框架的客户端组件采用分层设计，支持 TCP 和 WebSocket 两种协议，面向游戏服务器的连接需求。

## 目录

- [快速开始](#快速开始)
- [架构概览](#架构概览)
- [TCP 客户端](#tcp-客户端)
- [WebSocket 客户端](#websocket-客户端)
- [集群层客户端](#集群层客户端)
- [连接管理](#连接管理)
- [消息收发](#消息收发)
- [事件处理](#事件处理)
- [最佳实践](#最佳实践)

## 快速开始

### 安装

```bash
# TCP 客户端
go get github.com/dobyte/due/network/tcp/v2@latest

# WebSocket 客户端
go get github.com/dobyte/due/network/ws/v2@latest

# 集群层客户端
go get github.com/dobyte/due/cluster/client/v2@latest
```

### TCP 客户端最小示例

```go
package main

import (
    "github.com/dobyte/due/network/tcp/v2"
    "github.com/dobyte/due/v2/log"
    "github.com/dobyte/due/v2/network"
    "github.com/dobyte/due/v2/packet"
)

func main() {
    client := tcp.NewClient()

    client.OnConnect(func(conn network.Conn) {
        log.Info("连接已建立")
    })
    client.OnDisconnect(func(conn network.Conn) {
        log.Info("连接已断开")
    })
    client.OnReceive(func(conn network.Conn, data []byte) {
        msg, _ := packet.UnpackMessage(data)
        log.Infof("收到消息: seq=%d route=%d", msg.Seq, msg.Route)
    })

    conn, err := client.Dial()
    if err != nil {
        log.Fatal("拨号失败:", err)
    }
    defer conn.Close()

    // 发送消息
    msg, _ := packet.PackMessage(&packet.Message{
        Seq:    1,
        Route:  1,
        Buffer: []byte("hello"),
    })
    conn.Push(msg)

    select {} // 保持运行
}
```

### WebSocket 客户端最小示例

```go
package main

import (
    "github.com/dobyte/due/network/ws/v2"
    "github.com/dobyte/due/v2/log"
    "github.com/dobyte/due/v2/network"
    "github.com/dobyte/due/v2/packet"
)

func main() {
    client := ws.NewClient()

    client.OnConnect(func(conn network.Conn) {
        log.Info("连接已建立")
    })
    client.OnDisconnect(func(conn network.Conn) {
        log.Info("连接已断开")
    })
    client.OnReceive(func(conn network.Conn, data []byte) {
        msg, _ := packet.UnpackMessage(data)
        log.Infof("收到消息: seq=%d route=%d", msg.Seq, msg.Route)
    })

    conn, err := client.Dial()
    if err != nil {
        log.Fatal("拨号失败:", err)
    }
    defer conn.Close()

    // 发送消息
    msg, _ := packet.PackMessage(&packet.Message{
        Seq:    1,
        Route:  1,
        Buffer: []byte("hello"),
    })
    conn.Push(msg)

    select {} // 保持运行
}
```

## 架构概览

客户端组件采用三层架构：

```text
┌─────────────────────────────────────┐
│        集群层 (cluster/client)       │  ← 业务使用层
│  - 消息编解码 (codec)                │
│  - 消息加解密 (encryptor)            │
│  - 路由分发                          │
│  - 事件管理                          │
├─────────────────────────────────────┤
│        网络层 (network)              │  ← 接口定义层
│  - Client 接口                       │
│  - Conn 接口                         │
├─────────────────────────────────────┤
│   TCP (network/tcp) │ WS (network/ws) │  ← 协议实现层
└─────────────────────────────────────┘
```

### 核心接口

**network.Client**:

```go
type Client interface {
    Dial(addr ...string) (Conn, error)
    Protocol() string
    OnConnect(handler ConnectHandler)
    OnReceive(handler ReceiveHandler)
    OnDisconnect(handler DisconnectHandler)
}
```

**network.Conn**:

```go
type Conn interface {
    ID() int64
    UID() int64
    Attr() Attr
    Bind(uid int64)
    Unbind()
    Send(msg []byte) error      // 同步发送
    Push(msg []byte) error      // 异步发送
    State() ConnState
    Close(force ...bool) error
    LocalIP() (string, error)
    RemoteIP() (string, error)
}
```

## TCP 客户端

### 创建方式

```go
// 默认配置
client := tcp.NewClient()

// 自定义配置
client := tcp.NewClient(
    tcp.WithClientAddr("192.168.1.100:3553"),
    tcp.WithClientDialTimeout(5 * time.Second),
    tcp.WithClientWriteTimeout(3 * time.Second),
    tcp.WithClientWriteQueueSize(2048),
    tcp.WithClientHeartbeatInterval(15 * time.Second),
    tcp.WithClientCredentials("ca.crt", "server.example.com"),
)
```

### 配置选项

| 选项函数 | 配置键 | 默认值 | 说明 |
| --- | --- | --- | --- |
| `WithClientAddr` | `etc.network.tcp.client.addr` | `127.0.0.1:3553` | 拨号地址 |
| `WithClientCredentials` | `etc.network.tcp.client.caFile` | 空 | TLS CA 证书 |
| `WithClientDialTimeout` | `etc.network.tcp.client.dialTimeout` | `3s` | 拨号超时 |
| `WithClientWriteTimeout` | `etc.network.tcp.client.writeTimeout` | `0s` | 写入超时 |
| `WithClientWriteQueueSize` | `etc.network.tcp.client.writeQueueSize` | `1024` | 写队列大小 |
| `WithClientHeartbeatInterval` | `etc.network.tcp.client.heartbeatInterval` | `10s` | 心跳间隔 |

### YAML 配置

```yaml
network:
  tcp:
    client:
      addr: "127.0.0.1:3553"
      dialTimeout: "3s"
      writeTimeout: "0s"
      writeQueueSize: 1024
      heartbeatInterval: "10s"
      # caFile: "ca.crt"
      # serverName: "server.example.com"
```

### TLS 连接

```go
client := tcp.NewClient(
    tcp.WithClientAddr("secure.server.com:3553"),
    tcp.WithClientCredentials("ca.crt", "secure.server.com"),
)
```

## WebSocket 客户端

### 创建方式

```go
// 默认配置
client := ws.NewClient()

// 自定义配置
client := ws.NewClient(
    ws.WithClientUrl("ws://192.168.1.100:3553"),
    ws.WithClientDialTimeout(5 * time.Second),
    ws.WithClientWriteTimeout(3 * time.Second),
    ws.WithClientWriteQueueSize(2048),
    ws.WithClientHeartbeatInterval(15 * time.Second),
)
```

### 配置选项

| 选项函数 | 配置键 | 默认值 | 说明 |
| --- | --- | --- | --- |
| `WithClientUrl` | `etc.network.ws.client.url` | `ws://127.0.0.1:3553` | WebSocket URL |
| `WithClientDialTimeout` | `etc.network.ws.client.dialTimeout` | `3s` | 握手超时 |
| `WithClientWriteTimeout` | `etc.network.ws.client.writeTimeout` | `0s` | 写入超时 |
| `WithClientWriteQueueSize` | `etc.network.ws.client.writeQueueSize` | `1024` | 写队列大小 |
| `WithClientHeartbeatInterval` | `etc.network.ws.client.heartbeatInterval` | `10s` | 心跳间隔 |

### YAML 配置

```yaml
network:
  ws:
    client:
      url: "ws://127.0.0.1:3553"
      dialTimeout: "3s"
      writeTimeout: "0s"
      writeQueueSize: 1024
      heartbeatInterval: "10s"
```

### WSS (TLS) 连接

```go
client := ws.NewClient(
    ws.WithClientUrl("wss://secure.server.com:3553"),
)
```

## 集群层客户端

集群层客户端是面向业务的高层封装，支持消息编解码、加解密、路由分发。

### 创建方式

```go
import (
    "github.com/dobyte/due/cluster/client/v2"
    "github.com/dobyte/due/encoding/json"
    "github.com/dobyte/due/network/tcp/v2"
)

// TCP 客户端
c := client.NewClient(
    client.WithName("my-client"),
    client.WithCodec(json.NewCodec()),
    client.WithClient(tcp.NewClient()),
)

// WebSocket 客户端
c := client.NewClient(
    client.WithName("my-client"),
    client.WithCodec(json.NewCodec()),
    client.WithClient(ws.NewClient()),
)
```

### 配置选项

| 选项函数 | 配置键 | 默认值 | 说明 |
| --- | --- | --- | --- |
| `WithID` | `etc.cluster.client.id` | UUID | 实例唯一 ID |
| `WithName` | `etc.cluster.client.name` | `client` | 实例名称 |
| `WithCodec` | `etc.cluster.client.codec` | `proto` | 编解码器 |
| `WithClient` | - | 必填 | 底层网络客户端 |
| `WithEncryptor` | - | 无 | 消息加解密器 |
| `WithContext` | - | `context.Background()` | 上下文 |

### YAML 配置

```yaml
cluster:
  client:
    id: ""  # 留空自动生成 UUID
    name: "my-client"
    codec: "json"  # json 或 proto
```

### 完整示例

```go
package main

import (
    "github.com/dobyte/due/cluster/client/v2"
    "github.com/dobyte/due/cluster/v2"
    "github.com/dobyte/due/encoding/json"
    "github.com/dobyte/due/network/tcp/v2"
    "github.com/dobyte/due/v2"
    "github.com/dobyte/due/v2/log"
)

type GreetReq struct {
    Message string `json:"message"`
}

type GreetRes struct {
    Message string `json:"message"`
}

func main() {
    container := due.NewContainer()

    c := client.NewClient(
        client.WithName("my-client"),
        client.WithCodec(json.NewCodec()),
        client.WithClient(tcp.NewClient()),
    )

    initApp(c.Proxy())
    container.Add(c)
    container.Serve()
}

func initApp(proxy *client.Proxy) {
    // 注册生命周期钩子
    proxy.AddHookListener(cluster.Start, func(p *client.Proxy) {
        // 启动后建立连接
        conn, err := p.Dial(client.WithDialAddr("127.0.0.1:3553"))
        if err != nil {
            log.Error("拨号失败:", err)
            return
        }

        // 发送消息
        conn.Push(&cluster.Message{
            Route: 1,
            Data:  &GreetReq{Message: "hello"},
        })
    })

    // 注册事件监听器
    proxy.AddEventListener(cluster.Connect, func(conn *client.Conn) {
        log.Info("连接已建立, cid:", conn.ID())
    })

    proxy.AddEventListener(cluster.Disconnect, func(conn *client.Conn) {
        log.Info("连接已断开, cid:", conn.ID())
    })

    // 注册路由处理器
    proxy.AddRouteHandler(1, func(ctx *client.Context) {
        res := &GreetRes{}
        if err := ctx.Parse(res); err != nil {
            log.Error("解析消息失败:", err)
            return
        }
        log.Info("收到响应:", res.Message)
    })

    // 设置默认路由处理器
    proxy.SetDefaultRouteHandler(func(ctx *client.Context) {
        log.Warn("未处理的路由:", ctx.Route())
    })
}
```

## 连接管理

### 连接状态

```go
const (
    ConnOpened  ConnState = 1  // 已打开
    ConnHanged  ConnState = 2  // 挂起中（优雅关闭）
    ConnClosed  ConnState = 3  // 已关闭
)
```

### 关闭连接

```go
// 优雅关闭（等待消息发送完毕）
conn.Close()

// 强制关闭
conn.Close(true)
```

### 连接属性

```go
// 设置属性
conn.Attr().Set("token", "abc123")

// 获取属性
token, ok := conn.Attr().Get("token")

// 删除属性
conn.Attr().Del("token")

// 遍历属性
conn.Attr().Visit(func(key, value any) bool {
    log.Info("attr:", key, "=", value)
    return true
})
```

### 用户绑定

```go
// 绑定用户 ID
conn.Bind(uid)

// 解绑
conn.Unbind()

// 获取绑定的用户 ID
uid := conn.UID()
```

### 拨号选项

```go
conn, err := proxy.Dial(
    // 指定拨号地址（覆盖默认配置）
    client.WithDialAddr("192.168.1.100:3553"),

    // 设置连接属性
    client.WithConnAttr("token", "abc123"),
    client.WithConnAttr("uid", 10001),
)
```

## 消息收发

### 消息结构

```go
type Message struct {
    Seq    uint32  // 序列号
    Route  uint32  // 路由 ID
    Data   any     // 消息数据（[]byte 或可序列化对象）
}
```

### 发送消息（集群层）

```go
// 发送结构体（自动序列化）
conn.Push(&cluster.Message{
    Route: 1001,
    Data:  &GreetReq{Message: "hello"},
})

// 发送字节数据
conn.Push(&cluster.Message{
    Route: 1001,
    Data:  []byte("raw data"),
})

// 带序列号
conn.Push(&cluster.Message{
    Seq:   1,
    Route: 1001,
    Data:  &GreetReq{Message: "hello"},
})
```

### 发送消息（网络层）

```go
// 打包消息
msg, err := packet.PackMessage(&packet.Message{
    Seq:    1,
    Route:  1,
    Buffer: []byte("hello"),
})

// 异步发送
conn.Push(msg)

// 同步发送（TCP）
conn.Send(msg)
```

### 接收消息（集群层）

```go
proxy.AddRouteHandler(1001, func(ctx *client.Context) {
    req := &GreetReq{}
    if err := ctx.Parse(req); err != nil {
        log.Error("解析失败:", err)
        return
    }
    log.Info("收到消息:", req.Message)

    // 回复消息
    ctx.Conn().Push(&cluster.Message{
        Route: 1002,
        Data:  &GreetRes{Message: "world"},
    })
})
```

### 接收消息（网络层）

```go
client.OnReceive(func(conn network.Conn, data []byte) {
    msg, err := packet.UnpackMessage(data)
    if err != nil {
        log.Error("解包失败:", err)
        return
    }
    log.Infof("收到消息: seq=%d route=%d data=%s", msg.Seq, msg.Route, string(msg.Buffer))
})
```

## 事件处理

### 集群事件

```go
// 连接建立
proxy.AddEventListener(cluster.Connect, func(conn *client.Conn) {
    log.Info("连接建立:", conn.ID())
})

// 断线重连
proxy.AddEventListener(cluster.Reconnect, func(conn *client.Conn) {
    log.Info("重新连接:", conn.ID())
})

// 连接断开
proxy.AddEventListener(cluster.Disconnect, func(conn *client.Conn) {
    log.Info("连接断开:", conn.ID())
})
```

### 生命周期钩子

```go
// 初始化钩子
proxy.AddHookListener(cluster.Init, func(p *client.Proxy) {
    log.Info("客户端初始化")
})

// 启动钩子
proxy.AddHookListener(cluster.Start, func(p *client.Proxy) {
    log.Info("客户端启动")
    // 通常在此建立连接
    conn, _ := p.Dial()
    // ...
})

// 销毁钩子（可运行时注册）
proxy.AddHookListener(cluster.Destroy, func(p *client.Proxy) {
    log.Info("客户端销毁")
})
```

### 默认路由处理器

```go
proxy.SetDefaultRouteHandler(func(ctx *client.Context) {
    log.Warn("未处理的路由:", ctx.Route())
    log.Warn("消息数据:", string(ctx.Buffer()))
})
```

## 最佳实践

### 1. 协议选择

| 场景 | 推荐协议 |
| --- | --- |
| 游戏服务器内部通信 | TCP |
| 浏览器/移动端连接 | WebSocket |
| 高性能要求 | TCP |
| 防火墙穿透 | WebSocket |

### 2. 错误处理

```go
conn, err := proxy.Dial(client.WithDialAddr("server:3553"))
if err != nil {
    log.Error("拨号失败:", err)
    // 重试逻辑
    time.Sleep(3 * time.Second)
    conn, err = proxy.Dial(client.WithDialAddr("server:3553"))
    if err != nil {
        log.Fatal("重试失败:", err)
    }
}
```

### 3. 消息序列化

```go
// 使用 JSON 编解码器
c := client.NewClient(
    client.WithCodec(json.NewCodec()),
    client.WithClient(tcp.NewClient()),
)

// 使用 Protobuf 编解码器
c := client.NewClient(
    client.WithCodec(proto.NewCodec()),
    client.WithClient(tcp.NewClient()),
)
```

### 4. 心跳配置

```go
// 禁用心跳（适用于外部负载均衡器管理心跳）
client := tcp.NewClient(
    tcp.WithClientHeartbeatInterval(0),
)

// 自定义心跳间隔
client := tcp.NewClient(
    tcp.WithClientHeartbeatInterval(30 * time.Second),
)
```

### 5. 连接池模式

```go
type ConnPool struct {
    conns []*client.Conn
    mu    sync.RWMutex
}

func (p *ConnPool) Get() *client.Conn {
    p.mu.RLock()
    defer p.mu.RUnlock()
    // 简单轮询
    return p.conns[rand.Intn(len(p.conns))]
}
```

### 6. 断线重连

```go
proxy.AddEventListener(cluster.Disconnect, func(conn *client.Conn) {
    log.Warn("连接断开，尝试重连...")
    go func() {
        time.Sleep(3 * time.Second)
        newConn, err := proxy.Dial(client.WithDialAddr("server:3553"))
        if err != nil {
            log.Error("重连失败:", err)
            return
        }
        log.Info("重连成功:", newConn.ID())
    }()
})
```

## TCP 与 WebSocket 对比

| 特性 | TCP | WebSocket |
| --- | --- | --- |
| 底层库 | Go 标准库 `net` | `gorilla/websocket` |
| 默认地址 | `127.0.0.1:3553` | `ws://127.0.0.1:3553` |
| TLS | `caFile` + `serverName` | `wss://` 协议 |
| 消息读取 | 流式读取 | 帧读取 |
| `Send` 方法 | 同步写入 | 异步高优先级队列 |
| `Push` 方法 | 异步单一队列 | 异步低优先级队列 |
| 写队列 | 单一队列 | 双优先级队列 |
| 适用场景 | 服务器间通信 | 浏览器/移动端 |

## 常见问题

### Q: 如何切换 TCP 和 WebSocket？

只需替换 `client.WithClient()` 参数，业务代码无需改动：

```go
// TCP
c := client.NewClient(client.WithClient(tcp.NewClient()))

// WebSocket
c := client.NewClient(client.WithClient(ws.NewClient()))
```

### Q: 如何处理消息乱序？

使用序列号（Seq）进行消息匹配：

```go
var seq uint32

// 发送时记录序列号
seq++
conn.Push(&cluster.Message{Seq: seq, Route: 1, Data: req})

// 接收时匹配序列号
proxy.AddRouteHandler(1, func(ctx *client.Context) {
    if ctx.Seq() == expectedSeq {
        // 处理响应
    }
})
```

### Q: 如何限制消息频率？

```go
limiter := rate.NewLimiter(rate.Every(time.Second), 10) // 每秒 10 条

proxy.AddRouteHandler(1, func(ctx *client.Context) {
    if !limiter.Allow() {
        log.Warn("消息频率超限")
        return
    }
    // 处理消息
})
```

## 参考资料

- [due GitHub](https://github.com/dobyte/due)
- [due 文档](https://dobyte.github.io/due-docs/)
- [gorilla/websocket](https://github.com/gorilla/websocket)
