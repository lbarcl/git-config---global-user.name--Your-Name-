package protocol

import (
	"config"
	"fmt"
	"helper"
)

func HandleLogin(conn *connection, packet *incommingPacket) {
	switch packet.id {
	case 0x00:
		username := packet.ReadString()
		uuid := packet.ReadUUID()

		conn.player.UUID = uuid
		conn.player.Username = username

		fmt.Printf("[Server] %s joined with UUID of %s\n", username, uuid)

		if config.ReadConfig().Server.EnableCompression {
			outComPacket := &outgouingPacket{
				id: 0x03,
			}

			outComPacket.WriteVarInt(int(config.ReadConfig().Server.NetworkCompressionThreshold))

			conn.SendPacket(*outComPacket)

			conn.compression = true
		}

		// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Login_Success
		outPacket := &outgouingPacket{
			id: 0x02,
		}

		outPacket.WriteUUID(uuid)
		outPacket.WriteString(username)

		outPacket.WriteVarInt(0) // Number of properties
		outPacket.WriteVarInt(0) // Strict error

		conn.SendPacket(*outPacket)
		fmt.Printf("[LoginSuccess] id: %d uuid: %s username: %s\n", packet.id, uuid, username)

	case 0x03:
		// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Login_Acknowledged
		// This packet switches the connection state to configuration.
		conn.state = helper.Configuration
		fmt.Printf("[LoginAck] id: %d uuid: %s username: %s\n", packet.id, conn.player.UUID, conn.player.Username)
	}
}
