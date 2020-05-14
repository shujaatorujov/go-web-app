package main

import (
	"log"
	"os"
	"text/template"
)

func main()  {

	tpl, err :=template.ParseFiles("one.gmao")
	if err!=nil{
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err =tpl.ParseFiles("two.gmao","three.gmao")
	if err!=nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"two.gmao",nil)
	if err!=nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"three.gmao",nil)
	if err!=nil{
		log.Fatalln(err)
	}

	// Execute first finding template
	err = tpl.Execute(os.Stdout,nil)
	if err!=nil{
		log.Fatalln(err)
	}

}