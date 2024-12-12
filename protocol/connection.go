package protocol

import (
	"bytes"
	"compress/zlib"
	"config"
	"fmt"
	"helper"
	"io"
	"net"
)

type connection struct {
	socket      *net.Conn
	state       helper.States
	player      helper.Player
	compression bool
	encrypted   bool
}

func (conn *connection) GetPacket() (*incommingPacket, error) {
	// Read the total packet length
	length, err := helper.ReadVarInt(*conn.socket)
	if err != nil {
		return nil, fmt.Errorf("failed to read packet length: %w", err)
	}

	fmt.Print("New Packet")
	fmt.Printf(" Total Packet Length: %d", length)

	if conn.compression {
		dataLength, err := helper.ReadVarInt(*conn.socket)
		if err != nil {
			return nil, fmt.Errorf("failed to read data length: %w", err)
		}

		fmt.Printf(" Total Data Length: %d ", dataLength)

		if dataLength == 0 {
			// Read the packet ID
			id, err := helper.ReadVarInt(*conn.socket)
			if err != nil {
				return nil, fmt.Errorf("failed to read packet ID 1: %w", err)
			}

			fmt.Printf(" Total Packet ID: %d \n", id)

			reader := io.NopCloser(*conn.socket)

			// Construct and return the packet
			packet := &incommingPacket{
				id:     id,
				length: length,
				reader: reader,
			}

			return packet, nil
		} else {
			reader, err := zlib.NewReader(*conn.socket)
			if err != nil {
				return nil, fmt.Errorf("failed to decompress paket: %w", err)
			}

			id, err := helper.ReadVarInt(reader)
			if err != nil {
				return nil, fmt.Errorf("failed to read packet ID 2: %w", err)
			}

			fmt.Printf(" Total Packet ID: %d \n", id)

			packet := &incommingPacket{
				id:     id,
				length: length,
				reader: reader,
			}

			return packet, nil
		}
	}

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
	}

	return packet, nil
}

func (conn *connection) SendPacket(packet outgouingPacket) error {
	packetId := helper.WriteVarInt(packet.id)

	packetData := append(packetId, packet.data...)

	var rawPacket []byte

	if conn.compression {
		threshold := config.ReadConfig().Server.NetworkCompressionThreshold

		if len(packetData) > int(threshold) {
			dataLength := helper.WriteVarInt(len(packetData))

			var buffer bytes.Buffer
			writer := zlib.NewWriter(&buffer)

			if _, err := writer.Write(packetData); err != nil {
				writer.Close() // Ensure the writer is closed on error
				return fmt.Errorf("failed to write data to compression writer: %w", err)
			}

			if err := writer.Close(); err != nil {
				return fmt.Errorf("failed to close compression writer: %w", err)
			}

			packetData = append(dataLength, buffer.Bytes()...)
		} else {
			packetData = append(helper.WriteVarInt(0), packetData...)
		}
	}

	packetLength := helper.WriteVarInt(len(packetData))
	rawPacket = append(packetLength, packetData...)

	if _, err := (*conn.socket).Write(rawPacket); err != nil {
		return fmt.Errorf("error sending packet: %w", err)
	}

	return nil
}
