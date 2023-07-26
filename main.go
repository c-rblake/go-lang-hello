package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// In this program, we're creating a new mux.Router, registering a function to be called whenever someone accesses the root (/) of our website, and starting a web server on port 8000. The function we're registering simply sends the text "Hello, World!" to the client.

// If you run this program and navigate to http://localhost:8000 in your web browser, you should see the text "Hello, World!".

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8000", r)
	HelloWorld()

}

// GIT create a repository from the command line. needs the CLI
// After you've created the repository on GitHub, you can push your local Git repository to GitHub using the git push command. Here's how
//git remote add origin https://github.com/c-rblake/go-lang-hello
//git branch -M main
// git push -u origin main
