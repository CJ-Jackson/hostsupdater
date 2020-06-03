package util

import (
	"os"
	"runtime"
)

func GetHostsPath() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("SystemRoot") + "\\System32\\drivers\\etc\\hosts"
	}

	return "/etc/hosts"
}
