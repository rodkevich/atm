package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	log.Println("Hello World")
	_ = os.Setenv("DATABASE_URL", "postgresql://admin:password@db:5432/todos?sslmode=disable")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":11080", nil))
}
