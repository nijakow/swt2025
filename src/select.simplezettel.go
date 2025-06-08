// Abruf der Zettel-Liste
package main

import (
	"bytes"
	"net/http"
)

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

	return parseZettelstoreResponse(buf, err, true)
}

// Dieses File wurde mit Hilfe von GitHub Copilot erstellt.
// Für folgende Prompts wurde GitHub Copilot verwendet:
// Schreibe auf der Grundlage des bestehenden Codes eine Funktion, die ID und Name von Zetteln als struct zurückgibt.
// Die Funktion soll die Zettelliste vom Zettelstore abrufen und im SimpleZettel-struct parsen.
// Für Kommentare wurde teilweise der automatische Vervollständigungsvorschlag von GitHub Copilot übernommen.
