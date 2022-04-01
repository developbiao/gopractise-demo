package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		err    error
		output []byte
	)
	cmd = exec.Command("/bin/bash", "-c", "echo 1; ps;")

	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	// Output execute result
	fmt.Println("output:", string(output))

}
