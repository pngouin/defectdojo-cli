package main

import (
	"github.com/pngouin/defectdojo-cli/cmd"
	"github.com/pngouin/defectdojo-cli/config"
)

func main() {
	config.LoadConfigFromEnv()
	cmd.Execute()
}