package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	zettels, errMsg := get_zettel_list()
	var zettelListHTML string
	if errMsg != "" {
		zettelListHTML = fmt.Sprintf("<p>Error: %s</p>", errMsg)
	} else {
		zettelListHTML = "<ul>"
		for _, z := range zettels {
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
                <h2>Search Zettel</h2>
                <form action="/query" method="GET">
                    <input type="text" name="query" placeholder="Search for Zettel..." class="zs-input">
                    <button type="submit" class="zs-primary">Search</button>
                </form>
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

	zettels, errMsg := get_zettel_list()
	if errMsg != "" {
		http.Error(w, "Failed to fetch Zettel list: "+errMsg, http.StatusInternalServerError)
		return
	}

	var results []ZettelListEntry
	for _, z := range zettels {
		if containsIgnoreCase(z.Name, query) {
			results = append(results, z)
		}
	}

	w.Header().Set("Content-Type", "text/html")
	if len(results) == 0 {
		fmt.Fprintf(w, "<p>No results found for query: %s</p>", query)
	} else {
		fmt.Fprintf(w, "<h1>Search Results for '%s'</h1><ul>", query)
		for _, z := range results {
			fmt.Fprintf(w, "<li><a href=\"%s/h/%s\">%s</a></li>", ZETTELSTORE_URL, z.Id, z.Name)
		}
		fmt.Fprintf(w, "</ul>")
	}
}

func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
