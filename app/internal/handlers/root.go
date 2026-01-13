package handlers

import (
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	filePath := "templates/root.html"
	http.ServeFile(w, r, filePath)
}
