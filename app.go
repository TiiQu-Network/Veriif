package main

import "flag"

// TODO add proper error handling / logging
type App struct {
	Mode uint
}

// Application Initialisation
func (a App) Init() {
	a.getParams()

	switch a.Mode {
	case 1:
		a.initCLI()
		break
	case 2:
		a.initBrowser()
		break
	}
}

func (a App) getParams() {
	mode := flag.Uint("mode", 1, "System Mode :\n - 1 for CLI mode\n - 2 for browser UI mode\n -")
	flag.Parse()
	a.Mode = *mode
}

func (a App) initCLI() {
	printTitle()
	err := initContract()
	if err != nil {
		printError(err)
	}
}

func (a App) initBrowser() {

}

// Application Running

func (a App) Run() {
	switch a.Mode {
	case 1:
		a.runCLI()
		break
	case 2:
		a.runBrowser()
		break
	}
}

func (a App) runCLI() {
	for {
		a.verify()
	}
}

func (a App) runBrowser() {

}

func (a App) verify() {
	printPrompt()
	filename := getInput()
	printProcessing(filename)

	err := verify(filename)
	if err != nil {
		printVerifyFail(err)
	} else {
		printVerifySuccess()
	}
}