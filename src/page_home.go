package main

import (
	"fmt"
	"net/http"
)

func pageHome(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	zettelstoreURL := ZETTELSTORE_URL
	constructPage(w,
		fmt.Sprintf(`
                <h1>Wissenszetteltransfer</h1>
                <p>Willkommen zum Wissenszetteltransfer!</p>
                <form action="/query" method="GET">
                    <input type="text" name="query" placeholder="Search for Zettel..." class="zs-input">
                    <button type="submit" class="zs-primary">Search</button>
                </form>
				<h2>Informationen</h2>
				<ul>
					<li>Session ID: %s</li>
					<li>Laufender Zettelstore: <a href="%s">%s</a></li>
				</ul>
    	`,
			session.Name,
			zettelstoreURL, zettelstoreURL,
		),
	)
}
