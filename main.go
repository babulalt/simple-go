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
		userName := fmt.Sprintf("Hi %s", name)
		_, err := fmt.Fprint(w, userName)
		if ti != "" {
			log.Print(err)
		}
	})
	logrus.Info("Running on 8000 port...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
