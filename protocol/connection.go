package protocol

import (
	"fmt"
	"net"
)

func SocketHandle(conn net.Conn) {
	var currentState State = Handshaking

	for {
		// Parse the first VarInt
		packetLength, err := readVarInt(conn)
		if err != nil {
			fmt.Println("Connection closed")
			break
		}

		// Parse the second VarInt
		id, err := readVarInt(conn)
		if err != nil {
			fmt.Println("Error reading id VarInt:", err)
			break
		}
		fmt.Println("New Packet! - packetLength:", packetLength, "id:", id)

		// Handle the packet
		handlePacket(conn, &currentState, id)

	}
}
