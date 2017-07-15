package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
)

func getData() string {
	name, _ := os.Hostname()

	result := "<h2>Hostname: " + name + "</h2>\n"
	result += "<h2>Environment Variables</h2>\n"
	result += "<pre style=\"white-space:pre-wrap\">\n"

	// Sort envs vars and remove non-printable chars - just to make sure
	// each var is on its own line and only one line. Makes grep'ing easier.
	vars := sort.StringSlice(os.Environ())
	vars.Sort()
	for _, env := range vars {
		result += strings.Map(func(r rune) rune {
			if r < ' ' {
				return -1
			}
			return r
		}, env) + "\n"
	}

	result += "</pre>"
	return result
}

func main() {
	addr, port := "0.0.0.0", "80"

	if tmp := os.Getenv("PORT"); tmp != "" {
		port = tmp
	}

	if len(os.Args) > 2 {
		addr, port = os.Args[1], os.Args[2]
	} else if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Printf("Will show:\n==========\n%s\n==========\n\n", getData())
	fmt.Printf("Listening on: %s:%s\n", addr, port)

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(getData()))
	})
	if err := http.ListenAndServe(addr+":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
