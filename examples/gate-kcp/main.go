package main

import (
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/network/kcp/v2"
	"github.com/dobyte/due/registry/consul/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/gate"
	"github.com/dobyte/due/v2/log"
)

func main() {
	// 创建容器
	container := due.NewContainer()

	// 创建 KCP 服务器
	server := kcp.NewServer(
		kcp.WithServerListenAddr(":10000"),
		kcp.WithServerMaxConnNum(5000),
		kcp.WithServerMtu(1400),
		kcp.WithServerNoDelay(1, 10, 2, 1),
		kcp.WithServerAckNoDelay(true),
		kcp.WithServerWindowSize(128, 512),
		kcp.WithServerReadBuffer(4194304),
		kcp.WithServerWriteBuffer(4194304),
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
		gate.WithID("gate-kcp-001"),
		gate.WithName("gate"),
		gate.WithServer(server),
		gate.WithLocator(locator),
		gate.WithRegistry(registry),
	)

	// 添加组件到容器
	container.Add(component)

	log.Info("Gate KCP 服务启动中，端口: 10000")

	// 启动服务
	container.Serve()
}
