package protocol

import (
	"helper"
	"io"
)

type incommingPacket struct {
	id     int
	length int
	reader io.ReadCloser
}

// Close ensures the packet's reader is closed properly.
func (packet *incommingPacket) Close() error {
	return packet.reader.Close()
}

// ReadVarInt reads a VarInt from the packet's reader or socket.
func (packet *incommingPacket) ReadVarInt() (int, error) {
	return helper.ReadVarInt(packet.reader)
}

// ReadShort reads a short (uint16) from the packet's sender.
func (packet *incommingPacket) ReadShort() (uint16, error) {
	return helper.ReadShort(packet.reader)
}

// ReadLong reads a long (int64) from the packet's sender.
func (packet *incommingPacket) ReadLong() (int64, error) {
	return helper.ReadLong(packet.reader)
}

// ReadString reads a string from the packet's sender.
func (packet *incommingPacket) ReadString() string {
	return helper.ReadString(packet.reader)
}

// ReadUUID reads a UUID string from the packet's sender.
func (packet *incommingPacket) ReadUUID() string {
	return helper.ReadUUID(packet.reader)
}

// ReadBytes reads a byte slice of the given length from the packet's sender.
func (packet *incommingPacket) ReadBytes(length int) ([]byte, error) {
	return helper.ReadBytes(packet.reader, length)
}

type outgouingPacket struct {
	id   int
	data []byte
}

func (packet *outgouingPacket) WriteVarInt(data int) {
	packet.data = append(packet.data, helper.WriteVarInt(data)...)
}

func (packet *outgouingPacket) WriteShort(data int16) {
	packet.data = append(packet.data, byte(data))
}

func (packet *outgouingPacket) WriteLong(data int64) {
	packet.data = append(packet.data, helper.WriteLong(data)...)
}

func (packet *outgouingPacket) WriteString(data string) {
	packet.data = append(packet.data, helper.WriteString(data)...)
}

func (packet *outgouingPacket) WriteUUID(uuid string) {
	packet.data = append(packet.data, helper.WriteUUID(uuid)...)
}
