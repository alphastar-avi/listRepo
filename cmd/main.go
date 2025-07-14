package main

import (
	"encoding/json"
	"net/http"
	"listRepo/internals/gitapi"
	"fmt"
)

type RequestBody struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

type ResponseBody struct {
	Repos []string `json:"repos"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		var req RequestBody
		json.NewDecoder(r.Body).Decode(&req)
		repos := gitapi.GetRepoNames(req.Username, req.Token)
		resp := ResponseBody{Repos: repos}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
