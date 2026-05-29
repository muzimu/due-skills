---
name: due-skills
description: |
  Comprehensive knowledge base for due game server framework (v2.5.7).

  **Use this skill when:**
  - Working with due framework (any version, especially v2.5.7)
  - Building distributed game servers or real-time applications
  - Implementing Gate services (TCP/KCP/WebSocket) for client connections
  - Creating Node services with Actor model for stateful game logic
  - Setting up Mesh microservices for stateless business logic
  - **Creating HTTP/REST API services** with due's HTTP component
  - **Building web servers** with fiber-based routing and middleware
  - **Implementing Swagger API documentation** for HTTP endpoints
  - **Building TCP/WebSocket clients** for connecting to game servers
  - **Implementing client applications** with message routing and event handling
  - Configuring service discovery (Consul/Etcd/Nacos)
  - Implementing event buses (Redis/NATS/Kafka/RabbitMQ)
  - Adding caching layers (Redis/Memcache)
  - Handling message routing, serialization, or RPC communication
  - Using Container pattern to manage component lifecycle
  - Implementing player session management and binding

  **Always consult this skill for due-related tasks** - it contains v2.5.7 specific API changes, correct module paths (github.com/dobyte/due/v2), and production-ready patterns that prevent common mistakes like using wrong package imports or outdated APIs.

  **Features:**
  - Complete architecture guides with Gate → Node → Mesh patterns
  - Actor model implementation examples with Router handlers
  - Multi-protocol support (TCP/KCP/WebSocket)
  - **HTTP/REST API development** with fiber-based routing and middleware
  - **Swagger API documentation** integration
  - **TCP/WebSocket client development** with message routing and event handling
  - Production best practices for game servers
  - Common pitfall solutions (wrong imports, deprecated APIs)
  - Updated for due v2.5.7 API changes (enhanced RPC config, Session disconnect support, HTTP multi-style handlers)
license: Apache-2.0
allowed-tools:
  - Read
  - Grep
  - Glob
  - Edit
  - Write
  - Bash
  - task
  - todowrite
  - lsp_diagnostics
  - lsp_goto_definition
  - lsp_find_references
  - lsp_symbols
  - lsp_prepare_rename
  - lsp_rename
  - ast_grep_search
  - ast_grep_replace
  - glob
  - websearch
  - webfetch
  - codesearch
  - context7_resolve-library-id
  - context7_query-docs
trigger-keywords:
  - "due.NewContainer"
  - "gate.NewGate"
  - "node.NewNode"
  - "mesh.NewMesh"
  - "proxy.Router"
  - "AddRouteHandler"
  - "github.com/dobyte/due"
  - "due-skills"
  - "websocket"
  - "TCP"
  - "KCP"
  - "Actor"
  - "Consul"
  - "Etcd"
  - "Nacos"
  - "Redis"
  - "EventBus"
  - "rpcx"
  - "grpc"
  - "http.NewServer"
  - "http.Server"
  - "http.Context"
  - "http.Handler"
  - "fiber"
  - "REST API"
  - "Swagger"
  - "tcp.NewClient"
  - "ws.NewClient"
  - "client.NewClient"
  - "network.Client"
  - "network.Conn"
  - "cluster/client"
file-patterns:
  - "*.go"
directories:
  - "gate"
  - "node"
  - "mesh"
  - "cluster"
  - "cmd"
  - "internal"
---

# due Skills for AI Agents

This skill provides comprehensive due game server framework knowledge (v2.5.7), optimized for AI agents helping developers build production-ready distributed game servers. due is a lightweight, high-performance distributed game server framework (Apache 2.0 license), featuring standardized development patterns and proven deployment in enterprise game projects.

**v2.5.7 Key Changes:**
- Enhanced RPC configuration options for Gate/Node/Mesh (connNum, callTimeout, dialTimeout, dialRetryTimes, writeTimeout, writeQueueSize, faultRecoveryTime)
- Session API: Push/Multicast/Broadcast/Publish methods now support disconnect parameter for auto-closing connections after push
- Session performance: Concurrent message pushing using errgroup for Multicast/Broadcast/Publish
- HTTP Router: Added All method and support for multiple handler styles (due/fiber/express/net/http/fasthttp)
- WebSocket: Replaced handshakeTimeout with writeTimeout, added writeQueueSize
- **Breaking**: `WithAddr` renamed to `WithAddrs` for redis/etcd/kafka/memcache components (supports multiple addresses)
- **Breaking**: NATS EventBus uses `WithUrl` instead of `WithAddr`
- **Breaking**: Nacos uses `WithUrls` instead of `WithAddr`
- Various bug fixes and performance optimizations

