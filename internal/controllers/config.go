package controllers

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port string
	dir  string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}

func ConfigureAndStart() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dir := flag.String("dir", "./ui/static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	config := config{
		port: *addr,
		dir:  *dir,
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		config:   config,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:         app.config.port,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	infoLog.Printf("%s", config.dir)
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
