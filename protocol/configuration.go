package protocol

import (
	"fmt"
	"helper"
)

func HandleConfiguration(conn *connection, packet *incommingPacket) {
	switch packet.id {
	case 0x02:

		channelLength, _ := packet.ReadVarInt()
		channel, _ := packet.ReadBytes(channelLength)
		data, _ := packet.ReadBytes(packet.length - len(helper.WriteVarInt(channelLength)) - channelLength)

		fmt.Println("[ConfCustomPayload]", "id:", packet.id, "channel:", string(channel), "data:", string(data), "length:", len(data))

	case 0x05:

		// id, _ := packet.ReadInt()

		fmt.Println("[ConfPong]", "id:", packet.id)

	default:
		fmt.Println("Unknown packet ID in Configuration state:", packet.id)
	}
}
