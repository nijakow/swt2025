package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func constructPage(w http.ResponseWriter, content string) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Welcome</title>
			<link rel="stylesheet" href="/static/css/styles.css">
			<script src="/static/js/wzt.js"></script>
		</head>
		<body>
			<nav class="zs-menu">
				<a href="/">Home</a>
				<a href="/list">Zettelliste</a>
				<a href="/warenkorb">Warenkorb</a>
				<a href="/about">About</a>
			</nav>
			<main>
				%s
			</main>
		</body>
		</html>
	`, content)
}

func escapeHttpSafe(s string) string {
	// Return the string so that it can be safely used in an HTTP request as a URI parameter (e.g. by turning spaces into %20)
	return url.QueryEscape(s)
}
