package main

// 'getZettelIDToTagsMap" gibt eine Map zurück: ZettelID -> []Tag
// kehrt die Tag-Zuordnung um, sodass zu jeder Zettel-ID die zugehörigen Tags gefunden werden können
func getZettelIDToTagsMap() (map[string][]string, string) {

	// 'getAllTagsWithZettelIDs' holt die Map Tag -> []ZettelID
	// ruft die Funktion auf, die alle Tags mit ihren Zettel-IDs liefert
	tagMap, errMsg := getAllTagsWithZettelIDs()

	// if-Schleife prüft, ob beim Abrufen der Tags ein Fehler aufgetreten ist
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