## 🎯 When to Use This Skill

Invoke this skill when working with due:
- **Creating game servers**: Gate services, Node services, or Mesh microservices
- **Protocol implementation**: TCP, KCP, or WebSocket client connections
- **Actor model**: Implementing stateful game logic with due Actor system
- **HTTP/REST API development**: Building web servers, REST APIs, or HTTP microservices
- **Swagger documentation**: Generating and serving API documentation
- **Client development**: Building TCP/WebSocket clients for connecting to game servers
- **Service discovery**: Consul, Etcd, or Nacos integration
- **Event-driven architecture**: Redis, NATS, Kafka, or RabbitMQ event buses
- **Caching strategies**: Redis or Memcache integration
- **Message routing**: Custom route handling and serialization

## 📚 Knowledge Structure

This skill organizes due knowledge into focused modules. **Load specific guides as needed** rather than reading everything at once:

### Quick Start Guide
**File**: [getting-started/README.md](getting-started/README.md)
**When to load**: Starting a new due project, environment setup
**Contains**: Installation, first server, docker-compose setup, basic concepts

### Pattern Guides (Detailed Reference)

#### 1. Architecture Patterns
**File**: [references/architecture-patterns.md](references/architecture-patterns.md)
**When to load**: Understanding due architecture, designing server topology
**Contains**:
- Gate → Node → Mesh three-tier architecture
- Stateful vs Stateless service design
- Service discovery and registration
- Inter-service communication (gRPC/RPCX)

#### 2. Gate Development Patterns
**File**: [references/gate-patterns.md](references/gate-patterns.md)
**When to load**: Building gateway services for client connections
**Contains**:
- TCP protocol implementation
- KCP protocol implementation
- WebSocket protocol implementation
- Connection management and session handling
- Message packet format: `size + header + route + seq + message`
- Heartbeat mechanism

#### 3. Node Development Patterns
**File**: [references/node-patterns.md](references/node-patterns.md)
**When to load**: Implementing core game logic with Actor model
**Contains**:
- Actor model fundamentals in due
- Creating and managing Actors
- Message passing between Actors
- Stateful game logic implementation
- Actor lifecycle management

#### 4. Mesh Development Patterns
**File**: [references/mesh-patterns.md](references/mesh-patterns.md)
**When to load**: Building stateless microservices
**Contains**:
- Stateless service design
- Microservice communication patterns
- Load balancing strategies
- Service mesh integration

#### 5. Component Usage Patterns
**File**: [references/component-patterns.md](references/component-patterns.md)
**When to load**: Using due components (logging, caching, eventbus, etc.)
**Contains**:
- **Logging**: Console, File, Aliyun, Tencent log drivers
- **Cache**: Redis and Memcache integration
- **EventBus**: Redis, NATS, Kafka, RabbitMQ backends
- **Registry**: Consul, Etcd, Nacos service registration
- **Config**: Consul, Etcd, Nacos config centers with JSON/YAML/TOML/XML
- **Lock**: Distributed locks with Redis/Memcache
- **Crypto**: RSA and ECC encryption
- **Transport**: gRPC and RPCX communication

#### 6. HTTP Service Patterns
**File**: [references/http-patterns.md](references/http-patterns.md)
**When to load**: Building HTTP/REST APIs, web servers, or HTTP microservices
**Contains**:
- HTTP server creation and configuration
- Routing system with multiple handler styles (due/fiber/net/http/fasthttp)
- Middleware implementation (global and route-level)
- Request binding and response formatting
- Swagger API documentation integration
- CORS and TLS configuration
- Microservice integration via Proxy

#### 7. Client Development Patterns
**File**: [references/client-patterns.md](references/client-patterns.md)
**When to load**: Building TCP/WebSocket clients, connecting to game servers
**Contains**:
- TCP client creation and configuration
- WebSocket client creation and configuration
- Cluster-level client with codec and encryption
- Connection management and lifecycle
- Message sending and receiving
- Event handling and route dispatching
- Heartbeat mechanism and reconnection

