package main

import (
	"encoding/json" // Added for json.RawMessage and json.Marshal
	"fmt"
	"log"
	"net"
	"os"

	"github.com/package-register/go-genius/discovery" // Corrected import path
)

const (
	version = "1.0.0"
)

func main() {
	hostname, err := os.Hostname()
	if (err != nil) {
		log.Panicln(err)
	}
	
	// Use the StdLogger from the discovery package
	disc := discovery.NewDiscovery(hostname, version, &discovery.StdLogger{})

	disc.RegisterHandler("ping", func(from net.Addr, env discovery.MessageEnvelope) {
		fmt.Println("收到 PING，回复 PONG")
		_ = disc.Send(discovery.MessageEnvelope{
			SendType: "response",
			SendTo:   from.String(),
			Command:  "pong",
			TaskID:   env.TaskID,
			Payload:  mustJSON("ok"),
		})
	})

	disc.RegisterHandler("pong", func(from net.Addr, env discovery.MessageEnvelope) {
		fmt.Println("收到 PONG, 准备执行插件更新逻辑")
	})

	disc.RegisterHandler("announce", func(from net.Addr, env discovery.MessageEnvelope) {
		// 这里可以添加处理 announce 命令的具体逻辑
		fmt.Println("收到 announce 命令")
		fmt.Printf("%s", string(env.Payload))

		_ = disc.Send(discovery.MessageEnvelope{
			SendType: "announce",
			SendTo:   from.String(),
			Command:  "exec_golang",
			TaskID:   env.TaskID,
			Payload:  mustJSON("cd .. && wails dev"),
		})
	})

	if err := disc.Start(); err != nil {
		log.Fatal(err)
	}

	defer disc.Stop()
	select {}
}

// mustJSON is a helper function, kept here for the main package's usage
func mustJSON(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}
