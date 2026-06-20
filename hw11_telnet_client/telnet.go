package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type ClientImpl struct {
	address string
	conn    net.Conn
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func (c *ClientImpl) Connect() error {
	return nil
}

func (c *ClientImpl) Send() error {
	return nil
}

func (c *ClientImpl) Receive() error {
	return nil
}

func (c *ClientImpl) Close() error {
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &ClientImpl{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
