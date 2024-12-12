package protocol

import (
	"fmt"
	"helper"
	"net"
)

type Packet struct {
	id     int
	length int
	sender *net.Conn
}

func GetPacket(socket net.Conn) (*Packet, error) {
	// Read the total packet length
	length, err := helper.ReadVarInt(socket)
	if err != nil {
		return nil, fmt.Errorf("failed to read packet length: %w", err)
	}

	// Read the packet ID
	id, err := helper.ReadVarInt(socket)
	if err != nil {
		return nil, fmt.Errorf("failed to read packet ID: %w", err)
	}

	// Construct and return the packet
	packet := &Packet{
		id:     id,
		length: length,
		sender: &socket,
	}

	return packet, nil
}

func (packet *Packet) ReadVarInt() (int, error) {
	return helper.ReadVarInt(*packet.sender)
}

func (packet *Packet) ReadShort() (uint16, error) {
	return helper.ReadShort(*packet.sender)
}

func (packet *Packet) ReadLong() (int64, error) {
	return helper.ReadLong(*packet.sender)
}

func (packet *Packet) ReadString() string {
	return helper.ReadString(*packet.sender)
}

func (packet *Packet) ReadUUID() string {
	return helper.ReadUUID(*packet.sender)
}

func (packet *Packet) ReadBytes(length int) ([]byte, error) {
	return helper.ReadBytes(*packet.sender, length)
}

func SendPacket(conn net.Conn, id int, data []byte) {

	// Write the packet ID as a VarInt
	packetIdEncoded := helper.WriteVarInt(id)

	// Append the packet ID to the data
	packetData := append(packetIdEncoded, data...)

	// Write the length of the packet data as a VarInt
	packetLength := helper.WriteVarInt(len(packetData))

	// Combine the packet length and the packet data
	packet := append(packetLength, packetData...)

	// Send the packet
	_, err := conn.Write(packet)
	if err != nil {
		fmt.Println("Error sending packet:", err)
	}
}
