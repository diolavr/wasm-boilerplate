package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

var defaultListen = ":8080"
var defaultAssets = "./assets"

func init() {
	flag.StringVar(&defaultListen, "listen", defaultListen, "Listen address")
	flag.StringVar(&defaultAssets, "assets", defaultAssets, "Assets directory")
	flag.Parse()
}

func checkFile(filename string) (string, error) {
	path := filepath.Join(defaultAssets, filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}
	return path, nil
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func wasmExecHandler(w http.ResponseWriter, req *http.Request) {
	const file = "wasm_exec.js"

	path, err := checkFile(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println(path)

	f, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, req, file, time.Now(), f)
}

func codeWasmHandler(w http.ResponseWriter, req *http.Request) {
	const file = "code.wasm"

	path, err := checkFile(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	f, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, req, file, time.Now(), f)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/wasm_exec.js", wasmExecHandler)
	mux.HandleFunc("/code.wasm", codeWasmHandler)

	log.Printf("Listening on '%s' ...", defaultListen)
	log.Fatal(http.ListenAndServe(defaultListen, mux))
}
