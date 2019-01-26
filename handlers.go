package main

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
)

func (a App) Welcome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, err := template.ParseFiles(GetTemplates("base", "upload"))
		a.Log(err, "Parsing template files")
		a.Log(t.ExecuteTemplate(w, "base", nil), "Execute Welcome template")
		break

	case http.MethodPost:
		cert, err := getCert(r)
		a.Log(err, "Getting cert data")

		certPacket, results, err := verify(cert)

		data := map[string]interface{}{}
		data["results"] = results
		data["cert_data"] = certPacket.Data
		data["fail"] = err

		println("")
		t, err := template.ParseFiles(GetTemplates("base", "result"))

		a.Log(err, "Parsing template files")
		a.Log(t.ExecuteTemplate(w, "base", data), "Execute Welcome template")
		break
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
