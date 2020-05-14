package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc" : strings.ToUpper,
	"lw" : strings.ToLower,
	"tt" : strings.ToTitle,
}
func init()  {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main()  {

	err :=tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml","Shujaat")
	if err!=nil{
		log.Fatalln(err)
	}
}
