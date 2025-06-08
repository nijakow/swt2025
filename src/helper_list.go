package main

// struct für einen Zettel mit ID, Name und Tags
type ZettelListEntry struct {
	Id   string
	Name string
	Tags []string
}

func genZettelList([]ZettelListEntry) string {
	return "<p>TODO</p>"
}
