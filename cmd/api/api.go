package main

import "net/http"

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) run() error {

	server := &http.Server{
		Addr: a.config.addr,
	}

	return server.ListenAndServe()

}
