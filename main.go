package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := ":8080"

	// Serve /assets (images, pdfs, etc.)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Serve styles.css or other root-level static files
	http.Handle("/styles.css", http.FileServer(http.Dir(".")))

	// Serve index.html on root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			log.Printf("404 Not Found: %s\n", r.URL.Path)
			return
		}

		http.ServeFile(w, r, "index.html")
		log.Printf("200 OK: %s %s", r.Method, r.URL.Path)
	})

	log.Printf("🚀 Server running at http://localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("❌ Server failed: %v", err)
		os.Exit(1)
	}
}
