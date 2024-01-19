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

func ExampleCmd_CombinedOutput() {
	out, err := exec.Command("/bin/sh", "-c", "echo stdout; echo 1>&2 stderr").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %s\n", out)
	// Output: Output: stdout
	// stderr
}
