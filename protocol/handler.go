package protocol

import (
	"fmt"
	"helper"
	"net"
)

func SocketHandle(socket net.Conn) {
	conn := connection{
		socket:      &socket,
		state:       helper.Handshaking,
		compression: false,
		encrypted:   false,
	}

	for {
		packet, err := conn.GetPacket()

		if err != nil {
			conn.state = helper.Closed
			fmt.Printf("[SERVER] Client connection closed because of %s", err.Error())
			break
		}

		switch conn.state {
		case helper.Handshaking:
			// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Handshaking
			HandleHandshake(&conn, packet)
		case helper.Play:
			fmt.Println("Unhandled Play state")
		case helper.Status:
			// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Status
			HandleStatus(&conn, packet)
		case helper.Login:
			HandleLogin(&conn, packet)
		case helper.Configuration:
			HandleConfiguration(&conn, packet)
		}

		packet.Close()

		if conn.state == helper.Closed {
			break
		}
	}

	(*conn.socket).Close()
}
