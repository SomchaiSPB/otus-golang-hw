package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Client struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	connection net.Conn
}

func (c *Client) Close() (err error) {
	err = c.connection.Close()
	return err
}

func (c *Client) Receive() error {
	_, err := io.Copy(c.out, c.connection)
	return err
}

func (c *Client) Send() error {
	_, err := io.Copy(c.connection, c.in)
	return err
}

func (c *Client) Connect() (err error) {
	c.connection, err = net.DialTimeout("tcp", c.address, c.timeout)
	_, _ = fmt.Fprintf(os.Stdout, "...Connected to %s\n", c.address)
	return err
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
