// +build windows

package namedpipe

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/Microsoft/go-winio"
)

type WinNamedPipeServer struct {
	pipeName      string
	pipeListener  net.Listener
	activeClients map[string]net.Conn
	mu            sync.Mutex
}

func newNamedPipeServer(pipeName string) (namedPipe NamedPipeServer, err error) {
	pipePath := fmt.Sprintf(`\\.\pipe\%s`, pipeName)
	pipeListener, err := winio.ListenPipe(pipePath, nil)
	if err != nil {
		return nil, err
	}
	pipe := &WinNamedPipeServer{
		pipeName:      pipeName,
		pipeListener:  pipeListener,
		activeClients: make(map[string]net.Conn),
	}

	return pipe, nil
}

func (s *WinNamedPipeServer) GetReader(requestId string) (r io.Reader, err error) {
	return s.activeClients[requestId], nil
}

func (s *WinNamedPipeServer) GetWriter(requestId string) (w io.Writer, err error) {
	return s.activeClients[requestId], nil
}

func (s *WinNamedPipeServer) NewClient(requestId string) (err error) {
	// Accept a pipe connection from a client. pipeListener.Accept is a
	// blocking call until a client connects to the named pipe
	// When a new client connects to the named pipe, a new "instance" of
	// the pipe is created
	conn, err := s.pipeListener.Accept()
	if err != nil {
		return err
	}
	s.mu.Lock()
	s.activeClients[requestId] = conn
	s.mu.Unlock()
	return nil
}

type WinNamedPipeClient struct {
	pipeName string
	pipeConn net.Conn
}

func newNamedPipeClient(pipeName string, args ...string) (namedpipe NamedPipeClient, err error) {
	pipePath := fmt.Sprintf(`\\.\pipe\%s`, pipeName)
	duration := 10 * time.Second
	conn, err := winio.DialPipe(pipePath, &duration)
	if err != nil {
		return nil, err
	}
	return &WinNamedPipeClient{
		pipeName: pipeName,
		pipeConn: conn,
	}, nil
}

func (c *WinNamedPipeClient) GetWriter() (w io.Writer, err error) {
	return c.pipeConn, nil
}

// Read message from the named pipe
func (c *WinNamedPipeClient) GetReader() (r io.Reader, err error) {
	return c.pipeConn, nil
}
