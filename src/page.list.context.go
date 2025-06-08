package main

import "net/http"

func pageContext(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	id := r.URL.Query().Get("id")
	zettels, e := fetchEnrichedZettelstoreContext(id, session, true)
	handleListyPage("Context", w, r, zettels, e)
}
