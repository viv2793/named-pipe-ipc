package namedpipe

import (
	"io"
)

type NamedPipeServer interface {
	NewClient(requestId string) (err error)
	GetWriter(requestId string) (w io.Writer, err error)
	GetReader(requestId string) (r io.Reader, err error)
}

func NewNamedPipeServer(pipeName string) (namedPipe NamedPipeServer, err error) {
	return newNamedPipeServer(pipeName)
}

type NamedPipeClient interface {
	GetWriter() (w io.Writer, err error)
	GetReader() (r io.Reader, err error)
}

func NewNamedPipeClient(pipeName string, args ...string) (namedpipe NamedPipeClient, err error) {
	return newNamedPipeClient(pipeName, args...)
}
