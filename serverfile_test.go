package web_golang

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)


func ServeFile(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Query().Get("name") != ""{
		http.ServeFile(w,r,"./resources/ok.html")
	}else{
		http.ServeFile(w,r,"./resources/notfound.html")
	}
	
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}

}
//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNot string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Query().Get("name") != ""{
		fmt.Fprint(w, resourceOk)
	}else{
		fmt.Fprint(w, resourceNot)
	}
	
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}

}