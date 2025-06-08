package main

import "net/http"

func pageQuery(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	query := r.URL.Query().Get("query")
	zettels, e := queryEnrichedZettelstoreQuery(query, session, true)
	handleListyPage("Ergebnisse", w, r, zettels, e)
}
