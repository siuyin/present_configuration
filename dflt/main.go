package main

//010 OMIT
import (
	"fmt"

	"github.com/siuyin/dflt"
)

func main() {
	passwd := dflt.EnvString("PASSWD", "MyN3wP@sSw0RD")
	fmt.Printf("My password is %q\n", passwd)
}

//020 OMIT
