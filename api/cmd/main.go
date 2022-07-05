package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"mini-api/internal/driver"
	"mini-api/internal/models"

	"github.com/go-chi/cors"
)

const version = "1.0.0"

type config struct {
	port   int
	env    string
	docs   string
	secret string
	cors   []string
	db     struct {
		dsn string
	}
}

type application struct {
	conf    config
	info    *log.Logger
	err     *log.Logger
	version string
	cors    cors.Options
	DB      models.DB
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.conf.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.info.Printf("Running on :%d (%s)", app.conf.port, app.conf.env)
	return srv.ListenAndServe()
}

func main() {
	c := loadConfig()
	cfg := config{
		port:   c.Port,
		env:    c.Mode,
		docs:   c.Docs,
		secret: c.Secret,
		cors:   c.Cors,
	}

	cfg.db.dsn = c.DSN
	infoLog := log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	conn, err := driver.OpenDB(cfg.db.dsn)

	if err != nil {
		errLog.Fatal(err)
	}

	defer conn.Close()

	app := &application{
		conf: cfg,
		info: infoLog,
		err:  errLog,
		cors: cors.Options{
			AllowedOrigins:   cfg.cors,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Token"},
			AllowCredentials: false,
			MaxAge:           300,
		},
		version: version,
		DB: models.DB{
			Database: conn,
		},
	}

	if err = app.serve(); err != nil {
		log.Fatal(err)
	}
}
