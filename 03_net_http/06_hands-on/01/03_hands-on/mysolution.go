package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Welcome page")
}

func dog(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Dog Dog Doggy")
}

func name(w http.ResponseWriter, r *http.Request)  {
	tpl, err:=template.ParseFiles("something.gohtml")
	if err!=nil{
		log.Fatalln(err)
	}
	err=tpl.ExecuteTemplate(w,"something.gohtml","Shujaat")
	if err!=nil {
		log.Fatalln(err)
	}
}

func main()  {

	http.HandleFunc("/",index)
	http.HandleFunc("/dog/",dog)
	http.HandleFunc("/me", name)

	http.ListenAndServe(":8080",nil)

}
