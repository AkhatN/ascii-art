package art

import (
	"net/http"
	"strconv"
)

//MethodNotAllowed replies to the request with an HTTP 405 Method Not Allowed error.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
}

//InternalServerError replies to the request with an HTTP 500 Internal Server Error error.
func InternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
}

//BadRequest replies to the request with an HTTP 404 Bad Request Error error.
func BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Bad Request", http.StatusBadRequest)
}

//SendFileToClient ...
func SendFileToClient(w http.ResponseWriter, r *http.Request, s string) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Length", strconv.Itoa(len(s)))
	_, er := w.Write([]byte(s))
	if er != nil {
		return
	}
}
