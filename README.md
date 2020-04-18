# Interactive SSH
[![Go Report Card](https://goreportcard.com/badge/github.com/eze-kiel/interactive-ssh)](https://goreportcard.com/report/github.com/eze-kiel/interactive-ssh)

An interface to launch SSH connexions.

## Demo
[![asciicast](https://asciinema.org/a/E3D3zePVYguRDKujsypzhOblu.svg)](https://asciinema.org/a/E3D3zePVYguRDKujsypzhOblu)

## What's needed ?
You need to have `.ssh/config` file with hosts. For more informations, read this [article](https://www.devdungeon.com/content/ssh-tips).

## Usage
Once launched, the program will search for hosts in `~/.ssh/config`, and will list them in the prompt. Once you have selected the host you want to connect to, just press Enter. The program will quit and SSH will be launched.

## Author
Written by ezekiel, inspired by [DevDungeon's ISSH](https://www.devdungeon.com/content/issh).