package main

import (
	"fmt"
	"io"
	"net/http"
)

// gen_output2 holt Inhalte per HTTP für die gegebenen Inputs
func generateDownloadableFiles(ids []string) func() []File {
	var result []File

	for _, name := range ids {
		url := ZETTELSTORE_URL + "/z/" + name + "?enc=zmk&part=zettel"

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Fehler bei Anfrage für %s: %v\n", name, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Fehler beim Lesen der Antwort für %s: %v\n", name, err)
			continue
		}

		result = append(result, File{
			Name:    name + ".zettel",
			Content: string(body),
		})
	}

	return func() []File {
		return result
	}
}
