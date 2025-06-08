package main

func getZettelTitleById(id string) (string, error) {
	// This function should fetch the title of a zettel by its ID.
	// For now, we return a placeholder title.
	// In a real implementation, this would query the Zettelstore or database.
	return "Placeholder Title for " + id, nil
}

func enrichZettelId(id string) (ZettelListEntry, error) {
	title, err := getZettelTitleById(id)
	if err != nil {
		return ZettelListEntry{}, err
	}
	return ZettelListEntry{
		Id:   id,
		Name: title,
	}, nil
}

func enrichZettelIds(ids []string) ([]ZettelListEntry, error) {
	var entries []ZettelListEntry
	for _, id := range ids {
		entry, err := enrichZettelId(id)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func enrichSimpleZettel(zettel SimpleZettel) ZettelListEntry {
	return ZettelListEntry{
		Id:   zettel.Id,
		Name: zettel.Name,
		// Tags können hier später noch eingefügt werden
	}
}

func enrichSimpleZettelList(entries []SimpleZettel) []ZettelListEntry {
	var enrichedEntries []ZettelListEntry
	for _, zettel := range entries {
		enrichedEntries = append(enrichedEntries, enrichSimpleZettel(zettel))
	}
	return enrichedEntries
}
