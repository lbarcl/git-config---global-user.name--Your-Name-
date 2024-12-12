package main

import (
	"config"
	"fmt"
	"log"
	"net"
	"protocol"
	"time"
)

func startServer(config config.Conf) {
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Server.Ip, config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	fmt.Println("[SERVER]", time.Now().Format("02-01-2006 15:04:05"))
	fmt.Println("[SERVER] MoGo, using 'server.yaml' as configuration file")
	fmt.Println("[SERVER] Starting server at", server.Addr())
	fmt.Println("[SERVER] Quit the server with CONTROL-C")

	for {
		socket, err := server.Accept()
		fmt.Println("=====================================")
		fmt.Println(socket.RemoteAddr())
		if err != nil {
			fmt.Println(err)
		}

		go protocol.SocketHandle(socket)
	}

}

func main() {
	serverConfig := config.ReadConfig()
	config.GetEncryption()

	startServer(serverConfig)
}
