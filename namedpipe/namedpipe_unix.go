// +build !windows

package namedpipe

import (
	"errors"
	"io"
)

type UnixNamedPipeServer struct {
}

func newNamedPipeServer(pipeName string) (namedPipe NamedPipeServer, err error) {
	return nil, errors.New("not implemented")
}

func (s *UnixNamedPipeServer) GetReader(requestId string) (reader io.Reader, err error) {
	return nil, errors.New("not implemented")
}

func (s *UnixNamedPipeServer) GetWriter(requestId string) (writer io.Writer, err error) {
	return nil, errors.New("not implemented")
}

func (s *UnixNamedPipeServer) NewClient(requestId string) (err error) {
	return errors.New("not implemented")
}

type UnixNamedPipeClient struct {
}

func newNamedPipeClient(pipeName string, args ...string) (namedPipe NamedPipeClient, err error) {
	return nil, errors.New("not implemented")
}

func (c *UnixNamedPipeClient) GetReader() (r io.Reader, err error) {
	return nil, errors.New("not implemented")
}

func (c *UnixNamedPipeClient) GetWriter() (w io.Writer, err error) {
	return nil, errors.New("not implemented")
}