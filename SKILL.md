---
name: due-skills
description: |
  Comprehensive knowledge base for due game server framework (v2.5.8).

  **Use this skill when:**
  - Working with due framework (any version, especially v2.5.8)
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

  **Always consult this skill for due-related tasks** - it contains v2.5.8 specific API changes, correct module paths (github.com/dobyte/due/v2), and production-ready patterns that prevent common mistakes like using wrong package imports or outdated APIs.

  **Features:**
  - Complete architecture guides with Gate → Node → Mesh patterns
  - Actor model implementation examples with Router handlers
  - Multi-protocol support (TCP/KCP/WebSocket)
  - **HTTP/REST API development** with fiber-based routing and middleware
  - **Swagger API documentation** integration
  - **TCP/WebSocket client development** with message routing and event handling
  - Production best practices for game servers
  - Common pitfall solutions (wrong imports, deprecated APIs)
  - Built for due v2.5.8 APIs (etcd username/password auth for registry & config-center, fine-grained RPC config, Session disconnect support, HTTP multi-style handlers)
license: Apache-2.0
allowed-tools:
  - Read
  - Grep
  - Glob
  - Edit
  - Write
  - Bash
  - TodoWrite
  - lsp_diagnostics
  - lsp_goto_definition
  - lsp_find_references
  - lsp_symbols
  - lsp_prepare_rename
  - lsp_rename
  - ast_grep_search
  - ast_grep_replace
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
  - "dobyte/due"
  - "http.NewServer"
  - "http.Context"
  - "tcp.NewClient"
  - "ws.NewClient"
  - "client.NewClient"
  - "cluster/client"
  - "etcd.WithUsername"
  - "etcd.WithPassword"
  - "WithRefreshInterval"
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

> **IRON LAW**: 始终使用正确的模块路径 `github.com/dobyte/due/v2`，不得使用旧版路径。生成完整项目前，先确认服务拓扑（Gate/Node/Mesh 哪些）、协议（TCP/KCP/WS）、服务发现方案（Consul/Etcd/Nacos），再动手写代码。如不确定某 API 是否存在，先查阅对应 pattern 文件，不得凭记忆猜测。

This skill provides comprehensive due game server framework knowledge (v2.5.8) for building production-ready distributed game servers. due is a lightweight, high-performance distributed game server framework (Apache 2.0).

**v2.5.8 Core Capabilities:**
- Fine-grained RPC configuration for Gate/Node/Mesh (connNum, callTimeout, dialTimeout, dialRetryTimes, writeTimeout, writeQueueSize, faultRecoveryTime)
- Session API: Push/Multicast/Broadcast/Publish support a disconnect parameter for auto-closing connections after push; concurrent batch pushing via errgroup
- HTTP Router: All method plus multiple handler styles (due/fiber/express/net/http/fasthttp)
- Multi-address components use `WithAddrs` (redis/etcd/kafka/memcache); NATS uses `WithUrl`; Nacos uses `WithUrls`
- Etcd registry and config-center support `WithUsername`/`WithPassword` (and `username`/`password` in etc.yaml) for authenticated clusters

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

- **Target version**: due v2.5.8 (Apache 2.0 licensed)
- **Go version**: Go 1.25.0 or later recommended
- **Module path**: github.com/dobyte/due/v2
- **Dependencies**: grpc, rpcx, redis, nats, kafka, rabbitmq drivers as needed

## 🚀 Quick Start

```bash
go get -u github.com/dobyte/due/v2@latest
```

→ 完整安装和示例见 [getting-started/README.md](getting-started/README.md)
