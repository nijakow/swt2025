package main

import (
	"fmt"
	"net/http"
)

// File struct represents a file with a name and content
type File struct {
	Name    string
	Content string
}

func main() {
	process_command_line_args()

	// Serve static files (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", pageHome)
	http.HandleFunc("/download", downloader)
	http.HandleFunc("/query", query_downloader)
	http.HandleFunc("/contextDisplay", contextDisplay)
	http.HandleFunc("/warenkorb", pageWarenkorb)
	http.HandleFunc("/list", pageList)
	http.HandleFunc("/about", aboutHandler)

	port := OUR_PORT
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
