package util

type Hostsfile struct {
	Name  string `json:"name"`
	Hosts []Host `json:"host"`
}

type Host struct {
	Comment string   `json:"comment"`
	Domains []string `json:"domains"`
}
