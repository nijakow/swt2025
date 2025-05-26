package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Fetch zettel list
	zettels, errMsg := get_zettel_list()
	var zettelListHTML string
	if errMsg != "" {
		// Fehlermeldung aus select.go anzeigen
		zettelListHTML = errMsg
	} else {
		zettelListHTML = "<ul>"
		for _, z := range zettels {
			// Make a link with the zettel name and referencing ZETTELSTORE_URL + "/h/" + z.Id
			zettelListHTML += fmt.Sprintf("<li><a href=\"%s/h/%s\">%s</a></li>", ZETTELSTORE_URL, z.Id, z.Name)
		}
		zettelListHTML += "</ul>"
	}

	fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Welcome</title>
			<link rel="stylesheet" href="/static/css/styles.css">
        </head>
        <body>
            <nav class="zs-menu">
                <a href="/">Home</a>
                <a href="/download">Download ZIP</a>
                <a href="/query?query=example">Query</a>
            </nav>
            <main>
                <h1>Hello, World!</h1>
                <p>Welcome to the server!</p>
                <p><a href="/download">Download ZIP</a></p>
                <p>Or you can <a href="/query?query=example">query a file</a>.</p>
                <h2>Zettel List</h2>
                %s
            </main>
        </body>
        </html>
    `, zettelListHTML)
}

func query_downloader(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Missing 'query' query parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "You requested the file: %s\n", query)

	// Request to the Zettelstore
}
