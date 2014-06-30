package main

import ( 
	"time"
	"flag"
	"fmt"
	"crypto/rand"
	"math/big"
)

func main() {
	// get commandline flags
	randmax := flag.Int("randmax", 30, "maximum delay (seconds)")
	flag.Parse()

	// generate random delay

	bigrandmax := big.NewInt(int64(*randmax))
	randdelay, _ := rand.Int(rand.Reader, bigrandmax)

	var randdelayint = int(randdelay.Int64())

	fmt.Printf("Delaying for %v seconds", randdelayint)

	// sleep
	time.Sleep(time.Duration(randdelayint)*time.Second) // prints 10s
	
}
