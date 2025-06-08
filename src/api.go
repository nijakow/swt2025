package main

import "net/http"

func handleApiRequestEnd(w http.ResponseWriter, r *http.Request, session *Session) {
	// Diese Funktion wird aufgerufen, um eine API-Anfrage zu beenden. Sie leitet ggf.
	// wieder auf die Ursprungsseite zurück.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

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
	// Wir leiten den Benutzer zurück zur Ursprungsseite
	handleApiRequestEnd(w, r, session)
}
