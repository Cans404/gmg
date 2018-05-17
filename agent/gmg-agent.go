package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("edit.html")
	fmt.Fprintf(w, string(body))
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")

	command := exec.Command("bash", "-c", body)
	out, _ := command.CombinedOutput()
	ret := strings.TrimRight(string(out), "\n")

	fmt.Fprintln(w, ret)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	dir := r.FormValue("path")
	r.ParseMultipartForm(32 << 20)
	file, handler, _ := r.FormFile("uploadfile")
	defer file.Close()

	dest := path.Clean(dir + "/" + handler.Filename)
	f, _ := os.Create(dest)
	defer f.Close()

	io.Copy(f, file)
	w.Write([]byte("file uploaded!"))
}

func main() {
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/exec/", execHandler)
	http.HandleFunc("/upload/", uploadHandler)
	
	files := http.FileServer(http.Dir("js"))
	http.Handle("/web/", http.StripPrefix("/web/", files))

	http.ListenAndServe(":12580", nil)
}

