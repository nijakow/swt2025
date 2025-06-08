package main

import "fmt"

func getZettelTitleById(id string) (string, error) {
	// This function should fetch the title of a zettel by its ID.
	// For now, we return a placeholder title.
	// In a real implementation, this would query the Zettelstore or database.
	return getTitleOfZettel(id), nil
}

func enrichZettelId(id string, session *Session) (ZettelListEntry, error) {
	title, err := getZettelTitleById(id)
	if err != nil {
		return ZettelListEntry{}, err
	}
	return ZettelListEntry{
		Id:          id,
		Name:        title,
		InWarenkorb: session.ContainsZettel(id),
	}, nil
}

func enrichZettelIds(ids []string, session *Session) ([]ZettelListEntry, error) {
	var entries []ZettelListEntry
	for _, id := range ids {
		entry, err := enrichZettelId(id, session)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func enrichSimpleZettel(zettel SimpleZettel, session *Session) ZettelListEntry {
	return ZettelListEntry{
		Id:   zettel.Id,
		Name: zettel.Name,
		// Tags können hier später noch eingefügt werden
		InWarenkorb: session.ContainsZettel(zettel.Id),
	}
}

func enrichSimpleZettelList(entries []SimpleZettel, session *Session) []ZettelListEntry {
	var enrichedEntries []ZettelListEntry
	for _, zettel := range entries {
		enrichedEntries = append(enrichedEntries, enrichSimpleZettel(zettel, session))
	}
	return enrichedEntries
}

func queryEnrichedZettelstoreList(endpoint string, session *Session, sorted bool) ([]ZettelListEntry, error) {
	zettel, err := queryZettelstoreList(endpoint, sorted)
	if err != "" {
		return nil, fmt.Errorf("%s", err)
	}

	// Anreichern der einfachen Zettel mit ihren Titeln
	return enrichSimpleZettelList(zettel, session), nil
}

func queryEnrichedZettelstoreQuery(query string, session *Session, sorted bool) ([]ZettelListEntry, error) {
	zettel, err := queryZettelstoreQuery(query, sorted)
	if err != "" {
		return nil, fmt.Errorf("%s", err)
	}

	// Anreichern der einfachen Zettel mit ihren Titeln
	return enrichSimpleZettelList(zettel, session), nil
}
