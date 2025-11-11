package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"example.com/elastic-go/internal/es"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	es.Init()

	http.HandleFunc("/search", SearchHandler)
	log.Println("server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type SearchResult struct {
	Hits struct {
		Hits []struct {
			ID     string          `json:"_id"`
			Score  float64         `json:"_score"`
			Source json.RawMessage `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "missing q", http.StatusBadRequest)
		return
	}

	query := map[string]any{
		"query": map[string]any{
			"multi_match": map[string]any{
				"query":     q,
				"fields":    []string{"title", "slug"},
				"fuzziness": "AUTO",
			},
		},
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(query)

	res, err := es.Client.Search(
		es.Client.Search.WithIndex("metadata"),
		es.Client.Search.WithBody(&buf),
		es.Client.Search.WithSize(10),
	)
	if err != nil {
		http.Error(w, "SearchHandler es.Client.Search: "+err.Error(), 500)
		return
	}
	defer res.Body.Close()

	var result SearchResult
	json.NewDecoder(res.Body).Decode(&result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Hits.Hits)
}
