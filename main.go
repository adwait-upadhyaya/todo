package main

import (
	"fmt"
	"html/template"
	"os"
)

type page struct {
	Header    string
	Paragraph string
}

func main() {

	p := page{
		Header:    "Adwait",
		Paragraph: "Very bored very bored",
	}

	fmt.Println(p)
	templatePath := "/home/leapfrog/Desktop/go/todo/index.html"
	t, error := template.New("index.html").ParseFiles(templatePath)

	if error != nil {
		fmt.Println(error)
	}

	error = t.Execute(os.Stdout, p)

	if error != nil {
		fmt.Println(error)

	}
}
