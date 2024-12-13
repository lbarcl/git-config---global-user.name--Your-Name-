package protocol

import (
	"config"
	"fmt"
	"helper"
)

func HandleStatus(conn *connection, packet *incommingPacket) {
	switch packet.id {
	case 0x00:
		// Get the config
		config := config.ReadConfig()

		outPacket := &outgouingPacket{
			id: 0x00,
		}

		outPacket.WriteString(fmt.Sprintf(`{"version":{"name":"1.21.1","protocol":767},"players":{"max":%d,"online":0},"description":"%s"}`, config.Game.MaxPlayers, config.Misc.Motd))

		conn.SendPacket(*outPacket)

		fmt.Println("[Status]", "id:", packet.id)

	case 0x01:
		// Read the payload of the ping packet (64-bit long)
		payload, err := packet.ReadLong()
		if err != nil {
			fmt.Println("Error reading payload Long:", err)
			break
		}

		outPacket := &outgouingPacket{
			id: 0x01,
		}

		outPacket.WriteLong(payload)

		conn.SendPacket(*outPacket)

		fmt.Println("[Status]", "id:", packet.id, "payload:", payload)

		conn.state = helper.Closed

	default:
		fmt.Println("Unknown packet ID in Status state:", packet.id)
	}
}
