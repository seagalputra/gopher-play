package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func parsing(w http.ResponseWriter, r *http.Request) {
	urlString := "http://kalipare.com:80/hello?name=john wick&age=27"
	u, e := url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	name := u.Query()["name"][0]
	age := u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name":    "john wick",
			"Message": "have a nice day",
		}

		t, err := template.ParseFiles("./17-web-server/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)
	http.HandleFunc("/parsing", parsing)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
