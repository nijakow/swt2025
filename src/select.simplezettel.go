// Abruf der Zettel-Liste
package main

import (
	"bytes"
	"net/http"
)

// struct für einen Zettel mit ID und Name
type SimpleZettel struct {
	Id   string
	Name string
}

// 'get_zettel_list' gibt eine Liste aller Zettel zurück (ohne Tags)
// ruft die Zettelliste vom Zettelstore ab und parst sie in SimpleZettel-Structs
func get_zettel_list() ([]SimpleZettel, string) {

	// 'http.Get' führt einen HTTP-GET-Request an den Zettelstore aus
	// notwendig, um die Zettelliste zu erhalten
	resp, err := http.Get(ZETTELSTORE_URL + "/z")

	// if-Statement prüft, ob beim HTTP-Request ein Fehler aufgetreten ist
	// ermöglicht die Fehlerbehandlung und Rückgabe einer Fehlermeldung
	if err != nil {
		return nil, "Failed to retrieve the zettel list."
	}

	// 'defer resp.Body.Close()' schließt die HTTP-Verbindung am Ende der Funktion automatisch
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

	// 'lines' enthält alle Zeilen der Antwort (je Zettel eine Zeile)
	lines := bytes.Split(buf.Bytes(), []byte("\n"))

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

	// Gibt die fertige Liste von Zetteln und eine leere Fehlermeldung zurück
	return entries, ""
}

// Dieses File wurde mit Hilfe von GitHub Copilot erstellt.
// Für folgende Prompts wurde GitHub Copilot verwendet:
// Schreibe auf der Grundlage des bestehenden Codes eine Funktion, die ID und Name von Zetteln als struct zurückgibt.
// Die Funktion soll die Zettelliste vom Zettelstore abrufen und im SimpleZettel-struct parsen.
// Für Kommentare wurde teilweise der automatische Vervollständigungsvorschlag von GitHub Copilot übernommen.
