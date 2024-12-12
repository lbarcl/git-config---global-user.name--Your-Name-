package protocol

import (
	"fmt"
	"helper"
	"net"
)

func SocketHandle(conn net.Conn) {
	var currentState helper.States = helper.Handshaking

	for {
		packet, err := GetPacket(conn)
		if err != nil {
			fmt.Println("Connection closed")
		}

		fmt.Println("New Packet! - packetLength:", packet.length, "id:", packet.id)

		switch currentState {
		case helper.Handshaking:
			// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Handshaking
			HandleHandshake(&currentState, *packet)
		case helper.Play:
			fmt.Println("Unhandled Play state")
		case helper.Status:
			// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Status
			HandleStatus(&currentState, *packet)
		case helper.Login:
			HandleLogin(&currentState, *packet)
		}

		if currentState == helper.Closed {
			break
		}
	}
}
