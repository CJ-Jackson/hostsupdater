package util

type Hostsfile struct {
	Name  string `toml:"name"`
	Hosts []Host `toml:"host"`
}

type Host struct {
	Comment string   `toml:"comment"`
	Domains []string `toml:"domains"`
}
