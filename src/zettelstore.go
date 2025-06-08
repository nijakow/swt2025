package main

import "net/http"

func queryZettelstoreList(endpoint string, sorted bool) ([]SimpleZettel, string) {
	// Führt einen HTTP-GET-Request an den Zettelstore aus
	resp, err := http.Get(ZETTELSTORE_URL + endpoint)

	// Ruft die Funktion auf, um die Antwort zu parsen
	return parseSimpleZettelsFromResponse(resp, err, sorted)
}

func queryZettelstoreQuery(query string, sorted bool) ([]SimpleZettel, string) {
	return queryZettelstoreList("/z?q="+query, sorted)
}
