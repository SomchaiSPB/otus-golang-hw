package main

import (
	"context"
	"github.com/spf13/pflag"
	"log"
	"os"
	"sync"
	"time"
)

var (
	wg      = sync.WaitGroup{}
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

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	client := NewTelnetClient(host+":"+port, timeout, os.Stdin, os.Stdout)

	err := client.Connect()

	if err != nil {
		log.Println(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := client.Send()

		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := client.Receive()

		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	wg.Wait()

	client.Close()

	<-ctx.Done()
}

func init() {
	pflag.DurationVar(&timeout, "timeout", 10*time.Second, "connection timeout")
	pflag.Parse()
}
