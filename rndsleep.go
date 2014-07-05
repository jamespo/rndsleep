package main

import ( 
	"time"
	"flag"
	"fmt"
	"crypto/rand"
	"math/big"
	"os/exec"
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

	// run command
	out, err := exec.Command(*command).Output()
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Printf("Output is %v\n", out)
}
