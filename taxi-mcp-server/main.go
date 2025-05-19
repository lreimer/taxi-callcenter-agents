package main

import "github.com/lreimer/taxi-callcenter-agents/cmd"

var version string

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
