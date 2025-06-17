package main

import "net/http"

func fetchZettelstoreList(endpoint string, sorted bool) ([]SimpleZettel, string) {
	// Führt einen HTTP-GET-Request an den Zettelstore aus
	resp, err := http.Get(ZETTELSTORE_URL + endpoint)

	// Ruft die Funktion auf, um die Antwort zu parsen
	return parseSimpleZettelsFromResponse(resp, err, sorted)
}

func fetchZettelstoreAll(sorted bool) ([]SimpleZettel, string) {
	return fetchZettelstoreList("/z", sorted)
}

func fetchZettelstoreQuery(query string, sorted bool) ([]SimpleZettel, string) {
	// Escape the query string to ensure it is safe for use in a URL
	return fetchZettelstoreList("/z?q="+escapeHttpSafe(query), sorted)
}

func fetchZettelstoreContext(id string, sorted bool) ([]SimpleZettel, string) {
	return fetchZettelstoreQuery("CONTEXT "+id, sorted)
}

func getMetadataForZettel(id string) (SimpleZettelMeta, error) {
	resp, err := http.Get(ZETTELSTORE_URL + "/z/" + id + "?part=meta")

	return parseZettelMetadataFromResponse(resp, err)
}
