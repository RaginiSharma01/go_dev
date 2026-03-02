package posts

import "net/http"

func posts() {
	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8080", nil)
}
