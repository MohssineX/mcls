// mimi is a nice cat
// github : https://github.com/MohssineX
// Copyright (C) 2026 Mohssine <https://github.com/MohssineX>
// SPDX-License-Identifier: Apache-2.0

package main

// import libraries
import (
	"os"
	"os/exec"
	"runtime"
)

func main() {

	const ClearScreen = "\033[H\033[2J\033[3J"

	// Enable ANSI escape codes on Windows (not needed on Linux/Mac)
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	// Clear the terminal
	os.Stdout.WriteString(ClearScreen)

}
