package main

import (
	"bytes"
	"net/http"
	"archive/zip"
)

func downloader(w http.ResponseWriter, r *http.Request) {
	// Parse the selected zettel IDs from the form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	selectedIds := r.Form["zettelIds"]
	readOnlyIds := r.Form["readOnlyIds"]
	if len(selectedIds) == 0 {
		http.Error(w, "No zettel selected", http.StatusBadRequest)
		return
	}

		// Create a map for quick lookup of read-only zettel IDs
	readOnlyMap := make(map[string]bool)
	for _, id := range readOnlyIds {
		readOnlyMap[id] = true
	}

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	files := gen_output2(selectedIds)()


	for _, file := range files {
		// If the file is marked as read-only, modify its content accordingly
		content := file.Content
		if readOnlyMap[file.Name[:len(file.Name)-4]] { // remove ".txt" from filename to get ID
			// Prepend or append a "Read Only" marker, for example:
			content = "// READ ONLY\n" + content
		}

		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			http.Error(w, "Failed to create ZIP file", http.StatusInternalServerError)
			return
		}
		_, err = fileWriter.Write([]byte(content))
		if err != nil {
			http.Error(w, "Failed to write to ZIP file", http.StatusInternalServerError)
			return
		}
	}

	if err := zipWriter.Close(); err != nil {
		http.Error(w, "Failed to finalize ZIP file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=files.zip")
	w.Header().Set("Content-Type", "application/zip")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, "Failed to send ZIP file", http.StatusInternalServerError)
	}
}