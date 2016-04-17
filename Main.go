package main

import (
	"net/http"
	"html/template"
	"os"
	"strings"
	"bufio"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]
		template := templates.Lookup(requestedFile + ".html")
		if template != nil {
			template.Execute(w, nil)
		} else {
			w.WriteHeader(404)
		}

	})
	http.HandleFunc("/asset/img/",serveRequest)
	http.HandleFunc("/asset/css",serveRequest)
	http.ListenAndServe(":8070", nil)
}

func serveRequest(w http.ResponseWriter, req *http.Request){
	path := "public" + req.URL.Path
	var contentType string
	if(strings.HasSuffix(path,".css"){
		contentType="text/css"
	}else if (strings.HasSuffix(path,".png"){
		contentType= "image/png"
	}else if (strings.HasSuffix(path,".jpg"){
		contentType= "image/jpg"
	}else if (strings.HasSuffix(path,".gif"){
		contentType= "image/gif"
	}else {
		contentType="text/plain"
	}
	f,err:=os.Open(path)

	if err!=nil{
		defer f.Close()
		w.Header().Add("Content-Type",contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	}else{
		w.WriteHeader(404)
	}


}

func populateTemplates() *template.Template {
	result := template.New("templates")
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathRow, _ := templateFolder.Readdir(-1)
	templatePaths := new([] string)

	for _, pathInfo := range templatePathRow {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath + "/" + pathInfo.Name())
		}

	}
	result.ParseFiles(*templatePaths...)
	return result
}