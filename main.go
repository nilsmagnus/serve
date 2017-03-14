package main

import (
	"net/http"
	"os"
	"flag"
	"fmt"
)

func main() {

	currentDir, _ := os.Getwd()

	port:= flag.Int( "port", 8080, "http port")
	folder:= flag.String( "folder", currentDir, "folder to serve, defaults to current")
	flag.Parse()

	fmt.Printf("Serving %v on port %d\n", *folder, *port)
	fmt.Println("Press ctrl-c to end")

	http.Handle("/", http.FileServer(http.Dir(*folder)))

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
