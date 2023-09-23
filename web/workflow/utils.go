package main

import (
	"fmt"
	"os"
	"os/exec"
)

func BuildTagName(pluginName string, version string) string {
	return fmt.Sprintf("%s|v%s", pluginName, version)
}

func RunCommand(c string) {
	cmd := exec.Command("bash", "-c", c)
	cmd.Dir = RootPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	//if err := cmd.Run(); err != nil {
	//	panic(err) // When nothing to commit, it will panic
	//}
}
