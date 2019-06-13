package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// AutoDeploy Handler for the Github Webhook
func AutoDeploy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			log.Fatal(err)
		}
		request := Request{}
		log.Println(string(body))
		json.Unmarshal(body, &request)
		log.Println(request)
		w.Write([]byte("Gasto adicionado com Sucesso"))
	}
}

func main() {
	log.Println(os.Getenv("GITHUB_SECRET"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", AutoDeploy)

	srv := &http.Server{
		Addr:         "127.0.0.1:8787",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	log.Println("starting the server")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to star: %v", err)
	}
}

// Request bla bla bla
type Request struct {
	GithubSecret string `json:"ghsecret"`
	Teste        int    `json:"teste"`
}
