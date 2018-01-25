package instrumental

import "time"

const (
	clientVersion = "0.1"

	defaultTimeout    = time.Second * 10
	defaultBufferSize = 4096
	defaultAPIHost    = "collector.instrumentalapp.com"
	defaultAPIPort    = 8000

	commandIncrement = "increment"
	commandGauge     = "gauge"
	commandNotice    = "notice"
)
