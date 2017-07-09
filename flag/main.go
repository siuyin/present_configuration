package main

import (
	"flag"
	"fmt"
)

// 010 OMIT
var (
	goEnv  string
	numSvr int
)

const (
	envUsage    = "environment this program will run in"
	numSvrUsage = "number of servers available"
)

func main() {
	flag.StringVar(&goEnv, "env", "dev", envUsage)
	flag.IntVar(&numSvr, "num-svr", 1, numSvrUsage)
	flag.Parse()
	fmt.Printf("Hi, I'm in the %s environment with %d server(s).\n", goEnv, numSvr)
}

// 020 OMIT
