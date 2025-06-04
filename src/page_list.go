package main

import (
	"fmt"
	"net/http"
)

func pageList(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	constructPage(w,
		fmt.Sprintf(`
				<h1>Warenkorb</h1>
				%s
			`,
			genZettelList(listWarenkorb(session)),
		),
	)
}
