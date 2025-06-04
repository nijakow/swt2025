// Abruf der Zettel-Liste mit Tags
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// struct für einen Zettel mit ID, Name und Tags
type ZettelListEntry struct {
	Id   string
	Name string
	Tags []string
}

//1. Funktion: Ruft die Zettel-IDs und Tags vom Zettelstore ab

// 'getAllTagsWithZettelIDs' gibt eine Map zurück: Tag -> ZettelID
// ruft alle Tags mit zugehörigen Zettel-IDs vom Zettelstore ab
func getAllTagsWithZettelIDs() (map[string][]string, string) {

	// 'http.Get' führt einen HTTP-GET-Request an die Tag-API des Zettelstores aus
	// notwendig, um alle Tags mit ihren Zettel-IDs zu erhalten
	resp, err := http.Get(ZETTELSTORE_URL + "/z?q=|tags&enc=data")

	// if-Statement prüft, ob beim HTTP-Request ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Failed to retrieve the zettel list."
	}

	// 'resp.Body.Close()' schließt die HTTP-Verbindung am Ende der Funktion automatisch
	// verhindert, dass offene Verbindungen Ressourcen verbrauchen
	defer resp.Body.Close()

	// 'new(bytes.Buffer)' erstellt einen neuen Buffer, um die Antwort des HTTP-Requests zu speichern
	// 'buf.ReadFrom(resp.Body)' liest die Antwort des HTTP-Requests in einen Buffer
	// ermöglicht die Verarbeitung der gesamten Antwort als Ganzes
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)

	// if-Statement prüft, ob beim Lesen der Antwort ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Failed to read the zettel list."
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
		rawIDs := strings.Fields(match[2])
		var ids []string
		for _, id := range rawIDs {
			id14 := fmt.Sprintf("%014s", id) // ID auf 14 Stellen mit führenden Nullen auffüllen
			ids = append(ids, id14)
		}
		tagMap[tag] = ids
	}

	// 'return' gibt die Map und eine leere Fehlermeldung zurück
	// liefert das Ergebnis an die Weboberfläche oder andere Aufrufer
	return tagMap, ""
}

// 2. Funktion: Fügt den Zetteln die Tags hinzu (basiert auf der ersten Funktion)

// 'fetchZettelWithTags' gibt eine Liste von Zetteln mit ihren zugehörigen Tags zurück
// nutzt die zentrale Zettelliste und die Tag-Zuordnung
func fetchZettelWithTags() ([]ZettelListEntry, string) {
	// Hole alle einfachen Zettel (Id, Name)
	// 'get_zettel_list' ruft die Zettelliste vom Zettelstore ab
	// ermöglicht die weitere Verarbeitung der Zettel
	simpleZettel, errMsg := get_zettel_list()
	if errMsg != "" {
		// Fehlerbehandlung: Gibt eine Fehlermeldung zurück, falls die Zettelliste nicht geladen werden konnte
		return nil, errMsg
	}

	// Hole die Map ZettelID -> Tags
	// 'getZettelIDToTagsMap' ruft die Zuordnung von Zettel-IDs zu Tags ab
	// ermöglicht die Ergänzung der Zettel mit ihren Tags
	zettelIDToTags, tagErr := getZettelIDToTagsMap()
	if tagErr != "" {
		// Fehlerbehandlung: Gibt eine Fehlermeldung zurück, falls die Tag-Zuordnung nicht geladen werden konnte
		return nil, tagErr
	}

	// Baue die Liste mit Tags auf
	// Iteriert über alle Zettel und ergänzt sie um die zugehörigen Tags
	var entries []ZettelListEntry
	for _, sz := range simpleZettel {
		tags := zettelIDToTags[sz.Id] // Holt die Tags für die aktuelle Zettel-ID
		entries = append(entries, ZettelListEntry{
			Id:   sz.Id,
			Name: sz.Name,
			Tags: tags,
		})
	}

	// Gibt die fertige Liste von Zetteln mit Tags und eine leere Fehlermeldung zurück
	return entries, ""
}

// 3. Funktion: Kehrt die Tag-Zuordnung um

// 'getZettelIDToTagsMap" gibt eine Map zurück: ZettelID -> []Tag
// kehrt die Tag-Zuordnung um, sodass zu jeder Zettel-ID die zugehörigen Tags gefunden werden können
func getZettelIDToTagsMap() (map[string][]string, string) {

	// 'getAllTagsWithZettelIDs' holt die Map Tag -> []ZettelID
	// ruft die Funktion auf, die alle Tags mit ihren Zettel-IDs liefert
	tagMap, errMsg := getAllTagsWithZettelIDs()

	// if-Statement prüft, ob beim Abrufen der Tags ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if errMsg != "" {
		return nil, "Error while parsing the zettel list."
	}

	// 'zettelIDToTags' definiert eine leere Map, die Zettel-IDs auf ihre zugehörigen Tags abbildet
	// dient als Speicher für die Zettel-IDs und deren zugehörige Tags
	zettelIDToTags := make(map[string][]string)

	// for-Schleife iteriert über alle Tags und deren Zettel-IDs
	// ermöglicht die Zuordnung von Tags zu den jeweiligen Zettel-IDs
	for tag, ids := range tagMap {
		for _, id := range ids {

			// 'append' weist jedem Zettel die zugehörigen Tags hinzugefügt
			// ermöglicht die Zuordnung von Tags zu den Zettel-IDs
			zettelIDToTags[id] = append(zettelIDToTags[id], tag)
		}
	}

	// 'return' gibt die Map und eine leere Fehlermeldung zurück
	// liefert das Ergebnis an die Weboberfläche oder andere Aufrufer
	return zettelIDToTags, ""

}

// Dieses File wurde mit Hilfe von GitHub Copilot erstellt.
// Für folgende Prompts wurde GitHub Copilot verwendet:
// Schreibe auf der Grundlage des bestehenden Codes eine Funktion, die alle Tags mit ihren Zettel-IDs vom Zettelstore abruft.
// Schreibe auf der Grundlage des bestehenden Codes eine Funktion, die die Zettelliste mit den zugeordneten Tags zurückgibt, basierend auf der ersten Funktion.
// Schreibe auf der Grundlage des bestehenden Codes eine Funktion, die die Tag-Zuordnung umkehrt, sodass zu jeder Zettel-ID die zugehörigen Tags gefunden und angezeigt werden können.
// Für Kommentare wurde teilweise der automatische Vervollständigungsvorschlag von GitHub Copilot übernommen.
