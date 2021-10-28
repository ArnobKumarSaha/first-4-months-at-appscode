package main

import (
	"log"
	"os"
	"text/template"
)

type Student struct {
	Name string
}

type Person struct {
	Name   string
	Emails []string
}

func main() {
	s := Student{Name: "Arnob"}
	ss := Person{
		Name: "Masud",
		Emails: []string{
			"masud@gmail.com",
			"masud123@yahoo.com",
		},
	}
	tmp, err := template.New("test").Parse("Hello {{.Name}}!")
	if err != nil {
		log.Fatalf("tmp.Parse() failed with %s\n", err)
		return
	}

	err = tmp.Execute(os.Stdout, s)
	if err != nil {
		log.Fatalf("tmp.Execute() failed with %s\n", err)
		return
	}

	tmp2, err := template.New("test").Parse(
		`The name is {{.Name}}.
{{range .Emails}}
	His email id is {{.}}
{{end}}`)

	if err != nil {
		log.Fatalf("tmp.Parse() failed with %s\n", err)
		return
	}

	err = tmp2.Execute(os.Stdout, ss)
	if err != nil {
		log.Fatalf("tmp.Execute() failed with %s\n", err)
		return
	}

}
