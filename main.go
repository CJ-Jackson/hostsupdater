package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/CJ-Jackson/hostsupdater/util"
	toml "github.com/pelletier/go-toml"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must have one argument.")
		return
	}

	ip := os.Args[1]
	file, err := os.OpenFile(util.GetHostsPath(), os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	hostsFile, err := os.Open(wd + "/hosts.toml")
	if err != nil {
		log.Fatal(err)
	}
	hostFilesData := util.Hostsfile{}

	err = toml.NewDecoder(hostsFile).Decode(&hostFilesData)
	if err != nil {
		log.Fatal(err)
	}
	hostsFile.Close()

	hostsBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(0, io.SeekStart)

	hostsBytes = removeOldSetting(hostsBytes, hostFilesData.Name)

	err = file.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(hostsBytes)
	if err != nil {
		log.Fatal(err)
	}

	if ip == "--" {
		return
	}

	template.Must(template.New("HostFile Template").Parse(util.Template)).Execute(file, util.TemplateData{
		Ip:        ip,
		Hostsfile: hostFilesData,
	})
}

func removeOldSetting(hostBytes []byte, name string) []byte {
	pos := bytes.Index(hostBytes, []byte(fmt.Sprintf("# < %s", name)))
	if pos == -1 {
		return hostBytes
	}

	buf := &bytes.Buffer{}
	buf.Write(hostBytes[:pos])

	endB := []byte(fmt.Sprintf("# %s >", name))
	pos = bytes.Index(hostBytes, endB)
	if pos == -1 {
		log.Panic("Missing EndPoint")
	}

	buf.Write(hostBytes[pos+len(endB):])

	return buf.Bytes()
}
