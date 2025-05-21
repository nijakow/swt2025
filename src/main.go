package main

import (
	"bytes"
	"fmt"
	"net/http"
)

const ZETTELSTORE_URL = "http://localhost:23123"
const OUR_PORT = "8080"

// File struct represents a file with a name and content
type File struct {
	Name    string
	Content string
}

type ZettelListEntry struct {
	Id   string
	Name string
}

func get_zettel_list() ([]ZettelListEntry, error) {
	// Fetch from the Zettelstore by using /z as the endpoint. Use a HTTP GET request.

	resp, err := http.Get(ZETTELSTORE_URL + "/z")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	/*
			The format of the output is:
			...
		00001012051200 API: List all zettel
		00001012050600 API: Provide an access token
		00001012050400 API: Renew an access token
		00001012050200 API: Authenticate a client
		...
	*/

	// Parse the response body to extract the list of zettels
	var entries []ZettelListEntry

	// Read the response body
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(buf.Bytes(), []byte("\n"))
	for _, line := range lines {
		// Each line is expected to be like: "00001012051200 The name of the zettel (can contain spaces)"
		parts := bytes.SplitN(line, []byte(" "), 2)
		if len(parts) < 2 {
			continue
		}

		id := string(parts[0])
		name := string(parts[1])
		entries = append(entries, ZettelListEntry{Id: id, Name: name})
	}
	return entries, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Fetch zettel list
	zettels, err := get_zettel_list()
	var zettelListHTML string
	if err != nil {
		zettelListHTML = "<p>Could not load zettel list. Please start the Zettelstore.</p>"
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

func main() {
	// Serve static files (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/download", downloader)
	http.HandleFunc("/query", query_downloader)

	port := OUR_PORT
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
