package api

import (
	"net/http"
	"time"

	"github.com/karuppaiah/docxtopdf/pdfgenerate/usecase"
)

func Collector(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the delay.
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check to make sure the delay is anywhere from 1 to 10 seconds.
	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	// Now, we retrieve the person's name from the request.
	name := r.FormValue("name")

	// Just do a quick bit of sanity checking to make sure the client actually provided us with a name.
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}
	// Now, we retrieve the person's name from the request.
	filename := r.FormValue("filename")
	// Just do a quick bit of sanity checking to make sure the client actually provided us with a name.
	if filename == "" {
		http.Error(w, "You must specify a docx filename.", http.StatusBadRequest)
		return
	}
	usecase.CollectorJob(name, delay, filename)
	// And let the user know their work request was created.
	w.WriteHeader(http.StatusCreated)
	return
}
