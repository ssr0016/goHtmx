package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	// Step 1 Creating a web-server
	h1 := func(w http.ResponseWriter, r *http.Request) {

		// Step 2 Serving template from golang handler function
		tmp1 := template.Must(template.ParseFiles("index.html"))

		// Step 3 Adding server data into template
		films := map[string][]Film{
			"Films": {
				{Title: "The Matrix", Director: "Wachowski"},
				{Title: "The Matrix Reloaded", Director: "Wachowski"},
				{Title: "The Matrix Revolutions", Director: "Wachowski"},
			},
		}
		// Step 3 ends ... pass the data on tmp1.Execute(w, films)

		tmp1.Execute(w, films)
		// ->Step 2 ends
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		// Step 6 Submitting form with HTMX hx-post attribute
		// log.Print("HTMX request received")
		// log.Print(r.Header.Get("HX-Request"))
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		// Step 10 template fragments
		tmp1 := template.Must(template.ParseFiles("index.html"))
		tmp1.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})

		// // Step 7 Returning HTML from Go HTMX handler function
		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'> %s - %s</li>", title, director)
		// tmp1, _ := template.New("t").Parse(htmlStr)
		// tmp1.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Step 1 Creating a web-server
// Step 2 Serving template from golang handler function
// Step 3 Adding server data into template
// Step 4 Displaying server data in templates ... Films {{range .Films}} {{ end }}
// Step 5 Styling list with Bootstrap 5
// Step 6 Submitting form with HTMX hx-post attribute
// Step 7 Returning HTML from Go HTMX handler function
// Step 8 Set swapping element with hr-target attribute
// Step 9 Request feedback with hx-indicator attribute
// Step 10 template fragments
