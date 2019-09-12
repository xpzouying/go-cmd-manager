package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Task struct {
	ID  string
	cmd *exec.Cmd
}

func NewTask() *Task {
	cmd := exec.Command("python3", "../python-worker/main.py")
	cmd.Stdout = os.Stdout

	taskid := uuid.NewV4().String()

	return &Task{
		ID:  taskid,
		cmd: cmd,
	}
}

func (t *Task) Run() error {
	return t.cmd.Run()
}

func (t *Task) Stop() error {
	log.Printf("task(%s) will stop", t.ID)

	return t.cmd.Process.Kill()
}

func main() {
	t1 := NewTask()

	go func() {
		if err := t1.Run(); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second * 3)
	t1.Stop()
}
