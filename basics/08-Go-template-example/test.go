package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person2 struct {
	Name    string
	Age     string
	IsMohan bool
}

func main() {

	data := []Person2{
		{"Masudur Rahman", "34", false},
		{"Tahsin Rahman", "12", true},
	}

	const style = `

{{ define "abc"}}

My name is {{.Name}}
I'm {{.Age}} years old.

{{if .IsMohan}}
And I'm mohan.
{{- else}}
And I'm not mohan.
{{end}}
{{end}}

{{template "abc" .}}


`

	tmp, err := template.New("newtemp").Parse(style)
	if err != nil {
		panic(err)
	}

	for _, r := range data {
		fmt.Println(r)
		err := tmp.Execute(os.Stdout, r)
		if err != nil {
			panic(err)
		}
	}
}






