package main

// struct für einen Zettel mit ID und Name
// Dieses Struct wurde vollständig von GitHub Copilot generiert.
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
