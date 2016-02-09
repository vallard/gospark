package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vallard/gospark/models"
)

func GitCommit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	decoder := json.NewDecoder(r.Body)
	var c models.CommitMessage
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	log.Println(c.Repo.Name)

	hah, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "This is our body: %s", hah)
	log.Printf("%s", hah)

}

func GitHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	log.Println("Serving /hello")
	fmt.Fprintln(w, "Hello, is it me you're looking for?")

}
