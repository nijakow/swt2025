package main

import (
	"bytes"
	"sort"
)

// struct für einen Zettel mit ID und Name
type SimpleZettel struct {
	Id   string
	Name string
}

func simpleZettelCompare(a *SimpleZettel, b *SimpleZettel) bool {
	// Vergleicht zwei SimpleZettel nach ID
	// Gibt true zurück, wenn die ID von a kleiner ist als die von b
	return a.Id < b.Id
}

func parseZettelstoreResponse(buffer *bytes.Buffer, err error, sorted bool) ([]SimpleZettel, string) {
	// if-Statement prüft, ob beim Lesen der Antwort ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Failed to read the zettel list."
	}

	// 'lines' enthält alle Zeilen der Antwort (je Zettel eine Zeile)
	lines := bytes.Split(buffer.Bytes(), []byte("\n"))

	// 'entries' speichert die geparsten Zettel
	var entries []SimpleZettel

	// for-Schleife iteriert über alle Zeilen der Antwort
	// ermöglicht die Verarbeitung jedes Zettels
	for _, line := range lines {
		// if-Statement prüft, ob die Zeile leer ist
		// überspringt leere Zeilen
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		// 'parts' enthält die ID und den Namen des Zettels
		parts := bytes.SplitN(line, []byte(" "), 2)
		// if-Statement prüft, ob die Zeile korrekt formatiert ist
		// überspringt fehlerhafte Zeilen
		if len(parts) < 2 {
			continue
		}
		id := string(parts[0])   // Extrahiert die Zettel-ID
		name := string(parts[1]) // Extrahiert den Zettel-Namen
		// Fügt den Zettel dem struct SimpleZettel hinzu
		entries = append(entries, SimpleZettel{Id: id, Name: name})
	}

	// Sortieren der Zettel nach ID (falls gewünscht)
	if sorted {
		sort.Slice(entries, func(i int, j int) bool {
			return simpleZettelCompare(&entries[i], &entries[j])
		})
	}

	// Gibt die fertige Liste von Zetteln und eine leere Fehlermeldung zurück
	return entries, ""
}
