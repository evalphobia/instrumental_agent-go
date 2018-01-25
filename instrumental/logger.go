package instrumental

import "log"

// Logger is logging interface.
type Logger interface {
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

// DummyLogger does not ouput anything
type DummyLogger struct{}

// Infof does nothing.
func (*DummyLogger) Infof(format string, v ...interface{}) {}

// Debugf does nothing.
func (*DummyLogger) Debugf(format string, v ...interface{}) {}

// Errorf does nothing.
func (*DummyLogger) Errorf(format string, v ...interface{}) {}

// StdLogger use standard log package.
type StdLogger struct{}

// Infof logging information.
func (*StdLogger) Infof(format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}

// Debugf logging debug information.
func (*StdLogger) Debugf(format string, v ...interface{}) {
	log.Printf("[DEBUG] "+format, v...)
}

// Errorf logging error information.
func (*StdLogger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}
