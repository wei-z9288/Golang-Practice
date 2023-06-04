package main

import (
	cli "fops/cli"
)

func main() {
	var testString = "v0.0.1"
	cli.SetVersion(testString)
	cli.Execute()
}
