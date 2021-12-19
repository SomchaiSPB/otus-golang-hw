package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/pflag"
)

var (
	host    string
	port    string
	timeout time.Duration
)

func main() {
	host = pflag.Arg(0)
	port = pflag.Arg(1)

	if host == "" || port == "" {
		log.Fatal("host and port are mandatory")
	}
	client := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	defer client.Close()

	go func() {
		err := client.Receive()
		if err == nil {
			fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
		} else {
			log.Println(err)
		}
		cancel()
	}()

	go func() {
		err := client.Send()
		if err == nil {
			fmt.Fprintln(os.Stderr, "...EOF")
		} else {
			log.Println(err)
		}
		cancel()
	}()

	<-ctx.Done()
}

func init() {
	pflag.DurationVar(&timeout, "timeout", 10*time.Second, "connection timeout")
	pflag.Parse()
}
