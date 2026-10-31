// mimi is a nice cat
// github : https://github.com/MohssineX
// Copyright (C) 2026 Mohssine <https://github.com/MohssineX>
// SPDX-License-Identifier: Apache-2.0

// AppName : mcls
// Version : 1.0.1

package main

// import libraries
import "os"

const clearScreenSeq = "\033[H\033[2J\033[3J"

// Clear the terminal function
func clearScreen() {
	_, err := os.Stdout.WriteString(clearScreenSeq)

	// Report errors on stderr and exit with code 1
	if err != nil {
		_, _ = os.Stderr.WriteString("mcls: ERROR: " + err.Error())
		os.Exit(1)
	}
}

func main() {
	// Clear the terminal
	clearScreen()
}
