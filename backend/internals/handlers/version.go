package handlers

import (
	"encoding/json"
	"net/http"
	"pulley.com/shakesearch/pkg"
)

// Version returns the API version based on the constant version.
func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", pkg.JSONContentType)
	response := map[string]string{}
	response["version"] = pkg.APIVersion
	jsonContent, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}
