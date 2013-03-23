package router

import (
	"fmt"
	"github.com/VoiiD/goblog-prototype/db"
	"github.com/gorilla/mux"
	//"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

const authPath = len("/author/")

func authorHandler(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Path[authPath:]
	// Fetch from DB

	renderTemplate(w, "authorview", author)
}

const postPath = len("/post/")

func postHandler(w http.ResponseWriter, r *http.Request) {
	post := r.URL.Path[postPath:]

	// Fetch from DB

	renderTemplate(w, "postview", post)
}

const tagPath = len("/tag/")

func tagHandler(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Path[tagPath:]

	// Fetch from DB

	renderTemplate(w, "tagview", tag)
}

//var templates = template.Must(template.ParseFiles("postview", "tagview",
//	"home", "authorview"))

func renderTemplate(w http.ResponseWriter, tmpl string, s string) {
	//t, _ := template.ParseFiles(tmpl + ".html")
	//t.Execute(w)
}

func createDB() (*db.Persister, error) {
	return db.NewPersistance(db.NewSQLiter("test"))
}

func main() {
	persist, err := createDB()
	if err != nil {
		return
	}

	user := persist.NewUser()

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/author/{authorname}", authorHandler)
	r.HandleFunc("/post/{postname}", postHandler)
	r.HandleFunc("/tag/{tagname}", tagHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
