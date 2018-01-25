package instrumental

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config is a configuration to initialize client
type Config struct {
	Logger
	ProjectToken string
	Timeout      time.Duration
	Disabled     bool
	Debug        bool
	// APIHost and APIPort is used for overriding Instrumental server info.
	APIHost string
	APIPort int
	// Hostname is client's server hostname.
	Hostname string
}

// configuration for connection.
type connectionConfig struct {
	Logger
	ProjectToken string
	APIHost      string
	APIPort      int
	Hostname     string
	Timeout      time.Duration
	Debug        bool
}

func (c connectionConfig) GetProjectToken() string {
	if c.ProjectToken != "" {
		return c.ProjectToken
	}

	return os.Getenv("INSTRUMENTAL_PROJECT_TOKEN")
}

func (c connectionConfig) GetHostname() string {
	if c.Hostname != "" {
		return c.Hostname
	}

	env := os.Getenv("INSTRUMENTAL_HOSTNAME")
	if env != "" {
		return env
	}

	hostname, _ := os.Hostname()
	return hostname
}

// return `hostname:port`
func (c connectionConfig) GetDSN() string {
	return fmt.Sprintf("%s:%d", c.getHost(), c.getPort())
}

func (c connectionConfig) getHost() string {
	if c.APIHost != "" {
		return c.APIHost
	}

	env := os.Getenv("INSTRUMENTAL_API_HOST")
	if env != "" {
		return env
	}
	return defaultAPIHost
}

func (c connectionConfig) getPort() int {
	if c.APIPort > 0 {
		return c.APIPort
	}

	if env, err := strconv.Atoi(os.Getenv("INSTRUMENTAL_API_PORT")); err == nil {
		return env
	}
	return defaultAPIPort
}
