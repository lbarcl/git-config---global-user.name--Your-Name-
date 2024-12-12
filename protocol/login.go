package protocol

import (
	"fmt"
	"helper"
)

func HandleLogin(state *helper.States, packet Packet) {
	switch packet.id {
	case 0x00:
		username := packet.ReadString()
		uuid, _ := packet.ReadUUID()
		fmt.Printf("[Server] %s joined with UUID of %s", username, uuid)
	}
}
