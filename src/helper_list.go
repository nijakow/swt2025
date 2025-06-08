package main

import (
	"fmt"
	"html"
	"strings"
)

// genZettelList erzeugt aus einer Liste von ZettelListEntry einen HTML-String mit einer Liste plus Checkbox pro Eintrag
func genZettelList(entries []ZettelListEntry) string {
	var builder strings.Builder
	builder.WriteString("<ul>\n")
	for _, e := range entries {
		idEscaped := html.EscapeString(e.Id)
		nameEscaped := html.EscapeString(e.Name)

		// checkbox input mit id basierend auf der Zettel-ID (eindeutig)
		checkboxID := "chk-" + idEscaped

		builder.WriteString("<li>\n")

		// checkbox input für anklicken
		builder.WriteString(fmt.Sprintf(`<input type="checkbox" id="%s" name="zettel" value="%s"/>`, checkboxID, idEscaped))
		// label mit for=checkboxID, Name anzeigen
		builder.WriteString(fmt.Sprintf(`<label for="%s">%s</label>`, checkboxID, nameEscaped))

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
