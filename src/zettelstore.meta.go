package main

import (
	"fmt"
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

func getTitleOfZettel(id string) string {
	meta, err := getMetadataForZettel(id)
	if err != nil {
		return fmt.Sprintf(`(Untitled Zettel %s)`, id)
	}
	return meta.GetTitle()
}

func getZettelTitleById(id string) (string, error) {
	// This function should fetch the title of a zettel by its ID.
	// For now, we return a placeholder title.
	// In a real implementation, this would query the Zettelstore or database.
	return getTitleOfZettel(id), nil
}
