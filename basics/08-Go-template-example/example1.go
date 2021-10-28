package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Recipient struct {
	Name, Gift string
	Attended   bool
}
var recipients = []Recipient{
	{"Arnob Hasan", "iphone X", true},
	{"Masudur Rahman", "Nokia 3600", false},
	{"Fahim Abrar", "", false},
}

func main() {

	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It was a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	t := template.Must(template.New("letter").Parse(letter))

	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		fmt.Println("----------------------------------------------------------------")
		if err != nil {
			log.Fatalf("t.Execute() failed with %s\n", err)
		}
	}
}
