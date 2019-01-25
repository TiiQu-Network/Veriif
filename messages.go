package main

import (
	"bufio"
	"github.com/fatih/color"
	"os"
	"strings"
)

func printTitle() {
	color.Set(color.FgGreen, color.Bold)
	println(`____   ____           .__.__  _____  `)
	println(`\   \ /   /___________|__|__|/ ____\ `)
	println(` \   Y   // __ \_  __ \  |  \   __\  `)
	println(`  \     /\  ___/|  | \/  |  ||  |    `)
	println(`   \___/  \_____>__|  |__|__||__|    `)
	println("")
	color.Unset()

	println("application - Veriif (c) 2019 TiiQu Ltd")
	println("version - 0.2.0")
	println("")
}

func printError(err error) {
	color.Set(color.FgRed)
	println("- ERROR   - " + err.Error())
	color.Unset()
}

func printSuccess(message string) {
	color.Set(color.FgGreen)
	println("- SUCCESS - " + message)
	color.Unset()
}

func printProcessing(message string) {
	color.Set(color.FgHiYellow)
	println("- Processing : " + message)
	color.Unset()
}

func printPrompt() {
	print("Enter certificate filename: ")
}

func getInput() string{
	color.Set(color.FgHiYellow)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	color.Unset()

	// TODO check for "exit", then exit application

	return strings.TrimSpace(input)
}

func printVerifyFail(err error) {
	color.Set(color.FgRed)
	println("- FAIL    - " + err.Error())
	color.Unset()

	color.Set(color.FgRed, color.Bold)
	println("|--------------------------------------------------------------|")
	println("| {x}  NOT VERIFIED - The certificate failed verification  {x} |")
	println("|--------------------------------------------------------------|")
	println("")
	color.Unset()
}

func printVerifySuccess() {
	color.Set(color.FgGreen, color.Bold)
	println("|--------------------------------------------------------------|")
	println("| {+}     VERIFIED - The certificate has been verified     {+} |")
	println("|--------------------------------------------------------------|")
	println("")
	color.Unset()
}