#### 8. Message Protocol
**File**: [references/protocol-patterns.md](references/protocol-patterns.md)
**When to load**: Defining custom message formats, serialization
**Contains**:
- Default packet format: `size + header + route + seq + message`
- Heartbeat packet: `size + header + extcode + heartbeat_time`
- Custom route and sequence number configuration
- Message serialization patterns
- Request/Response correlation with sequence numbers

### Supporting Resources

#### Best Practices
**File**: [best-practices/overview.md](best-practices/overview.md)
**When to load**: Production deployment, performance optimization
**Contains**: Configuration management, logging strategies, monitoring, scaling

#### Troubleshooting
**File**: [troubleshooting/common-issues.md](troubleshooting/common-issues.md)
**When to load**: Debugging errors, connection issues, runtime problems
**Contains**: Common error messages, solutions, configuration pitfalls

#### Project Structure
**File**: [references/project-structure.md](references/project-structure.md)
**When to load**: Understanding due directory layout, creating new projects
**Contains**:
```
due/
├── .docker/          # Docker configurations
├── cluster/          # Cluster management
├── component/        # Reusable components
├── config/           # Configuration management
├── core/             # Core framework logic
├── encoding/         # Message encoding/decoding
├── eventbus/         # Event bus implementations
├── log/              # Logging components
├── network/          # Network protocols
├── registry/         # Service registration
├── session/          # Session management
└── transport/        # RPC transport
```

## 🚀 Common Workflows

| Task | Guide |
|------|-------|
| Gate Service | [gate-patterns.md](references/gate-patterns.md) |
| Actor Game Logic | [node-patterns.md](references/node-patterns.md) |
| Event-Driven Communication | [component-patterns.md](references/component-patterns.md) |
| Service Discovery | [architecture-patterns.md](references/architecture-patterns.md) |
| HTTP/REST API | [http-patterns.md](references/http-patterns.md) |
| TCP/WebSocket Client | [client-patterns.md](references/client-patterns.md) |

## ⚡ Key Principles

When generating or reviewing due code, always apply these principles:

### ✅ Always Follow

- **Three-tier separation**: Keep Gate (connection) → Node (logic) → Mesh (stateless) distinct
- **Actor model**: Use Actors for stateful game logic, not global state
- **Message routing**: Define clear route patterns for all messages
- **Configuration**: Externalize config with config center support
- **Service registration**: Always register services with discovery system
- **Error handling**: Use proper error wrapping and logging
- **Session management**: Store client state in Session, not global variables
- **HTTP best practices**: Use `ctx.Success()`/`ctx.Failure()` for consistent responses
- **Request validation**: Always validate and bind request data before processing
- **Middleware order**: Place global middleware (Logger, Recover, CORS) before custom ones
- **Client protocol selection**: Use TCP for server-to-server, WebSocket for browser/mobile
- **Connection lifecycle**: Always handle disconnect events and implement reconnection

### ❌ Never Do

- Put game logic directly in Gate layer (violates separation)
- Use global mutable state for game data (use Actors instead)
- Hard-code service addresses (use service discovery)
- Skip message validation in handlers
- Block Actor message loops with long operations
- Forget to handle connection close events
- Return raw errors from HTTP handlers (use `ctx.Failure()` instead)
- Skip CORS configuration for public APIs
- Expose internal error details in production responses
- Use `Send` for concurrent WebSocket writes (use `Push` instead)
- Ignore heartbeat timeout in client connections

## 📖 Learning Path

- **New to due?** → [getting-started/README.md](getting-started/README.md)
- **Building game servers?** → [architecture-patterns.md](references/architecture-patterns.md) → [node-patterns.md](references/node-patterns.md) → [gate-patterns.md](references/gate-patterns.md)
- **HTTP/REST APIs?** → [http-patterns.md](references/http-patterns.md)
- **Production deployment?** → [best-practices/overview.md](best-practices/overview.md) → [troubleshooting/common-issues.md](troubleshooting/common-issues.md)

## 🔗 Related Resources

