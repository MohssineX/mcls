// mimi is a nice cat
// github : https://github.com/MohssineX
// Copyright (C) 2026 Mohssine <https://github.com/MohssineX>
// SPDX-License-Identifier: Apache-2.0

package main

// import libraries
import "os"

func main() {

	const ClearScreen = "\033[H\033[2J\033[3J"

	// Clear the terminal
	_, _ = os.Stdout.WriteString(ClearScreen)

}
