package main

import (
	"fmt"
	"time"

	"github.com/evalphobia/instrumental_agent-go/instrumental"
)

func main() {
	fmt.Println("start")
	agent, err := instrumental.NewWithConfig(instrumental.Config{
		// ProjectToken: "<Your Instrumental project token>", // set it or use env var
		Debug:  true,                      // enables debug log
		Logger: &instrumental.StdLogger{}, // set logger
	})
	if err != nil {
		panic(err)
	}

	agent.Gauge("user.count", 100)
	agent.Increment("user.count")
	agent.Increment("user.count")
	agent.Increment("user.count")
	agent.Notice("Released new version")
	agent.Increment("user.count")
	agent.Increment("user.count")
	agent.Increment("user.count")
	time.Sleep(3 * time.Second)

	agent.Gauge("user.count", 200)
	agent.Increment("user.count")
	agent.Increment("user.count")
	agent.Notice("End example")
	time.Sleep(5 * time.Second)
}
