package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		tee := io.MultiWriter(res, log.Writer())
		fmt.Fprintln(tee, "Greetings!")
	})

	var (
		yaaargs int
		mu      sync.RWMutex
		stop    = make(chan struct{})
	)
	mux.HandleFunc("/ahoy", func(res http.ResponseWriter, req *http.Request) {
		tee := io.MultiWriter(res, log.Writer())
		fmt.Fprintln(tee, "YAAARG!")
		mu.Lock()
		defer mu.Unlock()
		yaaargs++
		if yaaargs > 3 {
			close(stop)
		}
	})

	srv := &http.Server{Addr: ":9000", Handler: mux}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-stop
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println(err)
		}
		return
	}()

	log.Println(srv.ListenAndServe())

	wg.Wait()
}
