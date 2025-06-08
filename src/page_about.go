package main

import (
	"fmt"
	"net/http"
)

func pageAbout(w http.ResponseWriter, r *http.Request) {
	session := HandleCookies(w, r)
	zettelstoreURL := ZETTELSTORE_URL
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
			<h2>Informationen</h2>
				<ul>
					<li>Session ID: %s</li>
					<li>Laufender Zettelstore: <a href="%s">%s</a></li>
				</ul>
    `,
			session.Name,
			zettelstoreURL, zettelstoreURL,
		))
}
