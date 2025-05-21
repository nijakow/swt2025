package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func get_zettel_list() ([]ZettelListEntry, error) {
	// Fetch from the Zettelstore by using /z as the endpoint. Use a HTTP GET request.

	resp, err := http.Get(ZETTELSTORE_URL + "/z")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	/*
			The format of the output is:
			...
		00001012051200 API: List all zettel
		00001012050600 API: Provide an access token
		00001012050400 API: Renew an access token
		00001012050200 API: Authenticate a client
		...
	*/

	// Parse the response body to extract the list of zettels
	var entries []ZettelListEntry

	// Read the response body
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(buf.Bytes(), []byte("\n"))
	for _, line := range lines {
		// Each line is expected to be like: "00001012051200 The name of the zettel (can contain spaces)"
		parts := bytes.SplitN(line, []byte(" "), 2)
		if len(parts) < 2 {
			continue
		}

		id := string(parts[0])
		name := string(parts[1])
		entries = append(entries, ZettelListEntry{Id: id, Name: name})
	}
	return entries, nil
}
