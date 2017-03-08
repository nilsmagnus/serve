package main

import (
	"net/http"
	"os"
	"flag"
	"fmt"
)

var port int
var folder string

func main() {

	currentDir, _ := os.Getwd()

	flag.IntVar(&port, "port", 8080, "http port")
	flag.StringVar(&folder, "folder", currentDir, "folder to serve, defaults to current")
	flag.Parse()

	fmt.Printf("Serving %v on port %d\n", folder, port)
	fmt.Println("Press ctrl-c to end")

	http.Handle("/", http.FileServer(http.Dir(folder)))

	http.ListenAndServe(":8081", nil)
}
