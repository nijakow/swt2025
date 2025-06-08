package main

import (
	"net/http"
)

func pageList(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	zettels, e := fetchEnrichedZettelstoreAll(session, true)
	handleListyPage("Alle Zettel", w, r, zettels, e)
}
