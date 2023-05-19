package handlers

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"go-intermediate/pkg/okta"
)

type commonData struct {
	IsAuthenticated bool
}

var (
	templates *template.Template
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	templates = template.Must(template.ParseGlob(wd + "/web/templates/*"))
}

func GetCommonData(request *http.Request) commonData {
	data := commonData{
		IsAuthenticated: okta.IsAuthenticated(request),
	}
	return data
}
