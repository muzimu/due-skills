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
	container := due.NewContainer()

	server := ws.NewServer(
		ws.WithServerAddr(":8800"),
		ws.WithServerMaxConnNum(10000),
		ws.WithServerAuthorizeTimeout(30*time.Second), // 授权超时：30秒内未Bind则断开
	)

	locator := redis.NewLocator(
		redis.WithAddrs("127.0.0.1:6379"),
	)

	registry := consul.NewRegistry(
		consul.WithAddr("127.0.0.1:8500"),
	)

	component := gate.NewGate(
		gate.WithID("gate-chat-001"),
		gate.WithName("gate"),
		gate.WithServer(server),
		gate.WithLocator(locator),
		gate.WithRegistry(registry),
	)

	container.Add(component)

	log.Info("Chat Room Gate 启动中，WebSocket 端口: 8800")
	container.Serve()
}
