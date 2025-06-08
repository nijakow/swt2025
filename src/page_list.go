package main

import (
	"fmt"
	"net/http"
)

func pageList(w http.ResponseWriter, r *http.Request) {
	HandleCookies(w, r)
	zettels, e := getEnrichedZettelList()
	if e != nil {
		constructPage(w, "<h1>Fehler, Zettel konnten nicht abgerufen werden!</h1>")
	} else {
		constructPage(w,
			fmt.Sprintf(`
					<h1>Alle Zettel</h1>
					%s
				`,
				genZettelList(zettels),
			),
		)
	}
}
