package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // New import
	"github.com/sparkycj328/Snippetbox/pkg/models/mysql"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	// Define a command-line flag for the MYSQL DSN string
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MYSQL database")
	// Define a default value command-line flag and parse it
	addr := flag.String("addr", ":4000", "HTTP network address")

	//Parse any command-line flags during application launch
	flag.Parse()

	// Create two new loggers to be used for informational messages and error messages respectively
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create a database connection
	db, err := openDB(*dsn)
	defer db.Close()

	// Initialize a new template cache
	templateCache, err := newTemplateCache("../../ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}
	// Initialize application struct
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// Create a new http.Server struct and pass it the new errorLog
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// start a new web server which listens on the passed address and serves the servemux created above
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB wraps sql.Open() and returns a sql.DB connectoon pool for the given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
