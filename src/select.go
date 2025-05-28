// Zettelliste abrufen und parsen

// 'package main' definiert das Hauptpaket der Anwendung
// ist notwendig, damit das Programm ausgeführt werden kann
package main

// 'import' importiert die benötigten Pakete für die Funktionalität
// 'bytes' importiert das Paket für Byte-Operationen
// 'io' importiert das Paket für Ein-/Ausgabe-Operationen
// 'net/http' importiert das Paket für HTTP-Kommunikation
// 'regexp' importiert das Paket für reguläre Ausdrücke
// 'strings' importiert das Paket für String-Operationen
import (
	"bytes"
	"io"
	"net/http"
)

// 'get_zettel_list' ruft eine Liste von Zetteln vom Zettelstore über HTTP ab
// dient als Schnittstelle zum Abrufen der Zettelliste
func get_zettel_list() ([]ZettelListEntry, string) {

	// 'http.Get' führt einen HTTP-GET-Request an die URL des Zettelstores aus
	// notwendig um die Zettelliste zu erhalten
	resp, err := http.Get(ZETTELSTORE_URL + "/z")

	// if-Schleife prüft, ob beim HTTP-Request ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Failed to retrieve the zettel list. Please ensure that Zettelstore is running."
	}

	// 'resp.Body.Close' schließt die HTTP-Verbindung am Ende der Funktion automatisch
	// verhindert, dass offene Verbindungen Ressourcen verbrauchen
	defer resp.Body.Close()

	// 'parse_zettel_list' verarbeitet den Body der HTTP-Antwort
	// ermöglicht eine bessere Testbarkeit und Wiederverwendbarkeit der Funktion
	entries, errMsg := parse_zettel_list(resp.Body)

	// if-Schleife prüft, ob beim Parsen der Zettelliste ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if errMsg != "" {
		return nil, "Error while parsing the zettel list."
	}

	// 'getZettelIDToTagsMap' holt die Map ZettelID -> Tags
	// ruft die Funktion auf, die für jede Zettel-ID die zugehörigen Tags liefert
	zettelIDToTags, tagErr := getZettelIDToTagsMap()

	// if-Schleife prüft, ob beim Abrufen der Tags ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if tagErr != "" {
		return nil, "Error while retrieving tags."
	}

	// for-Schleife weist jedem Zettel in der Liste die passenden Tags zu
	// ermöglicht die Zuordnung von Tags zu den Zetteln
	for i, entry := range entries {
		entries[i].Tags = zettelIDToTags[entry.Id]
	}

	// 'return' gibt die fertige Zettelliste (inklusive Tags) und eine leere Fehlermeldung zurück
	// liefert das Ergebnis an die Weboberfläche oder andere Aufrufer
	return entries, ""
}

// 'parse_zettel_list' verarbeitet die Zettelliste aus einem Reader
// ermöglicht das Parsen der Zettelliste unabhängig von der Datenquelle
func parse_zettel_list(r io.Reader) ([]ZettelListEntry, string) {

	// 'entries' definiert eine leere Liste von ZettelListEntry
	// dient als Speicher für die Zettel, die aus dem Reader gelesen werden
	var entries []ZettelListEntry

	// 'new(bytes.Buffer)' erstellt einen neuen Buffer, um die Daten aus dem Reader zu speichern
	// 'buf.ReadForm(r)' liest die Daten aus dem Reader in den Buffer
	// ermöglicht die Verarbeitung der gesamten Antwort als Ganzes
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)

	// if-Schleife prüft, ob beim Lesen der Daten ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Failed to read the zettel list. Please ensure that Zettelstore is running."
	}

	// 'bytes.Split' teilt den Buffer in einzelne Zeilen auf
	// ermöglicht die Verarbeitung jeder Zeile einzeln, um Zettel zu extrahieren
	lines := bytes.Split(buf.Bytes(), []byte("\n"))

	// for-Schleife iteriert über jede Zeile der Zettelliste
	// ermöglicht die Verarbeitung jeder Zeile, um Zettel-IDs und Namen zu extrahiere
	for _, line := range lines {

		// if-Schleife prüft, ob die Zeile leer ist
		// ermöglicht das Überspringen von leeren Zeilen, die keine Zettel enthalten
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}

		// 'bytes.SplitN' teilt die Zeile in zwei Teile: ID und Name
		// ermöglicht die Extraktion der Zettel-ID und des Namens
		parts := bytes.SplitN(line, []byte(" "), 2)

		// if-Schleife prüft, ob die Zeile mindestens ID und Name enthält
		// ermöglicht das Überspringen von Zeilen, die nicht genügend Informationen enthalten
		if len(parts) < 2 {
			continue
		}

		// 'strings.TrimLeft' entfernt führende Nullen von der Zettel-ID
		// 'string' wandelt die Byte-Slices in Strings um
		// ermöglicht die einfache Handhabung von Zettel-IDs und Namen als Strings
		id := string(parts[0])
		name := string(parts[1])

		// if-Schleife überprüft, ob ID und Name nicht leer sind
		// ermöglicht das Überspringen von ungültigen Einträgen, die keine Zettel repräsentieren
		if id == "" || name == "" {
			continue
		}

		// 'append' fügt einen neuen ZettelListEntry zur Liste hinzu
		// ermöglicht das Sammeln aller gültigen Zettel-IDs und Namen in der Liste
		entries = append(entries, ZettelListEntry{Id: id, Name: name})
	}

	// 'return' gibt die Liste der Zettel und eine leere Fehlermeldung zurück
	// liefert das Ergebnis an die Weboberfläche oder andere Aufrufer
	return entries, ""
}
