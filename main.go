package main

import (
	"github.com/pngouin/defectdojo-cli/cmd"
	"github.com/pngouin/defectdojo-cli/config"
)

var (
	Version = "development"
)


func main() {
	config.LoadConfigFromEnv()
	cmd.Execute(Version)
}