package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go serveHTTP()
	go serveStreams()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println(sig)
		done <- true
	}()
	log.Println("Server Start Awaiting Signal")
	log.Println("✅ 正在啟動 HTTP 伺服器...")

	<-done
	log.Println("Exiting")
}
