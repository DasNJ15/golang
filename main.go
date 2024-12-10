package main

import (
	"fmt"
	"lcorequest/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//helloHandler := handler.NewHello()
	//byeHandler := handler.NewBye()
	prHandler := handler.Newpr()

	sm := http.NewServeMux()

	sm.Handle("/", prHandler)

	newServer := &http.Server{
		Addr:         ":8000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		newServer.ListenAndServe()
	}()

	iochannel := make(chan os.Signal, 1)
	signal.Notify(iochannel, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Waiting for Intrupt..")
	fmt.Println("Waiting for Intrupt..")
	sig := <-iochannel
	fmt.Println("Intrupt by: ", sig)

}
