package cmdexecutions

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ExecuteCLICommand() {
	out, err := exec.Command("ifconfig", "en0").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}

func ExecuteShellScript() {
	path, _ := os.Getwd()
	fileName := "shell_script.sh"

	filePath := filepath.Join(path, "script", fileName)

	out, err := exec.Command("/bin/sh", filePath).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}
