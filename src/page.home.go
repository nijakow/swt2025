package main

import (
	"net/http"
)

func pageHome(w http.ResponseWriter, r *http.Request) {
	HandleCookies(w, r)
	constructPage(w,
		`
                <h1>Wissenszetteltransfer</h1>
                <p>Willkommen zum Wissenszetteltransfer!</p>
                <form action="/query" method="GET">
                    <input type="text" name="query" placeholder="Search for Zettel..." class="zs-input">
                    <button type="submit" class="zs-primary">Search</button>
                </form>
    	`,
	)
}
