package main

import (
	"bufio"
	"github.com/fatih/color"
	"os"
	"strings"
)

// TODO add proper error handling / logging
func main() {
	color.Set(color.FgGreen, color.Bold)
	println(`____   ____           .__.__  _____  `)
	println(`\   \ /   /___________|__|__|/ ____\ `)
	println(` \   Y   // __ \_  __ \  |  \   __\  `)
	println(`  \     /\  ___/|  | \/  |  ||  |    `)
	println(`   \___/  \_____>__|  |__|__||__|    `)
	println("")
	color.Unset()

	println("application - Veriif (c) 2019 TiiQu Ltd")
	println("version - 0.1.1")
	println("")

	err := initContract()

	// Get file name from input
	for {
		if err != nil {
			color.Set(color.FgRed)
			println("- ERROR   - " + err.Error())
			color.Unset()
		}

		reader := bufio.NewReader(os.Stdin)
		print("Enter certificate filename: ")
		color.Set(color.FgHiYellow)
		filename, _ := reader.ReadString('\n')

		// TODO check for "exit", then exit application

		filename = strings.TrimSpace(filename)
		println("- Processing : " + filename)
		color.Unset()

		err := verify(filename)
		if err != nil {
			color.Set(color.FgRed)
			println("- FAIL    - " + err.Error())
			color.Unset()

			color.Set(color.FgRed, color.Bold)
			println("|--------------------------------------------------------------|")
			println("| {x}  NOT VERIFIED - The certificate failed verification  {x} |")
			println("|--------------------------------------------------------------|")
			println("")
			color.Unset()

		} else {
			color.Set(color.FgGreen, color.Bold)
			println("|--------------------------------------------------------------|")
			println("| {+}     VERIFIED - The certificate has been verified     {+} |")
			println("|--------------------------------------------------------------|")
			println("")
			color.Unset()
		}
	}
}

