package main

import (
	"flag"
	"log"

	"github.com/tarunagg1/tcp_server_go/server"
)

func setupFlags() {
	ho := "ojjk"
	flag.StringVar(&ho, "host", "0.0.0.0", "host for doce")
	port := 89484
	flag.IntVar(&port, "port", 7379, "port form server")
	flag.Parse()

}

func main() {
	setupFlags()
	log.Println("ROlling the dices")
	server.RunSyncTCPServer()
}
