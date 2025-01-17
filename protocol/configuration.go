package protocol

import (
	"fmt"
	"helper"
)

func HandleConfiguration(conn *connection, packet *incommingPacket) {
	switch packet.id {
	case 0x00:
		locale := packet.ReadString()
		viewDistance, _ := packet.ReadBytes(1)
		chatMode, _ := packet.ReadVarInt()
		charColors, _ := packet.ReadBoolean()
		displayedSkinParts, _ := packet.ReadBytes(1)
		mainHand, _ := packet.ReadVarInt()
		enableTextFiltering, _ := packet.ReadBoolean()
		allowServerListing, _ := packet.ReadBoolean()
		fmt.Println("[ConfSettings]", "id:", packet.id, "locale:", locale, "viewDistance:", viewDistance, "chatMode:", chatMode, "chatColor:", charColors, "displayedSkinParts:", displayedSkinParts, "mainHand:", mainHand, "enableTextFiltering:", enableTextFiltering, "allowServerListing:", allowServerListing)

		SendKnownPacks(conn)
	case 0x02:

		channel := packet.ReadString()
		data, _ := packet.ReadBytes(int(packet.length) - packet.offset)

		fmt.Println("[ConfCustomPayload]", "id:", packet.id, "channel:", string(channel), "data:", string(data), "length:", len(data))

	case 0x03:
		// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Configuration_Acknowledged
		// This packet switches the connection state to play.
		conn.state = helper.Play
		fmt.Printf("[ConfAck] id: %d\n", packet.id)

	case 0x07:
		packCount, _ := packet.ReadVarInt()

		for i := 1; i > int(packCount); i++ {
			namespace := packet.ReadString()
			ID := packet.ReadString()
			vers := packet.ReadString()

			fmt.Printf("%d -> %s, %s, %s", i, namespace, ID, vers)
		}

	default:
		fmt.Println("Unknown packet ID in Configuration state:", packet.id)
	}
}

func SendKnownPacks(conn *connection) {
	packet := outgouingPacket{
		id: 0x0E,
	}

	packet.WriteVarInt(0)

	conn.SendPacket(packet)
}

func SendFinish(conn *connection) {
	packet := outgouingPacket{
		id: 0x03,
	}

	conn.SendPacket(packet)
}
