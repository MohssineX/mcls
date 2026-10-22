// mimi is a nice cat
// github : https://github.com/MohssineX
// Copyright (C) 2026 Mohssine <https://github.com/MohssineX>
// SPDX-License-Identifier: Apache-2.0

package main

// import libraries
import "os"

const clearScreenSeq = "\033[H\033[2J\033[3J"

// Clear the terminal function
func clearScreen() {
	_, err := os.Stdout.WriteString(clearScreenSeq)

	if err != nil {
		os.Exit(1)
	}
}

func main() {
	// Clear the terminal
	clearScreen()
}
