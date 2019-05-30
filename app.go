package main

import (
	"flag"
	"github.com/pkg/browser"
	"io/ioutil"
	"log"
	"net/http"
)

type App struct {
	ErrLogger *log.Logger
	Mode *uint
}

func NewApp() App {
	return App{
		ErrLogger: initErrLogger(),
		Mode: getParams(),
	}
}

// Application Initialisation
func (a App) Init() {
	switch *a.Mode {
	case 1:
		a.initCLI()
		break
	case 2:
		fallthrough
	case 3:
		a.initBrowser()
		break
	}
}

func getParams() *uint {
	mode := flag.Uint("mode", 2, "System Mode :\n - 1 for CLI mode\n - 2 for browser UI mode\n - 3 for Docker mode")
	flag.Parse()
	return mode
}

func (a App) initCLI() {
	printTitle()
	err := initContract()
	if err != nil {
		printError(err)
	}
}

func (a App) initBrowser() {
	a.initRoutes()
	a.initLoadingText()
	err := initContract()
	if err != nil {
		printError(err)
	}
}

func (a App) initRoutes() {
	// File routes
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(cssDir))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(jsDir))))

	// Application routes
	http.HandleFunc("/", a.Welcome)
	http.HandleFunc("/api", a.API)
	http.HandleFunc("/api/lite", a.APILite)
}

func (a App) initLoadingText() {
	printTitle()
	println("Veriif running ...")
	println("Opening http://localhost:8282 ...")
	println("Please wait... Loading...")
}

// Application Running

func (a App) Run() {
	switch *a.Mode {
	case 1:
		a.runCLI()
		break
	case 2:
		a.runBrowser()
		break
	case 3:
		a.runDocker()
		break
	}
}

func (a App) runCLI() {
	for {
		a.verify()
	}
}

func (a App) runBrowser() {
	a.Log(browser.OpenURL("http://localhost:8282"), "Opening browser")
	a.LogFatal(http.ListenAndServe(":8282", nil), "While serving")
}

func (a App) runDocker() {
	a.LogFatal(http.ListenAndServe(":80", nil), "While serving")
}

func (a App) verify() {
	printPrompt()
	filename := getInput()
	printProcessing(filename)

	// Get the cert file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		printError(err)
	}

	certPacket, err := extractCert(content)
	_, err = verify(certPacket)
	if err != nil {
		printVerifyFail(err)
	} else {
		printVerifySuccess()
	}
}