package protocol

import (
	"config"
	"fmt"
	"helper"
)

func HandleStatus(state *helper.States, packet Packet) {
	switch packet.id {
	case 0x00:
		// Get the config
		config := config.ReadConfig()

		// Raw JSON response
		response := []byte(fmt.Sprintf(`{"version":{"name":"1.21.1","protocol":767},"players":{"max":%d,"online":0},"description":"%s"}`, config.Game.MaxPlayers, config.Misc.Motd))

		// length of the response
		responseLength := helper.WriteVarInt(len(response))

		// Combine the length and the response
		responseData := append(responseLength, response...)

		// Send response
		SendPacket(*packet.sender, 0x00, responseData)

		fmt.Println("[Status]", "id:", packet.id)

	case 0x01:
		// Read the payload of the ping packet (64-bit long)
		payload, err := packet.ReadLong()
		if err != nil {
			fmt.Println("Error reading payload Long:", err)
			break
		}

		// Send the pong packet
		SendPacket(*packet.sender, 0x01, helper.WriteLong(payload))

		fmt.Println("[Status]", "id:", packet.id, "payload:", payload)

		*state = helper.Closed

	default:
		fmt.Println("Unknown packet ID in Status state:", packet.id)
	}
}
