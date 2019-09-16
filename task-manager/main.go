package main

import (
	"bufio"
	"log"
	"os/exec"
)

func main() {
	cmdName := "python ../python-worker/main.py"
	cmd := exec.Command("sh", "-c", cmdName)
	stdout, _ := cmd.StdoutPipe()

	cmd.Start()

	scan := bufio.NewScanner(stdout)
	for scan.Scan() {
		text := scan.Text()
		log.Printf("get: %s\n", text)
	}

	cmd.Wait()
	log.Println("golang finish/end")
}
