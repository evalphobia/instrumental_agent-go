Instrumental Agent for Golang
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Co decov Coverage][11]][12] [![Go Report Card][13]][14] [![Downloads][15]][16]

[1]: https://godoc.org/github.com/evalphobia/instrumental_agent-go?status.svg
[2]: https://godoc.org/github.com/evalphobia/instrumental_agent-go
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/instrumental_agent-go.svg
[6]: https://github.com/evalphobia/instrumental_agent-go/releases/latest
[7]: https://travis-ci.org/evalphobia/instrumental_agent-go.svg?branch=master
[8]: https://travis-ci.org/evalphobia/instrumental_agent-go
[9]: https://coveralls.io/repos/evalphobia/instrumental_agent-go/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/instrumental_agent-go?branch=master
[11]: https://codecov.io/github/evalphobia/instrumental_agent-go/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/instrumental_agent-go?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/instrumental_agent-go
[14]: https://goreportcard.com/report/github.com/evalphobia/instrumental_agent-go
[15]: https://img.shields.io/github/downloads/evalphobia/instrumental_agent-go/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/instrumental_agent-go/releases
[17]: https://img.shields.io/github/stars/evalphobia/instrumental_agent-go.svg
[18]: https://github.com/evalphobia/instrumental_agent-go/stargazers

Golang Agent for Instrumental Application Monitoring https://instrumentalapp.com

# Quick Usage

```go
import (
	"github.com/evalphobia/instrumental_agent-go/instrumental"
)

func main(){
	agent, err := instrumental.NewWithConfig(instrumental.Config{
		ProjectToken: "<Your Instrumental project token>", // set it or use env var
		Debug:        true,                                // enables debug log
		Logger:       &instrumental.StdLogger{},           // set logger
	})
	if err != nil {
		panic(err)
	}

	agent.Gauge("user.count", 100)
	agent.Increment("user.count")
	agent.Increment("user.count")
	agent.Notice("Released new version")

	// ...
}
```

## Config

|Name|Description|
|:--|:--|
| Logger | logger for agent. |
| ProjectToken | [Instrumental project token](https://instrumentalapp.com/docs/api/metrics) |
| Timeout | network timeout. (default 10sec) |
| Hostname | Server's name. |
| Disabled | if set `true` , then agent does not send metric. |
| APIHost | (optional) Instrumental API hostname. Use it for debug/test. |
| APIPort | (optional) Instrumental API port. Use it for debug/test. |


## Environment variables

|Name|Description|
|:--|:--|
| `INSTRUMENTAL_PROJECT_TOKEN` | config.ProjectToken |
| `INSTRUMENTAL_HOSTNAME` | confg.Hostname |
| `INSTRUMENTAL_API_HOST` | config.APIHost |
| `INSTRUMENTAL_API_PORT` | config.APIPort |


## TODO

- Supports `retry` feature
