package exec_test

import (
	"fmt"
	"log"

	"github.com/mrmarble/exec"
)

func ExampleCmd_Output() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
