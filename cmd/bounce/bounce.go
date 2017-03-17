package main

import (
	"log"
	"net/http"
	"net/url"
)

func redirectHandler(w http.ResponseWriter, req *http.Request) {
	baseURL, err := url.Parse("https://" + req.Host + "/")
	if err != nil {
		log.Println(err)
		http.Error(w, "", 400)
		return
	}

	absURL := req.URL.ResolveReference(baseURL)
	http.Redirect(w, req, absURL.String(), http.StatusMovedPermanently)
}

func main() {
	log.Println("Binding :http")
	log.Panic(http.ListenAndServe(":http", http.HandlerFunc(redirectHandler)))
}
