package main

import (
	"fmt"
	"os/exec"
)

func main() {
	execPvpn()
}

// execPvpn returns the result of a protonvpn connection attempt
func execPvpn() {
	cmd := exec.Command("/bin/sh", "-c", "sudo pvpn -f")
	out, err := cmd.CombinedOutput()

	if err != nil {
		printErr(err)
	}

	fmt.Printf("%s", out)
}

// printErr prints the given error message
func printErr(err error) {
	fmt.Println("Error occurred")
	fmt.Printf("%s", err)
}
