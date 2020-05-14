package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err!=nil{
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout,nil)
	if err!=nil{
		log.Fatal(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
