package instrumental

import (
	"time"
)

var (
	defaultLogger = &DummyLogger{}
)

// Agent is a client for Instrumental API.
type Agent struct {
	*Worker
	Config

	Timeout  time.Duration
	Disabled bool
}

// New creates a Instrumental API client agent.
func New() (Agent, error) {
	return newWithConfig(Config{
		Timeout: defaultTimeout,
		Logger:  defaultLogger,
	})
}

// NewWithConfig creates a Instrumental API client agent with configuration
func NewWithConfig(config Config) (Agent, error) {
	if config.Timeout == 0 {
		config.Timeout = defaultTimeout
	}
	if config.Logger == nil {
		config.Logger = defaultLogger
	}

	return newWithConfig(config)
}

func newWithConfig(config Config) (Agent, error) {
	a := Agent{
		Config:   config,
		Timeout:  config.Timeout,
		Disabled: config.Disabled,
	}
	if a.Disabled {
		return a, nil
	}

	w, err := newWorker(connectionConfig{
		ProjectToken: config.ProjectToken,
		APIHost:      config.APIHost,
		APIPort:      config.APIPort,
		Hostname:     config.Hostname,
		Debug:        config.Debug,
		Logger:       config.Logger,
		Timeout:      config.Timeout,
	})
	if err != nil {
		return a, err
	}

	a.Worker = w
	return a, nil
}

// Stop stops background worker.
func (a *Agent) Stop() {
	a.Worker.Stop()
}

// Increment sends data for incrementing metric.
// See: https://instrumentalapp.com/docs/metrics#metric-types
func (a *Agent) Increment(metricName string, options ...interface{}) {
	if a.Disabled {
		return
	}

	p := newPacket(commandIncrement)
	p.Name = metricName

	optSize := len(options)
	if optSize >= 1 {
		if v, ok := getFloatNumber(options[0]); ok {
			p.Value = v
		}
	}
	if optSize >= 2 {
		if v, ok := options[1].(time.Time); ok {
			p.Time = v
		}
	}
	if optSize >= 3 {
		if v, ok := options[2].(int); ok {
			p.Count = v
		}
	}

	a.call(p)
}

// Gauge sends numeric metric data.
// See: https://instrumentalapp.com/docs/metrics#metric-types
func (a *Agent) Gauge(metricName string, value float64, options ...interface{}) {
	if a.Disabled {
		return
	}

	p := newPacket(commandGauge)
	p.Name = metricName
	p.Value = value

	optSize := len(options)
	if optSize >= 1 {
		if v, ok := options[0].(time.Time); ok {
			p.Time = v
		}
	}
	if optSize >= 2 {
		if v, ok := options[1].(int); ok {
			p.Count = v
		}
	}

	a.call(p)
}

// Notice sends a notice.
// See: https://instrumentalapp.com/docs/notices
func (a *Agent) Notice(description string, options ...interface{}) {
	if a.Disabled {
		return
	}

	p := newPacket(commandNotice)
	p.Description = description

	optSize := len(options)
	if optSize >= 1 {
		if v, ok := options[0].(time.Time); ok {
			p.Time = v
		}
	}
	if optSize >= 2 {
		if v, ok := options[1].(time.Duration); ok {
			p.Duration = v
		}
	}

	a.call(p)
}

func (a *Agent) call(p packet) {
	a.Worker.Put(p)
}

// getFloatNumber converts from numeric value to float64 value.
func getFloatNumber(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case int:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case uint:
		return float64(v), true
	case uint8:
		return float64(v), true
	case uint16:
		return float64(v), true
	case uint32:
		return float64(v), true
	case uint64:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return v, true
	}
	return 0, false
}
