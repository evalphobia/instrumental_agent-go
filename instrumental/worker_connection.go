package instrumental

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

// Connection has TCP connection with Instrumental API.
type Connection struct {
	logger    Logger
	conn      net.Conn
	authCount int
	isDebug   bool
	timeout   time.Duration
}

// newConnection creates authentificated TCP connection with Instrumental API.
func newConnection(conf connectionConfig) (*Connection, error) {
	c := &Connection{
		logger:  conf.Logger,
		isDebug: conf.Debug,
		timeout: conf.Timeout,
	}

	conn, err := net.Dial("tcp", conf.GetDSN())
	if err != nil {
		return nil, err
	}

	c.conn = conn
	err = c.auth(conf.GetProjectToken(), conf.GetHostname())
	if err != nil {
		return nil, err
	}
	return c, nil
}

// auth performes authentification with Instrumental API.
func (c *Connection) auth(projectToken, hostName string) error {
	helloMsg := fmt.Sprintf("hello version go/instrumental_agent/%s hostname %s\n", clientVersion, hostName)
	c.debugLog("[Connection] Hello: `%s`", helloMsg)
	c.writeString(helloMsg)
	c.writeString(fmt.Sprintf("authenticate %s\n", projectToken))

	b, err := c.readWithTimeout()
	switch {
	case err != nil:
		return fmt.Errorf("Auth Error: %s", err)
	case strings.HasPrefix(string(b), "ok\nfail"):
		return fmt.Errorf("Auth Failed")
	}

	if !strings.HasPrefix(string(b), "ok\nok") {
		if c.authCount > 2 {
			return fmt.Errorf("Auth Retry Error: %s", string(b))
		}

		c.authCount++
		c.debugLog("retry auth")
		err := c.auth(projectToken, hostName)
		if err != nil {
			return err
		}
	}

	c.debugLog("[Connection] auth success")
	return nil
}

// close closes TCP connetion.
func (c *Connection) close() error {
	if c.conn != nil {
		c.debugLog("[Connection] closed")
		return c.conn.Close()
	}

	c.debugLog("[Connection] close failed: connection is nil")
	return nil
}

// read receive bytes from Instrumental.
func (c *Connection) read() ([]byte, error) {
	if c.conn == nil {
		return nil, errors.New("connection is nil")
	}

	buf := make([]byte, 1024)
	conn := c.conn
	_, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf, nil
}

func (c *Connection) readWithTimeout() ([]byte, error) {
	c.conn.SetReadDeadline(time.Now().Add(c.timeout))
	return c.read()
}

// write sends bytes to Instrumental.
func (c *Connection) write(b []byte) error {
	if c.conn == nil {
		return errors.New("connection is nil")
	}

	c.debugLog("[Connection] Write: %s", string(b))
	_, err := c.conn.Write(b)
	return err
}

func (c *Connection) writeString(s string) error {
	return c.write([]byte(s))
}

func (c *Connection) writeWithTimeout(b []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(c.timeout))
	return c.write(b)
}

func (c *Connection) debugLog(format string, v ...interface{}) {
	if !c.isDebug {
		return
	}
	c.logger.Debugf(format, v...)
}
