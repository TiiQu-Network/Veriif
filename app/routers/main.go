package routers

import (
	"net/http"

	"github.com/TiiQu-Network/Veriif/app/handlers"
	"github.com/TiiQu-Network/Veriif/app/interfaces"
	"github.com/TiiQu-Network/Veriif/app/paths"
)

func InitHttp(a interfaces.Application) {
	// File routes
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(paths.CssDir))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(paths.JsDir))))

	// Get a new handler
	h := handlers.New(a)

	// Application routes
	http.HandleFunc("/", h.Welcome)
	http.HandleFunc("/qr", h.QR)
	http.HandleFunc("/api", h.API)
	http.HandleFunc("/api/lite", h.APILite)
}
