package instrumental

import (
	"fmt"
	"time"
)

// MetricData is used for sending metric to Instrumental API.
type MetricData struct {
	Name  string
	Value float64
	Time  time.Time

	// Count is used by Increment,Gauge
	Count int

	// Description and Duration is used by Notice
	Description string
	Duration    time.Duration
}

// NewMetricData creates initialized MetricData.
func NewMetricData() MetricData {
	return MetricData{}.FillDefault()
}

// FillDefault fills default value.
func (d MetricData) FillDefault() MetricData {
	if d.Time.IsZero() {
		d.Time = time.Now()
	}
	return d
}

// packet is used for sending packet to Instrumental API.
type packet struct {
	Command     string
	Name        string
	Value       float64
	Time        time.Time
	Count       int
	Description string
	Duration    time.Duration
}

// newPacket creates initialized packet.
func newPacket(command string) packet {
	return packet{
		Command:  command,
		Value:    1.0,
		Time:     time.Now(),
		Count:    1,
		Duration: 0,
	}
}

// getBytes gets byte data for sending metric.
func (p packet) getBytes() []byte {
	switch p.Command {
	case commandIncrement,
		commandGauge:
		return []byte(fmt.Sprintf("%s %s %f %d %d\n", p.Command, p.Name, p.Value, p.Time.Unix(), p.Count))
	case commandNotice:
		return []byte(fmt.Sprintf("%s %d %d %s\n", p.Command, p.Time.Unix(), int(p.Duration.Seconds()), p.Description))
	default:
		return nil
	}
}
