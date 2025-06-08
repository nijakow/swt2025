// Abruf der Zettel-Liste
package main

import "fmt"

// 'get_zettel_list' gibt eine Liste aller Zettel zurück (ohne Tags)
// ruft die Zettelliste vom Zettelstore ab und parst sie in SimpleZettel-Structs
func get_zettel_list() ([]SimpleZettel, string) {
	return queryZettelstoreList("/z", true)
}

func getEnrichedZettelList() ([]ZettelListEntry, error) {
	simpleZettels, e := get_zettel_list()
	if e != "" {
		return nil, fmt.Errorf("Fehler! %s", e)
	}
	// Enrich the simple zettels with their titles
	return enrichSimpleZettelList(simpleZettels), nil
}

// Verbleibende Anmerkung aus Datei vor Änderungen am 08. Juni 2025:
// Dieses File wurde mit Hilfe von GitHub Copilot erstellt.
// Für folgende Prompts wurde GitHub Copilot verwendet:
// Schreibe auf der Grundlage des bestehenden Codes eine Funktion, die ID und Name von Zetteln als struct zurückgibt.
// Die Funktion soll die Zettelliste vom Zettelstore abrufen und im SimpleZettel-struct parsen.
// Für Kommentare wurde teilweise der automatische Vervollständigungsvorschlag von GitHub Copilot übernommen.
