package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mailtomainak/snippetbox/pkg/models/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippetModel  *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL Data Source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Println(err)
	}
	defer db.Close()
	snippetModel := &mysql.SnippetModel{
		DB: db,
	}
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{errorLog, infoLog, snippetModel, templateCache}

	infoLog.Printf("Starting server on %s", *addr)

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

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
