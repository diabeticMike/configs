package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Reader interface {
	Read(path string) (string, error)
}

type reader struct {
}

func (r *reader) Read(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

type controller struct {
	reader Reader
}

func (c *controller) users(w http.ResponseWriter, req *http.Request) {
	json, err := c.reader.Read("users.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "users upload error")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, json)
}

func main() {
	r := &reader{}
	c := &controller{reader: r}
	http.HandleFunc("/users", c.users)
	http.ListenAndServe(":8000", nil)
}
