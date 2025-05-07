package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"net/http"
)

// File struct represents a file with a name and content
type File struct {
	Name    string
	Content string
}

// gen_output generates a list of files to include in the ZIP archive
func gen_output() []File {
	return []File{
		{Name: "example1.txt", Content: "This is the content of example1.txt."},
		{Name: "example2.txt", Content: "This is the content of example2.txt."},
		{Name: "example3.txt", Content: "This is the content of example3.txt."},
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func downloader(w http.ResponseWriter, r *http.Request) {
	// Create a buffer to write the ZIP file to
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	// Get the list of files to include in the ZIP archive
	files := gen_output()

	// Add each file to the ZIP archive
	for _, file := range files {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			http.Error(w, "Failed to create ZIP file", http.StatusInternalServerError)
			return
		}
		_, err = fileWriter.Write([]byte(file.Content))
		if err != nil {
			http.Error(w, "Failed to write to ZIP file", http.StatusInternalServerError)
			return
		}
	}

	// Close the ZIP writer to finalize the archive
	if err := zipWriter.Close(); err != nil {
		http.Error(w, "Failed to finalize ZIP file", http.StatusInternalServerError)
		return
	}

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=files.zip")
	w.Header().Set("Content-Type", "application/zip")
	w.WriteHeader(http.StatusOK)

	// Write the ZIP file to the response
	_, err := w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, "Failed to send ZIP file", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/download", downloader)

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
