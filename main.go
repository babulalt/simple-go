package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		var dur int = 1
		a := r.URL.Query()
		ti := a.Get("time")
		name := a.Get("name")
		if name == "" {
			name = "Gophers!!"
		}
		if ti != "" {
			time, _ := strconv.Atoi(ti)
			dur = int(time)
		}
		logrus.Info("Sleep Duration:: ", dur)
		time.Sleep(time.Duration(dur) * time.Second)
		fmt.Fprintf(w, fmt.Sprintf("Hi %s", name))
	})
	logrus.Info("Running on 8000 port...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
