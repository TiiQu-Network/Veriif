package main

type App struct {}

// Initialisation
func (a App) Init() {
	printTitle()
	err := initContract()
	if err != nil {
		printError(err)
	}
}

func (a App) Run() {
	for {
		a.Verify()
	}
}

func (a App) Verify() {
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