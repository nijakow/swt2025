package main

import (
	"fmt"
	"net/http"
	"strings"
)

func constructPage(w http.ResponseWriter, content string) {
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
				<a href="/warenkorb">Warenkorb!!!</a>
			</nav>
			<main>
				%s
			</main>
		</body>
		</html>
	`, content)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	zettels, errMsg := fetchZettelWithTags()
	var zettelListHTML string
	if errMsg != "" {
		zettelListHTML = fmt.Sprintf("<p>Error: %s</p>", errMsg)
	} else {
		zettelListHTML = "<ul>"
		for _, z := range zettels {
			zettelListHTML += fmt.Sprintf(
				`<li>%s <a href="%s/h/%s">%s</a> %s</li>`,
				z.Id, ZETTELSTORE_URL, z.Id, z.Name, strings.Join(z.Tags, ", "),
			)
		}
		zettelListHTML += "</ul>"
	}

	constructPage(w,
		fmt.Sprintf(`
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
    `, zettelListHTML),
	)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <head><title>About</title></head>
        <body>
            <h1>Über dieses Tool</h1>
			<p>Session: %s</p>
            <p>Dieses Tool ermöglicht die Suche und den Export von Zetteln aus dem Zettelstore als ZIP-Datei.</p>
            <ul>
                <li>Suche mit Queries im Suchfeld</li>
                <li>Auswahl mehrerer Zettel für den Export</li>
                <li>Export als ZIP-Archiv</li>
            </ul>
            <p>Weitere Infos zu Startparametern siehe <code>--help</code> auf der Kommandozeile.</p>
            <p><a href="/">Zurück zur Startseite</a></p>
        </body>
        </html>
    `,
		session.Name)
}

func query_downloader(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Missing 'query' query parameter", http.StatusBadRequest)
		return
	}

	zettels, errMsg := fetchZettelWithTags()
	if errMsg != "" {
		http.Error(w, "Failed to fetch Zettel list: "+errMsg, http.StatusInternalServerError)
		return
	}

	var results []ZettelListEntry
	for _, z := range zettels {
		if containsIgnoreCase(z.Name, query) {
			results = append(results, ZettelListEntry{
				Id:   z.Id,
				Name: z.Name,
				Tags: z.Tags,
			})
		}
	}

	w.Header().Set("Content-Type", "text/html")
	if len(results) == 0 {
		fmt.Fprintf(w, "<p>No results found for query: %s</p>", query)
	} else {
		fmt.Fprintf(w, "<h1>Search Results for '%s'</h1><ul>", query)
		for _, z := range results {
			fmt.Fprintf(w, "<li><a href=\"%s/h/%s\">%s</a> <span class='tags'>[%s]</span></li>",
				ZETTELSTORE_URL, z.Id, z.Name, strings.Join(z.Tags, ", "))
		}
		fmt.Fprintf(w, "</ul>")
	}
}

func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
