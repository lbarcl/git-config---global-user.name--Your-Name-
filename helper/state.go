package helper

type States uint8

const (
	Handshaking States = iota
	Status
	Login
	Play
	Closed
)
