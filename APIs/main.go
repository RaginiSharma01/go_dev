package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"apiwithparameters/posts"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method { // specified method ( weather its a get , put, patch , delete)

		case http.MethodGet:
			id := strings.TrimPrefix(r.URL.Path, "/users/")
			if id == "" {
				http.Error(w, "invalid user id", http.StatusBadRequest)
				return
			}

			userID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "invalid user id", http.StatusBadRequest)
				return
			}

			active := r.URL.Query().Get("active")
			if active != "" && active != "true" && active != "false" {
				http.Error(w, "invalid active value", http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"user_id": userID,
				"active":  active,
			})

		case http.MethodPost:
			var user User

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}

			if user.Name == "" || user.Email == "" {
				http.Error(w, "missing fields", http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(user)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	posts.RegisterRoute()
	http.ListenAndServe(":8080", nil)
}
