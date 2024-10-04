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

	if *dir != "data" {
		helpe.CheckDir(*dir)
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
	var code int
	status := r.URL.Query().Get("code")
	if status == "" {
		code = http.StatusOK
	} else {
		fmt.Sscanf(status, "%d", &code)
	}
	w.WriteHeader(code)
	fmt.Fprintf(w, "Status code: %d OK\n", code)

	// fmt.Println(r.Method)

	// title := r.URL.Path[len(*dir)+2:]

	// folderName := *dir

	// profile := Profile{"Adil", []string{"programming", "listen music"}}

	// x, err := xml.MarshalIndent(profile, "", " ")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/xml")
	// w.Write(x)

	// fmt.Println(title)

	err1 := os.MkdirAll(*dir, 0o755)
	if err1 != nil {
		if os.IsExist(err1) {
			fmt.Fprintf(w, "Folder already exists\n")
			return
		} else {
			fmt.Fprintf(w, "Error creating folder: %v\n", err1)
		}
	}



	// resp, err := http.Get("http://localhost:" + *port + "/" + *dir + "/")
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error", err)
	// 	return
	// }
	// fmt.Println(resp)

	// p := loadPage(title)
	// fmt.Fprintln(w, "This is just the beginning", p.Title, p.Body)
}

func loadPage(title string) *Page {
	filename := title + ".csv"
	body, _ := os.ReadFile(filename) // ReadFile reads the named file and returns the contents.
	// fmt.Println(body)
	return &Page{Title: title, Body: body}
}
