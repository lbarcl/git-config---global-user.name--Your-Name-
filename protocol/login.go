package protocol

import (
	"fmt"
	"helper"
)

func HandleLogin(state *helper.States, packet Packet, player *helper.Player) {
	switch packet.id {
	case 0x00:
		username := packet.ReadString()
		uuid := packet.ReadUUID()

		player.Username = username
		player.UUID = uuid

		fmt.Printf("[Server] %s joined with UUID of %s\n", username, uuid)

		rawBytes := []byte{}

		// Set the player UUID and username
		rawBytes = append(rawBytes, helper.WriteUUID(player.UUID)...)
		rawBytes = append(rawBytes, helper.WriteString(player.Username)...)

		// Number of properties
		rawBytes = append(rawBytes, helper.WriteVarInt(0)...)

		// Strict error handling
		rawBytes = append(rawBytes, helper.WriteVarInt(0)...)

		SendPacket(*packet.sender, 0x02, rawBytes)
		fmt.Printf("[LoginSuccess] id: %d uuid: %s username: %s\n", packet.id, player.UUID, player.Username)

	case 0x03:
		// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Login_Acknowledged
		// This packet switches the connection state to configuration.
		*state = helper.Configuration
		fmt.Printf("[LoginAck] id: %d uuid: %s username: %s\n", packet.id, player.UUID, player.Username)
	}

}
