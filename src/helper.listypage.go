package main

import (
	"fmt"
	"net/http"
)

func handleListyPage(title string, w http.ResponseWriter, r *http.Request, zettels []ZettelListEntry, e error) {
	if e != nil {
		constructPage(w, "<h1>Fehler, Zettel konnten nicht abgerufen werden!</h1>")
	} else {
		constructPage(w,
			fmt.Sprintf(`
					<h1>%s</h1>
					%s
				`,
				title,
				genZettelList(zettels),
			),
		)
	}
}
