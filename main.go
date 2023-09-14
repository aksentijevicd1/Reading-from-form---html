package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aksentijevicd1/reading-from-form-go/handlers"
)

func main() {

	l := log.New(os.Stdout, "form-apply", log.LstdFlags)
	oh := handlers.NewOpinions(l)

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/form", oh.AddOpinion)
	http.HandleFunc("/opinions", oh.GetOpinions)

	s := http.Server{
		Addr:         ":9090",
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error while starting server: %s", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown!", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
