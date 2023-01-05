package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir(`./static`))
	http.Handle(`/`, fileServer)
	http.HandleFunc(`/hello`, helloPage)
	http.HandleFunc(`/index`, indexPage)
	http.HandleFunc(`/from`, fromPage)
	http.HandleFunc(`/json`, jsonPage)

	log.Println(`start listen http://127.0.0.1:2123`)
	if err := http.ListenAndServe(`:2123`, nil); err != nil {
		log.Fatal(err)
	}
}

// helloPage hello page
func helloPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != `/hello` {
		http.Error(w, `404, not found`, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, `404, method not supported`, http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, `hello`)
}

// indexPage index page
func indexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != `/index` {
		http.Error(w, `404, not found`, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, `404, method not supported`, http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, `index page`)
}

// fromPage from page
func fromPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != `/from` {
		http.Error(w, `404, not found`, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, `404, method not supported`, http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, `parse from error`, err)
		return
	}
	name := r.FormValue(`name`)
	address := r.FormValue(`address`)
	fmt.Fprintf(w, "Name:%s\n", name)
	fmt.Fprintf(w, "Address:%s\n", address)
}

// jsonPage json page
func jsonPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, `application/json`)
	if r.URL.Path != `/json` {
		http.Error(w, `404, not found`, http.StatusNotFound)
		return
	}
	peoples := []struct {
		Id      uint64
		Name    string
		Age     uint8
		Gender  bool
		Height  float32
		Address string
	}{
		{1, `payne`, 21, false, 178.21, `cc`},
		{2, `Tom`, 23, false, 181.21, `cc`},
		{3, `Tim`, 23, true, 163.21, `cc`},
	}
	json.NewEncoder(w).Encode(peoples)
}
