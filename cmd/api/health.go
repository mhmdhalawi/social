package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request)  {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"available","environment":"production","version":"1.0.0"}`))
	}