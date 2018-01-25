package instrumental

import (
	"sync"
)

// Worker sends message to Instrumental as a background goroutine.
type Worker struct {
	*Connection

	bufferMu   sync.Mutex
	buffer     chan packet
	wg         sync.WaitGroup
	stopSignal chan struct{}
	isRunning  bool
}

// newWorker creates Worker with connection.
func newWorker(conf connectionConfig) (*Worker, error) {
	conn, err := newConnection(conf)
	if err != nil {
		conf.Logger.Errorf("[Worker] error on `newConnection`: %s", err.Error())
		return nil, err
	}

	w := &Worker{
		Connection: conn,
		buffer:     make(chan packet, defaultBufferSize),
	}
	go w.Run()
	return w, nil
}

// Stop stops worker.
func (w *Worker) Stop() {
	w.Connection.debugLog("[Worker] stopping")
	w.stopSignal <- struct{}{}
}

// Put puts a packet into buffer (then send it to Instrumental API as a metric data).
func (w *Worker) Put(p packet) {
	w.Connection.debugLog("[Worker] Putting message")
	w.wg.Add(1)
	w.buffer <- p
}

// Flush flushes all the metric data in the buffer.
func (w *Worker) Flush() {
	w.bufferMu.Lock()
	defer w.bufferMu.Unlock()
	w.wg.Wait()
}

// Run runs Worker and process metric data as a background worker.
func (w *Worker) Run() {
	if w.isRunning {
		return
	}

	w.isRunning = true
	w.Connection.debugLog("[Worker] Running loop")
	for {
		select {
		case p := <-w.buffer:
			w.Connection.debugLog("[Worker] sendPacket")
			w.sendPacket(p)
			w.wg.Done()
		case <-w.stopSignal:
			w.Connection.debugLog("[Worker] connection closing")
			err := w.Connection.close()
			if err != nil {
				w.Connection.logger.Errorf("[Worker] close error: %s", err.Error())
			}
			return
		}
	}
}

// sendPacket sends a packet to Instrumental API.
func (w *Worker) sendPacket(p packet) {
	w.Connection.writeWithTimeout(p.getBytes())
}
