package protocol

import (
	"fmt"
	"helper"
	"io"
	"net"
)

type connection struct {
	socket          *net.Conn
	state           helper.States
	player          helper.Player
	compression     bool
	encrypted       bool
	protocolVersion int32
}

func (conn *connection) GetPacket() (*incommingPacket, error) {
	// Read the total packet length
	length, err := helper.ReadVarInt(*conn.socket)
	if err != nil {
		return nil, fmt.Errorf("failed to read packet length: %w", err)
	}

	fmt.Print("New Packet")
	fmt.Printf(" Total Packet Length: %d", length)

	// Read the packet ID
	id, err := helper.ReadVarInt(*conn.socket)
	if err != nil {
		return nil, fmt.Errorf("failed to read packet ID: %w", err)
	}

	fmt.Printf(" Total Packet ID: %d \n", id)

	reader := io.NopCloser(*conn.socket)

	// Construct and return the packet
	packet := &incommingPacket{
		id:     id,
		length: length,
		reader: reader,
		offset: helper.VarIntByteLength(id) + helper.VarIntByteLength(length),
	}

	return packet, nil
}

func (conn *connection) SendPacket(packet outgouingPacket) error {
	packetId := helper.WriteVarInt(packet.id)

	packetData := append(packetId, packet.data...)

	var rawPacket []byte

	packetLength := helper.WriteVarInt(int32(len(packetData)))
	rawPacket = append(packetLength, packetData...)

	if _, err := (*conn.socket).Write(rawPacket); err != nil {
		return fmt.Errorf("error sending packet: %w", err)
	}

	return nil
}
