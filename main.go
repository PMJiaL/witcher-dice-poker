package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/Depermitto/witcher-dice-poker/docs"
	"github.com/Depermitto/witcher-dice-poker/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//	@title			Witcher Dice Poker API
//	@version		1.0
//	@description	Webserver serving a complete implementation of Witcher 1 (2007) dice poker mini-game.

//	@contact.name	Piotr (Depermitto) Jabłoński
//	@contact.email	penciller@disroot.org

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// @BasePath
func main() {
	port := flag.String("port", "2007", "Port to listen on")
	flag.Parse()

	logger := log.New(os.Stdout, log.Prefix(), log.Flags())
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	{
		if os.Getenv("APP_ENV") != "production" {
			addr := fmt.Sprintf("http://localhost:%v/swagger/", *port)
			r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(addr+"doc.json")))
			logger.Printf("Swagger UI available at %v\n", addr+"index.html")
		}
		r.Get("/hands/random", handler.RandomHand)
		r.Post("/hands/switch", handler.UpdateHand)
		r.Post("/hands/eval", handler.EvaluateHand)
	}

	srv := &http.Server{Addr: "0.0.0.0:" + *port, Handler: r}
	go func() {
		logger.Printf("Listening on port %v\n", *port)
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
