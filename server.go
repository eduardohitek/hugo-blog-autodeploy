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

//PushWebHook Struct que representa uma chamda do GithubWebhook
type PushWebHook struct {
	Action     string `json:"action" bson:"action"`
	CheckSuite struct {
		CreatedAt  string `json:"created_at" bson:"created_at"`
		HeadBranch string `json:"head_branch" bson:"head_branch"`
		HeadCommit struct {
			Author struct {
				Email string `json:"email" bson:"email"`
				Name  string `json:"name" bson:"name"`
			} `json:"author" bson:"author"`
			ID        string `json:"id" bson:"id"`
			Message   string `json:"message" bson:"message"`
			Timestamp string `json:"timestamp" bson:"timestamp"`
			TreeID    string `json:"tree_id" bson:"tree_id"`
		} `json:"head_commit" bson:"head_commit"`
		HeadSha              string `json:"head_sha" bson:"head_sha"`
		ID                   int64  `json:"id" bson:"id"`
		LatestCheckRunsCount int64  `json:"latest_check_runs_count" bson:"latest_check_runs_count"`
		NodeID               string `json:"node_id" bson:"node_id"`
	} `json:"check_suite" bson:"check_suite"`
}

// AutoDeploy Handler for the Github Webhook
func AutoDeploy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			log.Fatal(err)
		}
		request := PushWebHook{}
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
