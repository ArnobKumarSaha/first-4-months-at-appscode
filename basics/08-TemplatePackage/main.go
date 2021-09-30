package main

// For more details : https://levelup.gitconnected.com/learn-and-use-templates-in-go-aa6146b01a38

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

var tpl = template.New("")

/*
type User struct {
	Name   string
	Coupon string
	Amount int64
}
var user = User{
Name:   "Rick",
Coupon: "IAMAWESOMEGOPHER",
Amount: 5000,
}*/

func allExamplesExceptFunction() {
	tpl = template.Must(template.ParseFiles("sliceOfStruct.gohtml"))
	//sages := []string{"Tom Cruise", "Jack Nicholson", "Demi Moore", "Kevin Bacon", "Wolfgang Bodison"}

	/*superheroes := map[string]string{
		"Lt. Daniel Kaffee":        "Tom Cruise",
		"Col. Nathan R. Jessep":    "Jack Nicholson",
		"Lt. Cdr. JoAnne Galloway": "Demi Moore",
		"Capt. Jack Ross": "Kevin Bacon",
		"Lance Cpl. Harold W. Dawson": "Wolfgang Bodison",
	}*/
	type superhero struct {
		Name  string
		Motto string
	}
	im := superhero{
		Name:  "Iron man",
		Motto: "I am iron man",
	}
	ca := superhero{
		Name:  "Captain America",
		Motto: "Avengers assemble",
	}
	ds := superhero{
		Name:  "Doctor Strange",
		Motto: "I see things",
	}
	superheroes := []superhero{im, ca, ds}
	tpl.Execute(os.Stdout, superheroes)
}




// Func example start

func monthDayYear(t time.Time) string {
	return t.Format("January 2, 2006")
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// signature :: type FuncMap map[string]interface{}
	var fm = template.FuncMap{
		"fdateMDY": monthDayYear,
	}
	tpl := template.Must(tpl.Funcs(fm).ParseFiles("functions.gohtml"))
	if err := tpl.ExecuteTemplate(w, "functions.gohtml", time.Now()); err != nil {
		log.Fatalln(err)
	}
}
func main()  {
	http.HandleFunc("/hello", homeFunc)
	http.ListenAndServe(":8090", nil)
}