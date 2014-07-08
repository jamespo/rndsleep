package main

import ( 
	"time"
	"flag"
	"fmt"
	"crypto/rand"
	"math/big"
	"os/exec"
	"strings"
)

func main() {
	// get commandline flags
	randmax := flag.Int("randmax", 30, "maximum delay (seconds)")
	command := flag.String("command", "ls", "command")
	flag.Parse()

	// generate random delay

	bigrandmax := big.NewInt(int64(*randmax))
	randdelay, _ := rand.Int(rand.Reader, bigrandmax)

	var randdelayint = int(randdelay.Int64())

	fmt.Printf("Delaying %v for %v seconds\n", *command, randdelayint)

	// sleep
	time.Sleep(time.Duration(randdelayint)*time.Second) // prints 10s

	// split cmd with arguments
	parts := strings.Fields(*command)
	head := parts[0]
        parts = parts[1:len(parts)]

	// run command
	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	fmt.Printf("Output is %s", out)
}
