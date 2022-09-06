package utils

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

// ExecCommand
func ExecCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Println("Error: The command is err: ", err.Error())
		return err
	}

	ch := make(chan string, 100)
	stdoutScan := bufio.NewScanner(stdout)
	stderrScan := bufio.NewScanner(stderr)
	go func() {
		for stdoutScan.Scan() {
			line := stdoutScan.Text()
			ch <- line
		}
	}()

	go func() {
		for stderrScan.Scan() {
			line := stderrScan.Text()
			ch <- line
		}
	}()

	var err error
	go func() {
		err = cmd.Wait()
		if err != nil && !strings.Contains(err.Error(), "exit status") {
			log.Println("wait: ", err.Error())
		}
		close(ch)
	}()
	for line := range ch {
		log.Println(line)
	}
	return err
}
