package main

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
