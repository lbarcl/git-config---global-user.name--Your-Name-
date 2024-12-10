package config

type Conf struct {
	Server Server `yaml:"server"`
	Misc   Misc   `yaml:"misc"`
}

type Server struct {
	Ip        string `yaml:"server-ip"`
	Port      uint16 `yaml:"server-port"`
	Online    bool   `yaml:"online-mode"`
	Status    bool   `yaml:"enable-status"`
	LogIPs    bool   `yaml:"log-ips"`
	RateLimit uint32 `yaml:"rate-limit"`
}

type Misc struct {
	Motd string `yaml:"motd"`
}
