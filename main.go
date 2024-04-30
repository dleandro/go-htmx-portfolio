package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "html/template"
    "io/ioutil"
)

func serveHTML(w http.ResponseWriter, r *http.Request) {
    // Load and parse the HTML file
    tmpl, err := template.ParseFiles("./src/index.html")
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Execute the template, which will write the HTML content to the response
    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}

func main() {
    // Create a new router using Gorilla Mux
    r := mux.NewRouter()

    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("src/static"))))

    // Use Gorilla Mux to specify the route and associate it with the handler function
    r.HandleFunc("/", serveHTML)
    
    r.HandleFunc("/zoom-in", func(w http.ResponseWriter, r *http.Request) {
        content, err := ioutil.ReadFile("./src/home.html")
        if err != nil {
            log.Println("Error reading HTML file:", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        // Send the content as the response to the client
        w.Header().Set("Content-Type", "text/html")
        w.Write(content)
    })

    // Start the server
    http.Handle("/", r)
    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}
