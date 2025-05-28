package main

import (
	"bytes"
	"net/http"
	"regexp"
	"strings"
)

// 'getAllTagsWithZettelIDs' gibt eine Map zurück: Tag -> ZettelID
// ruft alle Tags mit zugehörigen Zettel-IDs vom Zettelstore ab
func getAllTagsWithZettelIDs() (map[string][]string, string) {

	// 'http.Get' führt einen HTTP-GET-Request an die Tag-API des Zettelstores aus
	// notwendig, um alle Tags mit ihren Zettel-IDs zu erhalten
	resp, err := http.Get(ZETTELSTORE_URL + "/z?q=|tags&enc=data")

	// if-Schleife prüft, ob beim HTTP-Request ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Fehler beim Abrufen der Tags vom Zettelstore."
	}

	// 'resp.Body.Close()' schließt die HTTP-Verbindung am Ende der Funktion automatisch
	// verhindert, dass offene Verbindungen Ressourcen verbrauchen
	defer resp.Body.Close()

	// 'new(bytes.Buffer)' erstellt einen neuen Buffer, um die Antwort des HTTP-Requests zu speichern
	// 'buf.ReadFrom(resp.Body)' liest die Antwort des HTTP-Requests in einen Buffer
	// ermöglicht die Verarbeitung der gesamten Antwort als Ganzes
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)

	// if-Schleife prüft, ob beim Lesen der Antwort ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Fehler beim Lesen der Antwort vom Zettelstore."
	}

	// 'data' speichert die Antwort als String
	data := buf.String()

	// 'tagMap' speichert die Zuordnung Tag -> []ZettelID
	tagMap := make(map[string][]string)

	// 'regexp.MustCompile' erstellt einen Regex, um Tags und IDs aus der Antwort zu extrahieren
	// 're.FindAllStringSubmatch' findet alle Übereinstimmungen des Regex in der Antwort
	// ermöglicht die Extraktion von Tags und den zugehörigen Zettel-IDs
	re := regexp.MustCompile(`\("([^"]+)"\s+([0-9\s]+)\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	// for-Schleife iteriert über alle gefundenen Tag-Blöcke
	// ermöglicht die Verarbeitung jedes Tags und der zugehörigen Zettel-IDs
	for _, match := range matches {
		tag := match[1]
		ids := strings.Fields(match[2])
		tagMap[tag] = ids
	}

	// 'return' gibt die Map und eine leere Fehlermeldung zurück
	// liefert das Ergebnis an die Weboberfläche oder andere Aufrufer
	return tagMap, ""
}
