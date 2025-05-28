package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func parse_zettel_List(resp *http.Response, err error) ([]ZettelListEntry, error) {
	// Parse the response body to extract the list of zettels
	var entries []ZettelListEntry

	// Read the response body
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Lesen des Response-Bodys: %w", err)
	}
	lines := bytes.Split(buf.Bytes(), []byte("\n"))
	for i, line := range lines {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		// Each line is expected to be like: "00001012051200 The name of the zettel (can contain spaces)"
		parts := bytes.SplitN(line, []byte(" "), 2)
		if len(parts) < 2 {
			fmt.Printf("Warnung: Zeile %d konnte nicht geparst werden: %q\n", i+1, line)
			continue
		}

		id := string(parts[0])
		name := string(parts[1])
		if id == "" || name == "" {
			fmt.Printf("Warnung: Leere ID oder Name in Zeile %d: %q\n", i+1, line)
			continue
		}
		entries = append(entries, ZettelListEntry{Id: id, Name: name})
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("keine Zettel gefunden oder alle Zeilen fehlerhaft")
	}
	return entries, nil
}

func freetextQuery(query string) ([]ZettelListEntry, error) {
	// Fetch from the Zettelstore by using /z as the endpoint. Use a HTTP GET request. Encode the query string.
	encodedQuery := url.QueryEscape(query)
	resp, err := http.Get(ZETTELSTORE_URL + "/z?q=" + encodedQuery)
	if err != nil {
		return nil, fmt.Errorf("fehler beim HTTP-Request: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			fmt.Printf("Warnung: Fehler beim Schließen des Response-Bodys: %v\n", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unerwarteter Statuscode vom Zettelstore: %d", resp.StatusCode)
	}

	return parse_zettel_List(resp, err)
}

func contextQuery(id string) ([]ZettelListEntry, error) {
	return freetextQuery("CONTEXT " + id)
}

func contextDisplay(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "You requested the context for the file with ID: %s\n", id)

	// Request to the Zettelstore
	entries, err := contextQuery(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Fehler beim Abrufen des Kontexts: %v", err), http.StatusInternalServerError)
		return
	}

	for _, entry := range entries {
		fmt.Fprintf(w, "%s: %s\n", entry.Id, entry.Name)
	}
}
