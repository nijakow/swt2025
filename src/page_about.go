package main

import (
	"fmt"
	"net/http"
)

type AboutPageTableEntry struct {
	Name           string
	Matrikelnummer string
	Kuerzel        string
	Link           string
}

func genAboutPageTableWrapLink(name string, link string) string {
	if link == "" {
		return name
	}
	return fmt.Sprintf("<a href=\"%s\">%s</a>", link, name)
}

func genAboutPageTableWrapEmail(email string) string {
	if email == "" {
		return ""
	}
	return fmt.Sprintf("<a href=\"mailto:%s\">%s</a>", email, email)
}

func genAboutPageTable(entries []AboutPageTableEntry) string {
	// Das Grundgerüst dieser Funktion ist mit GitHub Copilot erstellt und danach manuell angepasst worden.
	var tableHTML string
	tableHTML += "<table>\n"
	tableHTML += "<tr><th>Name</th><th>Matrikelnummer</th><th>Kürzel</th><th>Email</th></tr>\n" // Manuelle Anpassung
	for _, entry := range entries {
		email := fmt.Sprintf(`%s@stud.hs-heilbronn.de`, entry.Kuerzel) // Manuelle Anpassung
		tableHTML += "<tr>"
		tableHTML += fmt.Sprintf("<td>%s</td>", genAboutPageTableWrapLink(entry.Name, entry.Link)) // Manuelle Anpassung
		tableHTML += fmt.Sprintf("<td>%s</td>", entry.Matrikelnummer)
		tableHTML += fmt.Sprintf("<td>%s</td>", entry.Kuerzel)                     // Manuelle Anpassung
		tableHTML += fmt.Sprintf("<td>%s</td>", genAboutPageTableWrapEmail(email)) // Manuelle Anpassung
		tableHTML += "</tr>\n"
	}
	tableHTML += "</table>\n"
	return tableHTML
}

func pageAbout(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	zettelstoreURL := ZETTELSTORE_URL
	table := []AboutPageTableEntry{
		{
			Name:           "Eric Felix Nijakowski",
			Matrikelnummer: "221107",
			Kuerzel:        "enijakowsk",
			Link:           "https://github.com/nijakow",
		},
		{
			Name:           "Lamia Oktay",
			Matrikelnummer: "218915",
			Kuerzel:        "loktay",
			Link:           "",
		},
		{
			Name:           "Mary Williams",
			Matrikelnummer: "221352",
			Kuerzel:        "mwilliams",
			Link:           "",
		},
		{
			Name:           "Mia Braun",
			Matrikelnummer: "220039",
			Kuerzel:        "braun1",
			Link:           "",
		},
		{
			Name:           "Dariana Barkov",
			Matrikelnummer: "220039",
			Kuerzel:        "dbarkov",
			Link:           "",
		},
		{
			Name:           "Melih Akbulut",
			Matrikelnummer: "220860",
			Kuerzel:        "makbulut",
			Link:           "",
		},
		{
			Name:           "Stefanie Haag",
			Matrikelnummer: "221351",
			Kuerzel:        "shaag1",
			Link:           "",
		},
	}
	constructPage(w,
		fmt.Sprintf(`
            <h1>Über dieses Tool</h1>
            <p>Dieses Tool ermöglicht die Suche und den Export von Zetteln aus dem Zettelstore als ZIP-Datei.</p>
            <ul>
                <li>Suche mit Queries im Suchfeld</li>
                <li>Auswahl mehrerer Zettel für den Export</li>
                <li>Export als ZIP-Archiv</li>
            </ul>
            <p>Weitere Infos zu Startparametern siehe <code>--help</code> auf der Kommandozeile.</p>
			<h2>Entwickler</h2>
			%s
			<h2>Informationen</h2>
				<ul>
					<li>Session ID: %s</li>
					<li>Laufender Zettelstore: <a href="%s">%s</a></li>
				</ul>
    `,
			genAboutPageTable(table),
			session.Name,
			zettelstoreURL, zettelstoreURL,
		))
}
