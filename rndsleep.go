package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

func server(port *int) {
	// listen on a port
	_, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		fmt.Println("Can't open lock socket")
		fmt.Println(err)
		os.Exit(1)
	}
}

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

func runCmd(command *string) ([]byte, error) {
	// split cmd with arguments
	parts := strings.Fields(*command)
	cmd := parts[0]
	args := parts[1:len(parts)]
	out, err := exec.Command(cmd, args...).Output()
	return out, err
}

func main() {
	// get commandline flags
	randmax := flag.Int("randmax", 0, "maximum delay (seconds)")
	command := flag.String("command", "", "command to run")
	verbose := flag.Bool("verbose", false, "enable verbose output")
	port := flag.Int("port", 0, "localhost TCP port to lock on")
	flag.Parse()

	if *port != 0 {
		// open socket
		verboseMsg(fmt.Sprintf("Locking on 127.0.0.1:%d", *port), verbose)
		go server(port)
	}

	if *randmax > 0 {
		// generate random delay
		var randdelayint = genRand(randmax)
		verboseMsg(fmt.Sprintf("Delaying %v for %v seconds", *command, randdelayint), verbose)
		// sleep
		time.Sleep(time.Duration(randdelayint) * time.Second) // prints 10s
	}

	// run command
	if *command == "" {
		//fmt.Printf("No command supplied\n")
		verboseMsg("No command supplied", verbose)
	} else {
		verboseMsg(fmt.Sprintf("Running %s\n", *command), verbose)
		out, err := runCmd(command)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		fmt.Printf("%s", out)
	}
}
