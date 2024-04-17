package client

import (
	"log"
	"os/exec"
)

func Start() {
	cmd := exec.Command("go", "run", "../server/controller")
	log.Printf("Running worker node and detaching!")
	err := cmd.Start()

	if err != nil {
		log.Fatal("cmd.Start failed: ", err)
	}

	err = cmd.Process.Release()
	if err != nil {
		log.Fatal("cmd.Process.Release failed: ", err)
	}
}

func Stop() {
	log.Printf("Stopping worker node!")
	cmd := exec.Command("fuser", "-k", "3000/tcp")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Stop server failed: ", err)
	}
}
