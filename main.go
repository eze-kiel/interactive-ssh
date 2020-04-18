package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func main() {
	hosts, err := getHosts()
	hosts = append(hosts, "quit")

	prompt := promptui.Select{
		Label: "Select host",
		Items: hosts,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if result == "quit" {
		color.Red("[*] exiting...")
	} else {
		color.Green("[*] launching ssh on: %s\n", result)

		binary, lookErr := exec.LookPath("ssh")
		if lookErr != nil {
			panic(lookErr)
		}
		err := syscall.Exec(binary, []string{"ssh", result}, os.Environ())

		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
	}
}

func getHosts() ([]string, error) {
	var hostsList []string

	configFile, err := ioutil.ReadFile("/home/ezekiel/.ssh/config")
	if err != nil {
		return nil, err
	}

	configFileStr := string(configFile)
	hostsList = strings.Split(configFileStr, "\n")
	var newHL []string

	for i := 0; i < len(hostsList); i++ {
		if strings.Contains(hostsList[i], "Host ") {
			newHL = append(newHL, hostsList[i])
		}
	}

	for i := 0; i < len(newHL); i++ {
		newHL[i] = strings.ReplaceAll(newHL[i], "Host ", "")
	}
	return newHL, nil
}
