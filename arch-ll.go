//    Copyright 2017 Ivan Kot
// 
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
// 
//        http://www.apache.org/licenses/LICENSE-2.0
// 
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

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
