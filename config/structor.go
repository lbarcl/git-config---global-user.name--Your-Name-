package config

type Conf struct {
	Server Server `yaml:"server"`
	Misc   Misc   `yaml:"misc"`
	Game   Game   `yaml:"game"`
}

type Server struct {
	Ip                          string `yaml:"server-ip"`
	Port                        uint16 `yaml:"server-port"`
	Online                      bool   `yaml:"online-mode"`
	Status                      bool   `yaml:"enable-status"`
	LogIPs                      bool   `yaml:"log-ips"`
	RateLimit                   uint32 `yaml:"rate-limit"`
	NetworkCompressionThreshold uint   `yaml:"network-compression-threshold"`
	EnableCompression           bool   `yaml:"enable-compression"`
}

type Misc struct {
	Motd string `yaml:"motd"`
}

type Game struct {
	MaxPlayers int `yaml:"max-players"`
}
