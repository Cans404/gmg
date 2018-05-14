package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
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

func main() {
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/exec/", execHandler)
	http.ListenAndServe(":12580", nil)
}
