package handlers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/TiiQu-Network/Veriif/app/interfaces"
	"github.com/TiiQu-Network/Veriif/app/models"
	"github.com/TiiQu-Network/Veriif/app/templates"
	"github.com/TiiQu-Network/Veriif/app/verify"
)

type Handler struct{
	app interfaces.Application
}

func New(a interfaces.Application) Handler {
	return Handler{a}
}

func (h Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles(templates.GetTemplates("base", "upload"))
		h.app.ErrLogger().Log(err, "Parsing template files")
		h.app.ErrLogger().Log(t.ExecuteTemplate(w, "base", nil), "Execute Welcome template")
		break

	case http.MethodPost:
		cert, err := getCert(r)
		h.app.ErrLogger().Log(err, "Getting cert data")

		certPacket, err := verify.ExtractCert(cert)

		results, err := verify.Verify(certPacket, h.app.EthCaller())

		data := map[string]interface{}{}
		data["results"] = results
		data["cert_data"] = certPacket.Data
		data["fail"] = err

		println("")
		t, err := template.ParseFiles(templates.GetTemplates("base", "result"))

		h.app.ErrLogger().Log(err, "Parsing template files")
		h.app.ErrLogger().Log(t.ExecuteTemplate(w, "base", data), "Execute Welcome template")
		break
	}

}

func (h Handler) API(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:

		decoder := json.NewDecoder(r.Body)

		certPacket := models.CertPacket{}
		err := decoder.Decode(&certPacket)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(`{"error":"Invalid request payload : `+err.Error()+`""}`))
			return
		}

		// TODO add Validation of the incoming data

		results, err := verify.Verify(certPacket, h.app.EthCaller())

		data := map[string]interface{}{}
		data["results"] = results
		data["cert_data"] = certPacket.Data
		if err !=nil {
			data["fail"] = err.Error()
		} else {
			data["fail"] = err
		}

		response, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(response)
	}

}

func (h Handler) APILite(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:

		decoder := json.NewDecoder(r.Body)

		certPacketLite := models.CertPacketLite{}
		err := decoder.Decode(&certPacketLite)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(`{"error":"Invalid request payload : `+err.Error()+`""}`))
			return
		}

		// TODO add Validation of the incoming data

		// Convert the CPL to a standard CP
		certPacket, err := certPacketLite.ToCertPacket()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(`{"error":"Invalid request payload : `+err.Error()+`""}`))
			return
		}

		results, err := verify.Verify(*certPacket, h.app.EthCaller())

		data := map[string]interface{}{}
		data["results"] = results
		data["cert_data"] = certPacket.Data
		if err !=nil {
			data["fail"] = err.Error()
		} else {
			data["fail"] = err
		}

		response, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(response)
	}

}

func getCert(r *http.Request) ([]byte, error) {

	// Parse the form data with files
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, err
	}

	// Open the file data
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, err
	}

	// Close the file
	err = file.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
