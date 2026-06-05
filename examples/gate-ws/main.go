package main

import (
	"time"

	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/gate"
	"github.com/dobyte/due/v2/log"
)

func main() {
	// 创建容器
	container := due.NewContainer()

	// 创建 WebSocket 服务器
	server := ws.NewServer(
		ws.WithServerAddr(":8800"),
		ws.WithServerMaxConnNum(10000),
		ws.WithServerAuthorizeTimeout(30*time.Second), // 授权超时：30秒内未Bind则断开，0s表示不检测
	)

	// 创建定位器
	locator := redis.NewLocator(
		redis.WithAddrs("127.0.0.1:6379"),
	)

	// 创建注册中心
	registry := consul.NewRegistry(
		consul.WithAddr("127.0.0.1:8500"),
	)

	// 创建 Gate 组件
	component := gate.NewGate(
		gate.WithID("gate-ws-001"),
		gate.WithName("gate"),
		gate.WithServer(server),
		gate.WithLocator(locator),
		gate.WithRegistry(registry),
	)

	// 添加组件到容器
	container.Add(component)

	log.Info("Gate WebSocket 服务启动中，端口: 8800")

	// 启动服务
	container.Serve()
}
