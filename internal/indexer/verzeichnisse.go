// Package indexer is responsible for indexing the tables from the database
package indexer

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"

	"example.com/elastic-go/internal/db"
	"example.com/elastic-go/internal/es"
)

func IndexVerzeichnisse() {
	maindataRows := fetchMetadataRows()
	defer maindataRows.Close()

	metadataBulkBody := buildMetadataBulkBody(maindataRows)

	sendBulkToES(metadataBulkBody)

	log.Println("metadata indexed")
}

func fetchMetadataRows() *sql.Rows {
	rows, err := db.DB.Query(`SELECT uid, pid, title, area, number, 
		short_title, sub_area, slug FROM tx_spiaport_domain_model_metadata`)
	if err != nil {
		log.Fatalf("IndexMetadata db.DB.Query: %v", err)
	}
	return rows
}

func buildMetadataBulkBody(rows *sql.Rows) []byte {
	var buf bytes.Buffer

	for rows.Next() {
		var uid, pid int
		var title, area, number, shortTitle, subArea, slug string
		rows.Scan(&uid, &pid, &title, &area, &number, &shortTitle, &subArea, &slug)

		meta := map[string]any{"index": map[string]any{
			"_index": "metadata",
			"_id":    uid,
		}}
		doc := map[string]any{
			"uid":         uid,
			"pid":         pid,
			"title":       title,
			"area":        area,
			"number":      number,
			"short_title": shortTitle,
			"sub_area":    subArea,
			"slug":        slug,
		}

		m, _ := json.Marshal(meta)
		d, _ := json.Marshal(doc)

		buf.Write(m)
		buf.WriteByte('\n')
		buf.Write(d)
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

func sendBulkToES(body []byte) {
	res, err := es.Client.Bulk(bytes.NewReader(body))
	if err != nil {
		log.Fatalf("sendBulkToES es.Client.Bulk: %v", err)
	}
	if res.IsError() {
		log.Fatalf("sendBulkToES es.Client.Bulk HTTP error: %v", res.String())
	}
	defer res.Body.Close()
}
