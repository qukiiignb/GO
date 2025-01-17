package main

import (
	"html/template"
	"os"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

/*** global variable ***/
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

/** use */
func main(){
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080",nil))
}

/*** function clousure ***/
func makeHandler( fun  func(http.ResponseWriter, r *http.Request, string)) *http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m:= validPath.FindStringSubmach(r.URL.Path)
		if m == nil {
			http.NotFound(w,r)
			return
		}
		fn(w, r, m[2])
	}
]

/*** validate path and get title ***/
func getTitle(w http.ResponseWriter, r *http.Request) (string, error){
	m := validadPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", error.New("invalid Page Title")
	}
	return m[2], nil		//the title is the second subexpression 
}

/*** refactoring ***/
func renderTemplate(w http.ResponseWrite, tmpl string, p *Page){
	err := templates.ExecuteTemplete(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*** address persistent storage ***/
func (p *Page) save() error {		// *Page this mean that is a type of date, see line #10
					// save() is the name of the func
					// error this mean bacically that the func don't return any except and error if occour typeof error
	filename := p.Title + ".txt"		
	return os.WriteFile(filename, p.Body, 0600)		// os: documentation: https://pkg.go.dev/os#pkg-index
								// .WriteFile(): documentation: https://pkg.go.dev/os#example-WriteFile
								//exam: func WriteFile(name string, data []byte, perm FileMode) error
								//how you can see this func return a error you must handle this in the future
}

/*** load Pages ***/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"		// handle the parameter title
	body, err := os.ReadFile(filename) 	// The _ (underscore) symbol is used to ignore the second return value, which is an error indicating whether the file reading operation encountered any errors.
	if err != nil {
		return nil,err
	}
	return &Page{Title:title, Body:body},nil		//return a pointer to a Page literal constructed with the proper title and body values
}

/*** view Handler ***/
func viewHandler(w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage(title)
	if err != nil{
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w,"view",p)		// standart func of go that parce one or more template and create a new template that
						// be execueted with dynamic data.
}

/*** editHandler ***/
func editHandler(w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage(title)
	if err != nil{
		p = &Page{title:title}
	}
  	renderTemplate(w, "edit", p)
}
/*** saveHandler ***/
func saveHandler(w http.ResponseWriter, r *http.Resquest, title string){
	body := r.FormValue("body")
	p := &Page{title: title, Body: []byte(body)}
	err := p.save()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	http.redirect(w, r, "/view/" +title, http.StatusFound)
}




