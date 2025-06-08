package main

import "fmt"

func fetchEnrichedZettelstoreList(endpoint string, session *Session, sorted bool) ([]ZettelListEntry, error) {
	zettel, err := fetchZettelstoreList(endpoint, sorted)
	if err != "" {
		return nil, fmt.Errorf("%s", err)
	}

	// Anreichern der einfachen Zettel mit ihren Titeln
	return enrichSimpleZettelList(zettel, session), nil
}

func fetchEnrichedZettelstoreAll(session *Session, sorted bool) ([]ZettelListEntry, error) {
	zettel, err := fetchZettelstoreAll(sorted)
	if err != "" {
		return nil, fmt.Errorf("%s", err)
	}

	// Anreichern der einfachen Zettel mit ihren Titeln
	return enrichSimpleZettelList(zettel, session), nil
}

func fetchEnrichedZettelstoreQuery(query string, session *Session, sorted bool) ([]ZettelListEntry, error) {
	zettel, err := fetchZettelstoreQuery(query, sorted)
	if err != "" {
		return nil, fmt.Errorf("%s", err)
	}

	// Anreichern der einfachen Zettel mit ihren Titeln
	return enrichSimpleZettelList(zettel, session), nil
}

func fetchEnrichedZettelstoreContext(id string, session *Session, sorted bool) ([]ZettelListEntry, error) {
	zettel, err := fetchZettelstoreContext(id, sorted)
	if err != "" {
		return nil, fmt.Errorf("%s", err)
	}

	// Anreichern der einfachen Zettel mit ihren Titeln
	return enrichSimpleZettelList(zettel, session), nil
}
