package main

import ( 
	"time"
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// get commandline flags
	var randmax = flag.Int("randmax", 30, "maximum delay (seconds)")
	flag.Parse()

	// generate random delay
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randdelay := r.Intn(*randmax)

	//fmt.Printf("Delaying for %v seconds", randdelay)

	// sleep
	time.Sleep(time.Duration(randdelay)*time.Second) // prints 10s
	
}
