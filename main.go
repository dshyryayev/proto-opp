package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"realtimemap-go/backend/protocluster"
	"realtimemap-go/backend/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	stopOnSignals(cancel)

	cluster := protocluster.StartNode()
	time.Sleep(2 * time.Second)

	server := server.NewHttpServer(cluster, ctx)
	serverDone := server.ListenAndServe()

	<-serverDone

	cluster.Shutdown(true)
}

func stopOnSignals(cancel func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		fmt.Println("*** STOPPING ***")
		cancel()
	}()
}
