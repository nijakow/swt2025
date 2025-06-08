package main

import (
	"archive/zip"
	"bytes"
	"net/http"
)

func downloadFiles(w http.ResponseWriter, r *http.Request, files []File) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

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

	err := zipWriter.Close()

	if err != nil {
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

func apiDownload(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	selectedIds := session.GetZettels()

	if len(selectedIds) == 0 {
		http.Error(w, "No zettel selected", http.StatusBadRequest)
		return
	}

	downloadFiles(w, r, generateDownloadableFiles(selectedIds)())
}

func apiDownloadOld(w http.ResponseWriter, r *http.Request) {
	// Parse the selected zettel IDs from the form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	selectedIds := r.Form["zettelIds"]

	if len(selectedIds) == 0 {
		http.Error(w, "No zettel selected", http.StatusBadRequest)
		return
	}

	downloadFiles(w, r, generateDownloadableFiles(selectedIds)())
}
