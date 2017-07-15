package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
)

func getData() string {
	name, _ := os.Hostname()

	result := "<h2>Hostname: " + name + "</h2>\n"
	result += "<h2>Environment Variables</h2>\n<pre style=\"white-space:pre-wrap\">\n"

	keys := sort.StringSlice(os.Environ())
	keys.Sort()
	for _, env := range keys {
		env1 := ""
		for _, ch := range env {
			if ch < ' ' {
				continue
			}
			env1 += string(ch)
		}
		result += env1 + "\n"
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

	fmt.Printf("Will show:\n%s\nListening on: %s:%s\n", getData(), addr, port)

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
