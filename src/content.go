package main

import (
	"fmt"
	"io"
	"net/http"
)

// gen_output generates a list of files to include in the ZIP archive
func gen_output() []File {
	return []File{
		{Name: "example1.txt", Content: "This is the content of example1.txt."},
		{Name: "example2.txt", Content: "This is the content of example2.txt."},
		{Name: "example3.txt", Content: "This is the content of example3.txt."},
	}
}

// gen_output2 holt Inhalte per HTTP für die gegebenen Inputs
func gen_output2(inputs []string) func() []File {
	var result []File

	for _, name := range inputs {
		url := ZETTELSTORE_URL + "/z" + name + "?enc=zmk"

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
			Name:    name + ".txt",
			Content: string(body),
		})
	}

	return func() []File {
		return result
	}
}
