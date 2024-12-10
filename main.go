package main

import (
	"config"
	"fmt"
	"log"
	"net"
	"protocol"
)

func startServer(config config.Conf) {
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Server.Ip, config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	fmt.Printf("[SERVER] started on %s \n", server.Addr())

	for {

		socket, err := server.Accept()
		fmt.Println(socket.RemoteAddr())
		if err != nil {
			fmt.Println(err)
		}

		protocol.SocketHandle(socket)
	}

}

func main() {
	serverConfig := config.ReadConfig()

	startServer(serverConfig)
}
