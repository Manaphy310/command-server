package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type CommandList struct {
	Commands []string `json:"commands"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("commands.json")
	if err != nil {
		http.Error(w, "Could not read commands file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var cmdList CommandList
	if err := json.NewDecoder(file).Decode(&cmdList); err != nil {
		http.Error(w, "Invalid JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applicatin/json")
	json.NewEncoder(w).Encode(cmdList)
}

func main() {
	http.HandleFunc("/commands", handler)
	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
