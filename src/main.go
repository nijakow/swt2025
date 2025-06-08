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

	/*
	 * API
	 */
	http.HandleFunc("/api/add", apiAdd)
	http.HandleFunc("/api/remove", apiRemove)

	/*
	 * Static
	 */
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	/*
	 * Pages
	 */
	http.HandleFunc("/", pageHome)
	http.HandleFunc("/list", pageList)
	http.HandleFunc("/warenkorb", pageWarenkorb)
	http.HandleFunc("/about", pageAbout)
	http.HandleFunc("/query", pageQuery)

	/*
	 * Debug endpoints
	 */
	http.HandleFunc("/download", apiDownload)
	http.HandleFunc("/contextDisplay", contextDisplay)

	port := OUR_PORT
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
