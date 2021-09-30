package main

// Template caching , Validation, Literals & Closure section ... is not added yet

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
	/*
	func WriteFile(filename string, data []byte, perm fs.FileMode) error
	    WriteFile writes data to a file named by filename. If the file does not
	    exist, WriteFile creates it with permissions perm (before umask); otherwise
	    WriteFile truncates it before writing, without changing permissions.
	*/
	// 0600 indicates that the file should be created with read-write permissions for the current user only
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	/*
	func ReadFile(filename string) ([]byte, error)
	    ReadFile reads the file named by filename and returns the contents. A
	    successful call returns err == nil, not err == EOF. Because ReadFile reads
	    the whole file, it does not treat an EOF from Read as an error to be
	    reported.
	*/
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	/*
	func ParseFiles(filenames ...string) (*Template, error)
	    ParseFiles creates a new Template and parses the template definitions from
	    the named files. The returned template's name will have the (base) name and
	    (parsed) contents of the first file. There must be at least one file. If an
	    error occurs, parsing stops and the returned *Template is nil.

	    When parsing multiple files with the same name in different directories, the
	    last one mentioned will be the one that results. For instance,
	    ParseFiles("a/foo", "b/foo") stores "b/foo" as the template named "foo",
	    while "a/foo" is unavailable.
	*/
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		/*
		func Error(w ResponseWriter, error string, code int)
		    Error replies to the request with the specified error message and HTTP code.
		    It does not otherwise end the request; the caller should ensure no further
		    writes are done to w. The error message should be plain text.
		 */
		return
	}
	err = t.Execute(w, p)
	/*
	func (t *Template) Execute(wr io.Writer, data interface{}) error
		Execute applies a parsed template to the specified data object, writing the
		output to wr. If an error occurs executing the template or writing its
		output, execution stops, but partial results may already have been written
		to the output writer. A template may be executed safely in parallel,
		although if parallel executions share a Writer the output may be
		interleaved.
	*/

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // StatusFound = 302
		/*
		func Redirect(w ResponseWriter, r *Request, url string, code int)
		    Redirect replies to the request with a redirect to url, which may be a path
		    relative to the request path.

		    The provided code should be in the 3xx range and is usually
		    StatusMovedPermanently, StatusFound or StatusSeeOther.

		    If the Content-Type header has not been set, Redirect sets it to "text/html;
		    charset=utf-8" and writes a small HTML body. Setting the Content-Type header
		    to any value, including nil, disables that behavior.

		*/
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}