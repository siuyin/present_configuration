package main

//010 OMIT
import (
	"fmt"
	"log"
	"os"

	"github.com/siuyin/dflt"
)

func main() {
	nSvr, err := dflt.EnvInt("NUM_SVR", 1)
	if err != nil {
		log.Fatalf("Bad NUM_SVR: %v\n", os.Getenv("NUM_SVR"))
	}
	fmt.Printf("I have %d servers\n", nSvr)
}

//020 OMIT
