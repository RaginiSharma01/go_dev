package posts

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Post struct {
	Id   int    `json:"id"`
	User string `json:"user"`
	Post string `json:"post"`
}

func RegisterRoute() {

	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			id := strings.TrimPrefix(r.URL.Path, "/posts/")
			if id == "" {
				http.Error(w, "invalid id", http.StatusBadRequest)
				return
			}

			postID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "invalid id", http.StatusBadRequest)
				return
			}
			active := r.URL.Query().Get("active")
			if active != "" && active != "true" && active != "false" {
				http.Error(w, "invalid active id", http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Post{
				Id:   postID,
				User: "test_user",
				Post: "tesr_post",
			})

		}

	})

}
