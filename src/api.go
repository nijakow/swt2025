package main

import "net/http"

func apiAdd(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	// Wir fragen die Zettel-ID über den URL-Parameter `id` ab
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}
	// Wir fügen den Zettel mit der angegebenen ID zum Warenkorb hinzu
	session.AddZettel(id)
}
