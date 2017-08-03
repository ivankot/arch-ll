

package main

import "fmt"
import "os"
import "os/exec"
import "log"

func main() {
	var arg string
	arg = "."
	if (len(os.Args) > 1) {
		arg = os.Args[1]
	} 
	cmd := exec.Command("ls", "-lha", arg)
	output, err := cmd.Output()
	if (err != nil) {
		log.Fatalln("Something went wrong")
	}
	fmt.Printf("%s", output)
}
