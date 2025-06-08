package main

import "net/http"

func fetchZettelstoreList(endpoint string, sorted bool) ([]SimpleZettel, string) {
	// Führt einen HTTP-GET-Request an den Zettelstore aus
	resp, err := http.Get(ZETTELSTORE_URL + endpoint)

	// Ruft die Funktion auf, um die Antwort zu parsen
	return parseSimpleZettelsFromResponse(resp, err, sorted)
}

func fetchZettelstoreQuery(query string, sorted bool) ([]SimpleZettel, string) {
	return fetchZettelstoreList("/z?q="+query, sorted)
}

func fetchZettelstoreContext(id string, sorted bool) ([]SimpleZettel, string) {
	return fetchZettelstoreQuery("CONTEXT "+id, sorted)
}
