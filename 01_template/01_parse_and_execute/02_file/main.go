package main

import (
	"log"
	"os"
	"text/template"
)

func main()  {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err!=nil{
		log.Fatal(err)
	}

	rf, err :=os.Create("index.html")
	if err!=nil{
		log.Fatal(err)
	}
	defer rf.Close()

	err = tpl.Execute(rf,nil)
	if err!=nil{
		log.Fatal(err)
	}
}
