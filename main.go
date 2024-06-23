package main

import (
	"bank/handlers"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)




func main() {


	l:=log.New(os.Stdout,"product-api",log.LstdFlags)
	h:=mainPage.NewStart(l)
	mux:=http.NewServeMux()


	fileServer := http.FileServer(http.Dir("./templates/css"))
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))
	mux.Handle("/",h)
	mux.HandleFunc("/login",func ( w http.ResponseWriter, r *http.Request){

		t,err:=template.ParseFiles("D:/Bank/templates/html/login.html")
	
		if err!=nil{
			http.Error(w,"oops",http.StatusBadRequest)
		}
	
		err = t.Execute(w,nil)
	
		if err!=nil{
			http.Error(w,"oops",http.StatusBadRequest)
		}
	
	
	
	})

	mux.HandleFunc("/register",func ( w http.ResponseWriter, r *http.Request){
		
		t,err:=template.ParseFiles("D:/Bank/templates/html/registration.html")
	
		if err!=nil{
			http.Error(w,"oops",http.StatusBadRequest)
		}
	
		err = t.Execute(w,nil)
	
		if err!=nil{
			http.Error(w,"oops",http.StatusBadRequest)
		}
	
	
	
	})

	mux.HandleFunc("/transfer",func ( w http.ResponseWriter, r *http.Request){
	
		t,err:=template.ParseFiles("D:/Bank/templates/html/transfer.html")
	
		if err!=nil{
			http.Error(w,"oops",http.StatusBadRequest)
		}
	
		err = t.Execute(w,nil)
	
		if err!=nil{
			http.Error(w,"oops",http.StatusBadRequest)
		}
	
	
	
	})

	s:=http.Server{
		Addr:":8080",
		Handler: mux,
		IdleTimeout: 120*time.Second,


	}
	s.ListenAndServe()

	
}