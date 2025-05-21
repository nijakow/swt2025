package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func get_zettel_list() ([]ZettelListEntry, error) {
	// Fetch from the Zettelstore by using /z as the endpoint. Use a HTTP GET request.

	resp, err := http.Get(ZETTELSTORE_URL + "/z")
	if err != nil {
		return nil, fmt.Errorf("Fehler beim HTTP-Request: %w", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			fmt.Printf("Warnung: Fehler beim Schließen des Response-Bodys: %v\n", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unerwarteter Statuscode vom Zettelstore: %d", resp.StatusCode)
	}

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
