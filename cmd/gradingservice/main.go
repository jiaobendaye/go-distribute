package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: registry.GradingServcie,
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
