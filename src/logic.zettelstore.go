package main

import "fmt"

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
