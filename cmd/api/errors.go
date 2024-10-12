package main

import "net/http"

func (api *application) internalServerError(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusInternalServerError, err.Error())
}

func (api *application) invalidPayloadError(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusBadRequest, err.Error())
}
