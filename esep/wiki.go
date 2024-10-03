package main

import (
	// "encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"esep/helpe"
	// "path/filepath"
)

var (
	dir  = flag.String("dir", "data", "Path to the directory")
	port = flag.String("port", "7070", "Port number")
	help = flag.Bool("help", false, " Show information")
)

type Profile struct {
	Name    string
	Hobbies []string `xml:"Hobbies>Hobby"`
}

type Page struct {
	Title string
	Body  []byte
}

func main() {
	flag.Parse()

	if *help {
		helpe.HelpShow()
		return
	}

	portNum := ":" + *port

	var dirName string
	dirName += "/"
	dirName += *dir
	dirName += "/"
	fmt.Println(dirName)

	fmt.Println("http://localhost:" + *port + "/" + *dir)

	http.HandleFunc(dirName, viewHandler)

	// http.HandleFunc(dirName, foo)
	log.Fatal(http.ListenAndServe(portNum, nil))
}

// func foo(w http.ResponseWriter, r *http.Request) {
// 	profile := Profile{"Alex", []string{"listen music", "programming"}}

// 	x, err := xml.MarshalIndent(profile, "", " ")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/xml")
// 	w.Write(x)
// }

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(*dir)+2:]

	// profile := Profile{"Adil", []string{"programming", "listen music"}}

	// x, err := xml.MarshalIndent(profile, "", " ")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/xml")
	// w.Write(x)

	// fmt.Println(title)

	// err1 := os.Mkdir("kinder", 0o755)
	// if err1 != nil {
	// 	fmt.Println("Error creating file", err1)
	// 	os.Exit(1)
	// }

	// resp, err := http.Get("http://localhost:" + *port + "/" + *dir + "/")
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error", err)
	// 	return
	// }
	// fmt.Println(resp)

	p := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1>This is just the beginning<div>%s</div>", p.Title, p.Body)
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename) // ReadFile reads the named file and returns the contents.
	// fmt.Println(body)
	return &Page{Title: title, Body: body}
}
