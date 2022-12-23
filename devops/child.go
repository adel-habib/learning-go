package main

import (
	"fmt"
	"os"
	"os/exec"
)

// cmd.Run -> excutes command the returns an error if there was one
// cmd.Start -> starts the process asynchronously and lets the parent process (this program for example) continue its flow
// cmd.Wait waits on process started by cmd.Start
// cmd.Output runs the command standalone and returns its result

func main() {
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Cmd: ", cmd.Args[0])
	fmt.Println("Args:", cmd.Args[1:])
	fmt.Println("PID: ", cmd.Process.Pid)
	cmd.Wait()
}
