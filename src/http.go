package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	zettels, err := get_zettel_list()
	var zettelListHTML string
	if err != nil {
		zettelListHTML = "<p>Could not load zettel list. Please start the Zettelstore.</p>"
	} else {
		zettelListHTML = "<form action=\"/download\" method=\"post\"><ul>"
		for _, z := range zettels {
			zettelListHTML += fmt.Sprintf("<li><input type=\"checkbox\" name=\"zettelIds\" value=\"%s\"> %s</li>", z.Id, z.Name)
		}
		zettelListHTML += "</ul><input type=\"submit\" value=\"Download Selected Zettel\"></form>"
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
                <a href="#" onclick="document.getElementById('downloadForm').submit();">Download ZIP</a>
                <a href="/query?query=example">Query</a>
            </nav>
            <main>
                <h1>Hello, World!</h1>
                <p>Welcome to the server!</p>
                <h2>Zettel List</h2>
                <form id="downloadForm" action="/download" method="post">
                    <ul>
                        %s
                    </ul>
                    <input type="submit" value="Download Selected Zettel" style="display:none;">
                </form>
            </main>
        </body>
        <script>
            // Submit the form when the "Download ZIP" link is clicked
            document.querySelector('.zs-menu a[href="#"]').addEventListener('click', function(event) {
                event.preventDefault(); // Prevent the default link behavior
                document.getElementById('downloadForm').submit(); // Submit the form
            });
        </script>
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