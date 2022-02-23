package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// Define a default value command-line flag and parse it
	addr := flag.String("addr", ":4000", "HTTP network address")

	//Parse any command-line flags during application launch
	flag.Parse()

	// Create two new loggers to be used for informational messages and error messages respectively
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Create a new http.Server struct and pass it the new errorLog
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// start a new web server which listens on the passed address and serves the servemux created above
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
