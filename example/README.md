Example of Instrumental Agent for Golang
----

# Quick Usage

```bash
$ INSTRUMENTAL_PROJECT_TOKEN=<Your project token> go run example_send_metric.go


start
2018/01/25 19:18:24 [DEBUG] [Connection] Hello: `hello version go/instrumental_agent/0.1 hostname foobar
`
2018/01/25 19:18:24 [DEBUG] [Connection] Write: hello version go/instrumental_agent/0.1 hostname foobar
2018/01/25 19:18:24 [DEBUG] [Connection] Write: authenticate <Your project token>
2018/01/25 19:18:24 [DEBUG] [Connection] auth success
2018/01/25 19:18:24 [DEBUG] [Worker] Putting message
2018/01/25 19:18:24 [DEBUG] [Worker] Putting message
2018/01/25 19:18:24 [DEBUG] [Worker] Putting message
2018/01/25 19:18:24 [DEBUG] [Worker] Putting message
2018/01/25 19:18:24 [DEBUG] [Worker] Putting message
2018/01/25 19:18:24 [DEBUG] [Worker] Putting message
2018/01/25 19:18:24 [DEBUG] [Worker] Putting message
2018/01/25 19:18:24 [DEBUG] [Worker] Putting message
2018/01/25 19:18:24 [DEBUG] [Worker] Running loop
2018/01/25 19:18:24 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:24 [DEBUG] [Connection] Write: gauge user.count 100.000000 1516875504 1
2018/01/25 19:18:24 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:24 [DEBUG] [Connection] Write: increment user.count 1.000000 1516875504 1
2018/01/25 19:18:24 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:24 [DEBUG] [Connection] Write: increment user.count 1.000000 1516875504 1
2018/01/25 19:18:24 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:24 [DEBUG] [Connection] Write: increment user.count 1.000000 1516875504 1
2018/01/25 19:18:24 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:24 [DEBUG] [Connection] Write: notice 1516875504 0 Released new version
2018/01/25 19:18:24 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:24 [DEBUG] [Connection] Write: increment user.count 1.000000 1516875504 1
2018/01/25 19:18:24 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:24 [DEBUG] [Connection] Write: increment user.count 1.000000 1516875504 1
2018/01/25 19:18:24 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:24 [DEBUG] [Connection] Write: increment user.count 1.000000 1516875504 1
2018/01/25 19:18:27 [DEBUG] [Worker] Putting message
2018/01/25 19:18:27 [DEBUG] [Worker] Putting message
2018/01/25 19:18:27 [DEBUG] [Worker] Putting message
2018/01/25 19:18:27 [DEBUG] [Worker] Putting message
2018/01/25 19:18:27 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:27 [DEBUG] [Connection] Write: gauge user.count 200.000000 1516875507 1
2018/01/25 19:18:27 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:27 [DEBUG] [Connection] Write: increment user.count 1.000000 1516875507 1
2018/01/25 19:18:27 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:27 [DEBUG] [Connection] Write: increment user.count 1.000000 1516875507 1
2018/01/25 19:18:27 [DEBUG] [Worker] sendPacket
2018/01/25 19:18:27 [DEBUG] [Connection] Write: notice 1516875507 0 End example
```