package main

import (
	"fmt"
	"html"
	"strings"
)

// struct für einen Zettel mit ID, Name und Tags
type ZettelListEntry struct {
	Id          string
	Name        string
	Tags        []string
	InWarenkorb bool // Ob im Warenkorb
}

// genZettelList erzeugt aus einer Liste von ZettelListEntry einen HTML-String mit einer Liste plus Checkbox pro Eintrag
func genZettelList(entries []ZettelListEntry) string {
	var builder strings.Builder
	builder.WriteString("<ul>\n")
	for _, e := range entries {
		idEscaped := html.EscapeString(e.Id)
		nameEscaped := html.EscapeString(e.Name)
		checkboxEnabled := "checked"
		functionToCall := "removeZettelFromWarenkorb"
		zettelURL := ZETTELSTORE_URL + "/h/" + idEscaped

		if !e.InWarenkorb {
			checkboxEnabled = ""
			functionToCall = "addZettelToWarenkorb"
		}

		// checkbox input mit id basierend auf der Zettel-ID (eindeutig)
		checkboxID := "chk-" + idEscaped

		builder.WriteString("<li>\n")

		// checkbox input für anklicken
		builder.WriteString(fmt.Sprintf(`<input type="checkbox" id="%s" name="%s" onclick="%s('%s')" %s>`, checkboxID, idEscaped, functionToCall, e.Id, checkboxEnabled))
		builder.WriteString(fmt.Sprintf(`<a href="/context?id=%s">🔍</a> `, idEscaped))
		builder.WriteString(fmt.Sprintf(`<label for="%s"><a href="%s">%s</a></label>`, checkboxID, zettelURL, nameEscaped))

		// Falls Tags vorhanden sind, diese in <small> listen
		if len(e.Tags) > 0 {
			builder.WriteString(" <small>(Tags: ")
			for i, tag := range e.Tags {
				if i > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(html.EscapeString(tag))
			}
			builder.WriteString(")</small>")
		}

		builder.WriteString("\n</li>\n")
	}
	builder.WriteString("</ul>")
	return builder.String()
}

// Diese Datei wurde in Teilen von BLACKBOX.AI generiert.
// Die folgenden Zeilen sind betroffen: (Z.3-7), (Z.17-23), (Z.33-38), (Z.43-58)
