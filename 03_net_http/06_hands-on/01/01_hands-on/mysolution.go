package main

import (
	"fmt"
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Welcome page")
}

func dog(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Dog Dog Doggy")
}

func name(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Shujaat")
}

func main()  {

	http.HandleFunc("/",index)
	http.HandleFunc("/dog/",dog)
	http.HandleFunc("/me", name)

	http.ListenAndServe(":8080",nil)

}
