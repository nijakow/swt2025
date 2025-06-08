package main

import (
	"fmt"
	"net/http"
)

func pageWarenkorb(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	warenkorb, err := listWarenkorb(session)
	warenkorbHtml := "<p>Warenkorb ist leer.</p>"
	if err != nil {
		warenkorbHtml = fmt.Sprintf("<p>Fehler beim Abrufen des Warenkorbs: %v</p>", err)
	} else if len(warenkorb) > 0 {
		warenkorbHtml = genZettelList(warenkorb)
	}
	constructPage(w,
		fmt.Sprintf(`
				<h1>Warenkorb</h1>
				%s
			`,
			warenkorbHtml,
		),
	)
}
