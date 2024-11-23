package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func genrateShortUrl(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))

	fmt.Println(hasher)

	data := hasher.Sum(nil)
	fmt.Println("hasher data", data)

	hash := hex.EncodeToString(data)
	fmt.Println("Ecoding string", hash)
	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := genrateShortUrl(originalURL)
	id := shortURL

	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}

	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]

	if !ok {
		return URL{}, errors.New("URL not found")
	}

	return url, nil
}

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	shortURL_ := genrateShortUrl(data.URL)

	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func redirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]

	url, err := getURL(id)

	if err != nil {
		http.Error(w, "Invalid url", http.StatusBadRequest)
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)

}

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// OriginalURL := "http://localhost:3000/"
	// genrateShortUrl(OriginalURL)

	http.HandleFunc("/", ShortUrlHandler)
	http.HandleFunc("/redirect", redirectUrlHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Err o starting new server", err)
	}
}
