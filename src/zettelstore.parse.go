package main

import (
	"bytes"
	"net/http"
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
	/*
	 * Der folgende Vergleich `a.Id < b.Id` wurde mithilfe von GitHub Copilot erzeugt.
	 */
	return a.Id < b.Id
}

func parseZettelstoreResponseContent(buffer *bytes.Buffer, err error, sorted bool) ([]SimpleZettel, string) {
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
		/*
		 * Der folgende Funktionsaufruf `sort.Slice(...)` wurde mithilfe von GitHub Copilot erstellt.
		 */
		sort.Slice(entries, func(i int, j int) bool {
			return simpleZettelCompare(&entries[i], &entries[j])
		})
	}

	// Gibt die fertige Liste von Zetteln und eine leere Fehlermeldung zurück
	return entries, ""
}

func parseZettelstoreResponse(resp *http.Response, err error, sorted bool) ([]SimpleZettel, string) {
	// if-Statement prüft, ob beim HTTP-Request ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Failed to read the response body."
	}

	// 'defer resp.Body.Close()' schließt die HTTP-Verbindung am Ende der Funktion automatisch
	// verhindert, dass offene Verbindungen Ressourcen verbrauchen
	defer resp.Body.Close()

	// Erstellt einen neuen Buffer, um die Antwort zu lesen
	buf := new(bytes.Buffer)
	// Liest den Inhalt der Antwort in den Buffer
	_, err = buf.ReadFrom(resp.Body)

	// Ruft die Funktion auf, um die Zettel zu parsen
	return parseZettelstoreResponseContent(buf, err, sorted)
}

func queryZettelstoreList(endpoint string, sorted bool) ([]SimpleZettel, string) {
	// Führt einen HTTP-GET-Request an den Zettelstore aus
	resp, err := http.Get(ZETTELSTORE_URL + endpoint)

	// Ruft die Funktion auf, um die Antwort zu parsen
	return parseZettelstoreResponse(resp, err, sorted)
}

// Teile dieses Codes wurden aus Datei select.simplezettel.go übernommen. Siehe Anmerkung dort.
//  - EFN
