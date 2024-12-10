package protocol

import (
	"fmt"
	"net"
)

func SocketHandle(conn net.Conn) {
	defer conn.Close()

	// Parse the first VarInt
	packetLength, err := readVarInt(conn)
	if err != nil {
		fmt.Println("Error reading packetLength VarInt:", err)
		return
	}

	// Parse the second VarInt
	id, err := readVarInt(conn)
	if err != nil {
		fmt.Println("Error reading id VarInt:", err)
		return
	}

	// Parse the second VarInt
	protocolversion, err := readVarInt(conn)
	if err != nil {
		fmt.Println("Error reading PV VarInt:", err)
		return
	}

	serverAdress := readString(conn)

	serverPort, err := readShort(conn)
	if err != nil {
		fmt.Println("Error reading sp VarInt:", err)
		return
	}

	nextState, err := readVarInt(conn)
	if err != nil {
		fmt.Println("Error reading NS VarInt:", err)
		return
	}

	fmt.Println("Parsed VarInts - packetLength:", packetLength, "id:", id, "protocol version:", protocolversion, "server address:", serverAdress, "server port:", serverPort, "next state:", nextState)
}
