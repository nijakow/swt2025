package main

import "os"

// const ZETTELSTORE_URL = "http://localhost:23123"
// const OUR_PORT = "8080"

var ZETTELSTORE_URL string
var OUR_PORT string

func process_command_line_args() {
    args := os.Args[1:]

    // Hilfe anzeigen, wenn --help übergeben wurde
    for _, arg := range args {
        if arg == "--help" || arg == "-h" {
            println("Verwendung:")
            println("  --zettelstore-url <URL>   URL des Zettelstores (Standard: http://localhost:23123)")
            println("  --port <PORT>             Port für den Server (Standard: 8080)")
            println("  --help                    Diese Hilfe anzeigen")
            os.Exit(0)
        }
    }

    // Process --zettelstore-url and --port arguments
    for i := 0; i < len(args); i++ {
        if args[i] == "--zettelstore-url" && i+1 < len(args) {
            ZETTELSTORE_URL = args[i+1]
            i++
        } else if args[i] == "--port" && i+1 < len(args) {
            OUR_PORT = args[i+1]
            i++
        }
    }

    if ZETTELSTORE_URL == "" {
        ZETTELSTORE_URL = "http://localhost:23123"
    }

    if OUR_PORT == "" {
        OUR_PORT = "8080"
    }

    println("ZETTELSTORE_URL: ", ZETTELSTORE_URL)
    println("OUR_PORT:        ", OUR_PORT)
}