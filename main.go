package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mcgtrt/gtx/handler"
	"github.com/phuslu/log"
)

var (
	r          *mux.Router
	listenAddr string
)

func main() {
	if log.IsTerminal(os.Stderr.Fd()) {
		log.DefaultLogger = log.Logger{
			TimeFormat: "15:04:05",
			Caller:     1,
			Writer: &log.ConsoleWriter{
				ColorOutput:    true,
				QuoteString:    true,
				EndWithMessage: true,
			},
		}
	}

	log.Info().
		Str("port", listenAddr).
		Msg("HTTP SendiForce Server Running")

	log.Error().
		Err(http.ListenAndServe("localhost:"+listenAddr, r)).
		Msg("http server error")
}

func init() {
	r = handler.NewRouter(&handler.Handler{})
	listenAddr = os.Getenv("HTTP_SERVER_PORT")
}
