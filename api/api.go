package api

import (
	"net/http"
)

const Version = "v1"

type api interface {
	Version() string
	Default(w http.ResponseWriter, r *http.Request)
	Hello(w http.ResponseWriter, r *http.Request)
	StockReport(w http.ResponseWriter, r *http.Request)
}
