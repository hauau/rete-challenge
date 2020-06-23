package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":8080"
	logfile, logerr := os.OpenFile(
		"file.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)

	if logerr != nil {
		panic(logerr)
	}
	defer logfile.Close()

	logger := log.New(logfile, "", 0)
	logHandler := newLogHandler(logger)

	http.HandleFunc("/", logHandler)

	log.Fatal(http.ListenAndServe(addr, nil))
}

func newLogHandler(logger *log.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if "POST" == r.Method {
			if err := r.ParseForm(); err != nil {
				panic(err)
			}

			for formKey := range r.PostForm {
				logger.Println(formKey)
			}
		}
	}
}
