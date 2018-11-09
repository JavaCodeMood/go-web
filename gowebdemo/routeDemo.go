package main

import (
	"net/http"
	"fmt"
)

type MyMux struct {

}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/" {
		helloName(w,r)
		return
	}
	http.NotFound(w,r)
	return
}

func helloName(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello myroute!")
}

func main(){
	mux := &MyMux{}
	http.ListenAndServe(":9090",mux)
}

