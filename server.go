package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//PushWebHook Struct que representa uma chamda do GithubWebhook
type PushWebHook struct {
	HeadCommit struct {
		Author struct {
			Username string `json:"username" bson:"username"`
		} `json:"author" bson:"author"`
		Message string `json:"message" bson:"message"`
	} `json:"head_commit" bson:"head_commit"`
	Ref string `json:"ref" bson:"ref"`
}

// AutoDeploy Handler for the Github Webhook
func AutoDeploy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	signature := r.Header.Get("X-Hub-Signature")
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			log.Fatal(err)
		}
		a := HashPayload("H3lpM3Str4ng3r", body)
		request := PushWebHook{}
		log.Println(string(body))
		json.Unmarshal(body, &request)
		log.Println(request)
		log.Println(a, signature)
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

// HashPayload função para gerar o hash de verificação do Github Webhook
func HashPayload(secret string, playloadBody []byte) string {
	hm := hmac.New(sha1.New, []byte(secret))
	hm.Write(playloadBody)
	sum := hm.Sum(nil)
	return fmt.Sprintf("%x", sum)
}
