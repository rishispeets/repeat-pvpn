package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	repeatConnAttempts()
}

func repeatConnAttempts() {
	for {
		connResult, connError := execPvpn()
		if connError != nil {
			// Log message and let loop retry conn after short wait
			printErr(connError)
			time.Sleep(1 * time.Second)
		} else {
			fmt.Printf("%s", connResult)
			break
		}
	}
}

// execPvpn returns the result of a protonvpn connection attempt
func execPvpn() (out []byte, err error) {
	cmd := exec.Command("/bin/sh", "-c", "pvpn -f")
	out, err = cmd.CombinedOutput()
	return
}

func printErr(err error) {
	fmt.Println("Error occurred: ")
	fmt.Printf("%s", err)
	fmt.Println("\nRetrying...")
}