- **Official docs**: [https://github.com/dobyte/due](https://github.com/dobyte/due)
- **fiber documentation**: [https://docs.gofiber.io/](https://docs.gofiber.io/)

## 📝 Version Compatibility

- **Target version**: due v2.5.7 (Apache 2.0 licensed)
- **Go version**: Go 1.25.0 or later recommended
- **Module path**: github.com/dobyte/due/v2
- **Dependencies**: grpc, rpcx, redis, nats, kafka, rabbitmq drivers as needed

## 🚀 Quick Start (v2.5.7)

```bash
# Get due v2.5.7
go get -u github.com/dobyte/due/v2@latest

# Get required components
go get -u github.com/dobyte/due/locate/redis/v2@latest
go get -u github.com/dobyte/due/network/ws/v2@latest
go get -u github.com/dobyte/due/registry/consul/v2@latest
go get -u github.com/dobyte/due/transport/rpcx/v2@latest
go get -u github.com/dobyte/due/component/http/v2@latest
```

**Gate Server Example (v2.5.7):**
```go
package main

import (
   "github.com/dobyte/due/locate/redis/v2"
   "github.com/dobyte/due/network/ws/v2"
   "github.com/dobyte/due/registry/consul/v2"
   "github.com/dobyte/due/v2"
   "github.com/dobyte/due/v2/cluster/gate"
)

func main() {
   container := due.NewContainer()
   server := ws.NewServer()
   locator := redis.NewLocator()
   registry := consul.NewRegistry()
   component := gate.NewGate(
      gate.WithServer(server),
      gate.WithLocator(locator),
      gate.WithRegistry(registry),
      // v2.5.7: 新增 RPC 配置选项
      gate.WithCallTimeout(5*time.Second),
      gate.WithDialTimeout(3*time.Second),
      gate.WithWriteTimeout(0),
      gate.WithWriteQueueSize(4096),
   )
   container.Add(component)
   container.Serve()
}
```

**Node Server Example (v2.5.7):**
```go
package main

import (
   "github.com/dobyte/due/locate/redis/v2"
   "github.com/dobyte/due/registry/consul/v2"
   "github.com/dobyte/due/v2"
   "github.com/dobyte/due/v2/cluster/node"
)

func main() {
   container := due.NewContainer()
   locator := redis.NewLocator()
   registry := consul.NewRegistry()
   component := node.NewNode(
      node.WithLocator(locator),
      node.WithRegistry(registry),
      // v2.5.7: 新增 RPC 配置选项
      node.WithConnNum(10),
      node.WithCallTimeout(5*time.Second),
   )
   initListen(component.Proxy())
   container.Add(component)
   container.Serve()
}

func initListen(proxy *node.Proxy) {
   proxy.Router().AddRouteHandler(routeID, isSync, handlerFunc)
}
```

**HTTP Server Example (v2.5.7):**
```go
package main

import (
   "github.com/dobyte/due/component/http/v2"
   "github.com/dobyte/due/v2"
   "github.com/dobyte/due/v2/codes"
)

func main() {
   container := due.NewContainer()
   component := http.NewServer(
      http.WithName("api-server"),
      http.WithAddr(":8080"),
      http.WithConsole(true),
   )
   initApp(component.Proxy())
   container.Add(component)
   container.Serve()
}

func initApp(proxy *http.Proxy) {
   router := proxy.Router()
   router.Get("/api/v1/greet", func(ctx http.Context) error {
      return ctx.Success("Hello")
   })
   router.Post("/api/v1/users", func(ctx http.Context) error {
      // 处理创建用户
      return ctx.Failure(codes.InvalidArgument)
   })
}
```

**Client Example (v2.5.7):**
```go
package main

import (
   "github.com/dobyte/due/cluster/client/v2"
   "github.com/dobyte/due/cluster/v2"
   "github.com/dobyte/due/encoding/json"
   "github.com/dobyte/due/network/tcp/v2"  // 或 "github.com/dobyte/due/network/ws/v2"
   "github.com/dobyte/due/v2"
   "github.com/dobyte/due/v2/log"
)

func main() {
   container := due.NewContainer()
   c := client.NewClient(
      client.WithName("my-client"),
      client.WithCodec(json.NewCodec()),
      client.WithClient(tcp.NewClient()),  // 或 ws.NewClient()
   )
   initApp(c.Proxy())
   container.Add(c)
   container.Serve()
}

func initApp(proxy *client.Proxy) {
   proxy.AddHookListener(cluster.Start, func(p *client.Proxy) {
      conn, _ := p.Dial(client.WithDialAddr("127.0.0.1:3553"))
      conn.Push(&cluster.Message{Route: 1, Data: "hello"})
   })
   proxy.AddRouteHandler(1, func(ctx *client.Context) {
      var res string
      ctx.Parse(&res)
      log.Info("收到响应:", res)
   })
}
```

---

**Quick invocation**: Use `/due-skills` or ask "How do I [task] with due?"
**Need help?** Reference the specific pattern guide for detailed examples.
