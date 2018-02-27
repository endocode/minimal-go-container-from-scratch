
package main

import (
    "log"
    "net/http"
)

func main() {
    d := "."
    p := "5000"
    http.Handle("/", http.FileServer(http.Dir(d)))

    log.Printf("Serving %s on HTTP port: %s\n", d, p)
    log.Printf("Ctrl+c to quit")
    log.Fatal(http.ListenAndServe(":"+p, nil))
}
