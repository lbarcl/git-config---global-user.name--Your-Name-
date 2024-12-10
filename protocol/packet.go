package protocol

import (
	"config"
	"fmt"
	"net"
)

type State uint8

const (
	Handshaking State = iota
	Status
	Login
	Play
)

func handlePacket(conn net.Conn, state *State, id int) {

	switch *state {
	// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Handshaking
	case Handshaking:
		switch id {
		case 0x00:
			// Parse the second VarInt
			protocolversion, err := readVarInt(conn)
			if err != nil {
				fmt.Println("Error reading PV VarInt:", err)
				break
			}

			serverAdress := readString(conn)

			serverPort, err := readShort(conn)
			if err != nil {
				fmt.Println("Error reading sp VarInt:", err)
				break
			}

			nextState, err := readVarInt(conn)
			if err != nil {
				fmt.Println("Error reading NS VarInt:", err)
				break
			}

			fmt.Println("[Handshake]", "id:", id, "protocol version:", protocolversion, "server address:", serverAdress, "server port:", serverPort, "next state:", nextState)

			*state = State(nextState)

		default:
			fmt.Println("Unknown packet ID in Handshaking state:", id)
		}
	case Play:
		fmt.Println("Unhandled Play state")
	case Status:
		// https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Status
		switch id {
		case 0x00:
			// Get the config
			config := config.ReadConfig()

			// Raw JSON response
			response := []byte(fmt.Sprintf(`{"version":{"name":"1.21.1","protocol":767},"players":{"max":%d,"online":0},"description":"%s"}`, config.Game.MaxPlayers, config.Misc.Motd))

			// length of the response
			responseLength := writeVarInt(len(response))

			// Combine the length and the response
			responseData := append(responseLength, response...)

			// Send response
			sendPacket(conn, 0x00, responseData)

			fmt.Println("[Status]", "id:", id)

		case 0x01:
			// Read the payload of the ping packet (64-bit long)
			payload, err := readLong(conn)
			if err != nil {
				fmt.Println("Error reading payload Long:", err)
				break
			}

			// Send the pong packet
			sendPacket(conn, 0x01, writeLong(payload))

			fmt.Println("[Status]", "id:", id, "payload:", payload)

		default:
			fmt.Println("Unknown packet ID in Status state:", id)
		}
	case Login:
		fmt.Println("Unhandled Login state")
	}

}

func sendPacket(conn net.Conn, id int, data []byte) {

	// Write the packet ID as a VarInt
	packetIdEncoded := writeVarInt(id)

	// Append the packet ID to the data
	packetData := append(packetIdEncoded, data...)

	// Write the length of the packet data as a VarInt
	packetLength := writeVarInt(len(packetData))

	// Combine the packet length and the packet data
	packet := append(packetLength, packetData...)

	// Send the packet
	_, err := conn.Write(packet)
	if err != nil {
		fmt.Println("Error sending packet:", err)
	}
}
