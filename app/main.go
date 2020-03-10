package app

import (
	"flag"
	"io/ioutil"
	"net/http"

	"github.com/pkg/browser"

	"github.com/TiiQu-Network/Veriif/app/errors"
	"github.com/TiiQu-Network/Veriif/app/ethereum"
	"github.com/TiiQu-Network/Veriif/app/messages"
	"github.com/TiiQu-Network/Veriif/app/routers"
	"github.com/TiiQu-Network/Veriif/app/verify"
)

type App struct {
	ErrorLogger errors.Logger
	Mode uint
	Caller ethereum.Caller
}

func NewApp() App {
	return App{
		ErrorLogger: errors.InitLogger(),
		Mode: getParams(),
	}
}

func (a *App) ErrLogger() errors.Logger {
	return a.ErrorLogger
}

func (a *App) EthCaller() ethereum.Caller {
	return a.Caller
}

// Application Initialisation
func (a *App) Init() {
	c, err := ethereum.Init()
	if err != nil {
		messages.PrintError(err)
	}
	a.Caller = c

	switch a.Mode {
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

func getParams() uint {
	mode := flag.Uint("mode", 2, "System Mode :\n - 1 for CLI mode\n - 2 for browser UI mode\n - 3 for Docker mode")
	flag.Parse()
	return *mode
}

func (a *App) initCLI() {
	messages.PrintTitle()
}

func (a *App) initBrowser() {
	routers.InitHttp(a)
	a.initLoadingText()
}

func (a App) initLoadingText() {
	messages.PrintTitle()
	println("Veriif running ...")
	println("Opening http://localhost:8282 ...")
	println("Please wait... Loading...")
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
	a.ErrLogger().Log(browser.OpenURL("http://localhost:8282"), "Opening browser")
	a.ErrLogger().LogFatal(http.ListenAndServe(":8282", nil), "While serving")
}

func (a App) runDocker() {
	a.ErrLogger().LogFatal(http.ListenAndServe(":80", nil), "While serving")
}

func (a App) verify() {
	messages.PrintPrompt()
	filename := messages.GetInput()
	messages.PrintProcessing(filename)

	// Get the cert file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		messages.PrintError(err)
	}

	certPacket, err := verify.ExtractCert(content)
	_, err = verify.Verify(certPacket, a.Caller)
	if err != nil {
		messages.PrintVerifyFail(err)
	} else {
		messages.PrintVerifySuccess()
	}
}