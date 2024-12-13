package protocol

import (
	"helper"
	"io"
)

type incommingPacket struct {
	id     int
	length int
	offset int
	reader io.ReadCloser
}

// Close ensures the packet's reader is closed properly.
func (packet *incommingPacket) Close() error {
	return packet.reader.Close()
}

// ReadVarInt reads a VarInt from the packet's reader or socket.
func (packet *incommingPacket) ReadVarInt() (int, error) {
	value, err := helper.ReadVarInt(packet.reader)
	if err != nil {
		return 0, err
	}
	packet.offset += helper.VarIntByteLength(value)
	return value, nil
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
	str := helper.ReadString(packet.reader)
	packet.offset += len([]byte(str))
	return str
}

// ReadUUID reads a UUID string from the packet's sender.
func (packet *incommingPacket) ReadUUID() string {
	return helper.ReadUUID(packet.reader)
}

// ReadBytes reads a byte slice of the given length from the packet's sender.
func (packet *incommingPacket) ReadBytes(length int) ([]byte, error) {
	bytes, err := helper.ReadBytes(packet.reader, length)
	packet.offset += length
	return bytes, err
}

// ReadInt reads an int from the packet's sender.
func (packet *incommingPacket) ReadInt() (int, error) {
	return helper.ReadInt(packet.reader)
}

// ReadBoolean reads a boolean from the packet's sender.
func (packet *incommingPacket) ReadBoolean() (bool, error) {
	boolean, err := helper.ReadBoolean(packet.reader)
	packet.offset++
	return boolean, err
}

// Read all remaining bytes
func (packet *incommingPacket) ReadAll() ([]byte, error) {
	return io.ReadAll(packet.reader)
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
