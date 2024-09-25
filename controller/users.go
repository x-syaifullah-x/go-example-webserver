package controller

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

func Users() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data = []student{
			{"A_001", "ethan", 21},
			{"A_002", "wick", 22},
			{"A_003", "bourne", 23},
		}
		var result, _ = json.Marshal(data)
		w.Write(result)
	}
}
