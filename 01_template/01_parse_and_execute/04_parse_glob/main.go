package main

import (
	"log"
	"os"
	"text/template"
)
var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("template/*.gmao"))
}

func main()  {

	err :=tpl.Execute(os.Stdout,nil)
	if err!=nil{
		log.Fatalln(err)
	}

	err=tpl.ExecuteTemplate(os.Stdout,"one.gmao",nil)
	if err!=nil{
		log.Fatalln(err)
	}
	err=tpl.ExecuteTemplate(os.Stdout,"two.gmao",nil)
	if err!=nil{
		log.Fatalln(err)
	}
	err=tpl.ExecuteTemplate(os.Stdout,"three.gmao",nil)
	if err!=nil{
		log.Fatalln(err)
	}
}
