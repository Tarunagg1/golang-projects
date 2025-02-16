package server

import (
	"io"
	"log"
	"net"
	"strconv"
)

func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)

	n, err := c.Read(buf[:])

	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}

func RunSyncTCPServer() {
	log.Println("starting a syns tcp servrr")

	var con_client int = 0

	lsnr, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(5000))

	if err != nil {
		panic(err)
	}

	for {
		c, err := lsnr.Accept()

		if err != nil {
			panic(err)
		}

		con_client += 1

		log.Println("client connected with address: ", c.RemoteAddr(), "Concurrent", con_client)

		for {

			cmd, err := readCommand(c)

			if err != nil {
				c.Close()
				con_client -= 1
				log.Println("client disconnected", c.RemoteAddr(), "conncurrent clien", con_client)

				if err == io.EOF {
					break
				}
				log.Println("error", err)
			}
			log.Println("command", cmd)

			if err = respond(cmd, c); err != nil {
				log.Println("error", err)
			}
		}
	}
}
