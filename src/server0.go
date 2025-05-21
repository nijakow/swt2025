package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"net/http"
)

const ZETTELSTORE_URL = "http://localhost:23123"
const OUR_PORT = "8080"

