package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// struct für Metadaten
type SimpleZettelMeta struct {
	Meta map[string]string
}

func (meta *SimpleZettelMeta) GetTitle() string {
	// Gibt den Titel des Zettels zurück, falls vorhanden
	// Falls kein Titel vorhanden ist, wird "Untitled" zurückgegeben
	// Diese Funktion wurde vollständig von GitHub Copilot generiert.
	if title, ok := meta.Meta["title"]; ok {
		return title
	}
	return "Untitled"
}

func parseZettelMetadata(buffer []byte) (SimpleZettelMeta, error) {
	// Hier wird angenommen, dass die Metadaten im Format "key: value" vorliegen
	// Diese Funktion wurde vollständig von GitHub Copilot generiert.
	meta := make(map[string]string)
	lines := bytes.Split(buffer, []byte("\n"))

	for _, line := range lines {
		parts := bytes.SplitN(line, []byte(":"), 2)
		if len(parts) == 2 {
			key := string(bytes.TrimSpace(parts[0]))
			value := string(bytes.TrimSpace(parts[1]))
			meta[key] = value
		}
	}

	return SimpleZettelMeta{Meta: meta}, nil
}

func getMetadataForZettel(id string) (SimpleZettelMeta, error) {
	// Anfrage ist: // GET ZETTELSTORE_URL + "/z/id?part=meta"
	resp, err := http.Get(ZETTELSTORE_URL + "/z/" + id + "?part=meta")

	if err != nil {
		return SimpleZettelMeta{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return SimpleZettelMeta{}, fmt.Errorf("Whoops %s: %s", id, resp.Status)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)

	if err != nil {
		return SimpleZettelMeta{}, fmt.Errorf("Whoops II %s: %s", id, err)
	}

	meta, err := parseZettelMetadata(buf.Bytes())

	if err != nil {
		return SimpleZettelMeta{}, fmt.Errorf("Whoops III %s: %s", id, err)
	}

	return meta, nil
}

func getTitleOfZettel(id string) string {
	meta, err := getMetadataForZettel(id)
	if err != nil {
		return fmt.Sprintf(`(Untitled Zettel %s)`, id)
	}
	return meta.GetTitle()
}
