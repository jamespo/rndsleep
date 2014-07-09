package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os/exec"
	"strings"
	"time"
)

// generate random # up to *randmax
func genRand(randmax *int) int {
	bigrandmax := big.NewInt(int64(*randmax))
	randdelay, _ := rand.Int(rand.Reader, bigrandmax)
	return int(randdelay.Int64())
}

// Display msg if verbose set
func verboseMsg(msg string, verbose *bool) {
	if *verbose {
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	// get commandline flags
	randmax := flag.Int("randmax", 30, "maximum delay (seconds)")
	command := flag.String("command", "", "command")
	verbose := flag.Bool("verbose", false, "verbose")
	flag.Parse()

	// generate random delay
	var randdelayint = genRand(randmax)

	verboseMsg(fmt.Sprintf("Delaying %v for %v seconds", *command, randdelayint), verbose)

	// sleep
	time.Sleep(time.Duration(randdelayint) * time.Second) // prints 10s

	// run command
	if *command == "" {
		//fmt.Printf("No command supplied\n")
		verboseMsg("No command supplied", verbose)
	} else {
		// split cmd with arguments
		parts := strings.Fields(*command)
		cmd := parts[0]
		args := parts[1:len(parts)]
		out, err := exec.Command(cmd, args...).Output()
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		fmt.Printf("%s", out)
	}
}
