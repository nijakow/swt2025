package main

func listWarenkorb(session *Session) ([]ZettelListEntry, error) {
	return enrichZettelIds(session.GetZettels())
}
