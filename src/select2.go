package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Beispiel-Blacklist (IDs oder Namen, die ausgeschlossen werden sollen)
var blacklistIDs = map[string]bool{
	"00001012051200": true,
	// weitere IDs hinzufügen
}

// Prüft, ob ein Zettel auf der Blacklist steht
func isBlacklisted(id, name string) bool {
	return blacklistIDs[id]
}

// Ruft eine Liste von Zetteln vom Zettelstore über HTTP ab
func fetchZettelListZettelList() ([]ZettelListEntry, string) {
	resp, err := http.Get(ZETTELSTORE_URL + "/z")
	if err != nil {
		return nil, "Failed to retrieve the zettel list. Please ensure that Zettelstore is running."
	}
	defer resp.Body.Close()

	entries, errMsg := parsezettellist(resp.Body)
	return entries, errMsg
}

// Verarbeitet die Zettelliste aus einem Reader
func parsezettellist(r io.Reader) ([]ZettelListEntry, string) {
	var entries []ZettelListEntry
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, "Failed to read the zettel list. Please ensure that Zettelstore is running."
	}
	lines := bytes.Split(buf.Bytes(), []byte("\n"))
	for _, line := range lines {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		parts := bytes.SplitN(line, []byte(" "), 2)
		if len(parts) < 2 {
			continue
		}
		id := string(parts[0])
		name := string(parts[1])
		if id == "" || name == "" {
			continue
		}
		// Blacklist-Check
		if isBlacklisted(id, name) {
			continue
		}
		entries = append(entries, ZettelListEntry{Id: id, Name: name})
	}
	return entries, ""
}

// API-Handler, der die Zettelliste als JSON zurückgibt
func zettelListAPIHandler(w http.ResponseWriter, r *http.Request) {
	entries, errMsg := fetchZettelListZettelList()
	if errMsg != "" {
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
