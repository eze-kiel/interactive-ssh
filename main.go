package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func main() {
	hosts, err := getHosts()
	if err != nil {
		log.Fatal(err)
	}
	hosts = append(hosts, "quit iSSH")

	prompt := promptui.Select{
		Label: "Select a SSH host",
		Items: hosts,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if result == "quit iSSH" {
		color.Red("[*] exiting...")
		os.Exit(0)
	}

	color.Green("[*] opening ssh connection on: %s\n", result)
	binary, err := exec.LookPath("ssh")
	if err != nil {
		log.Fatal(err)
	}

	if err := syscall.Exec(binary, []string{"ssh", result}, os.Environ()); err != nil {
		log.Fatal(err)
	}
}

func getHosts() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configFile, err := ioutil.ReadFile(homeDir + "/.ssh/config")
	if err != nil {
		return nil, err
	}

	hostsList := strings.Split(string(configFile), "\n")
	var newHostsList []string

	for i := 0; i < len(hostsList); i++ {
		if strings.Contains(hostsList[i], "Host ") {
			newHostsList = append(newHostsList, hostsList[i])
		}
	}

	for i := 0; i < len(newHostsList); i++ {
		newHostsList[i] = strings.ReplaceAll(newHostsList[i], "Host ", "")
	}
	sort.Strings(newHostsList)
	return newHostsList, nil
}
