package handlers

import (
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	filePath := "static/app.html"
	http.ServeFile(w, r, filePath)
}
