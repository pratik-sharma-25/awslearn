package main

import "net/http"

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/greet", greetUser)

	http.ListenAndServe(":8040", router)
}

func greetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// CI/CD pipeline we are using github action
// actions -> automated task which run on some events
