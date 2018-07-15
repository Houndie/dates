//go:generate retool do go-bindata data

package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var dates = []string{"LaserTag!",
	"Go bowling!!! (Strike)",
	"Go Skating",
	"Go stargazing",
	"\"Tipsy\" Mini Golf :)",
	"Make Pasta + have a quiet diner at home like Lady + the Tramp",
	"Go bowling",
	"Bubble bath with candles :) (with wine of course)",
	"Rob a bank",
	"Dinner & movie",
	"Dress up fancy for dinner out at Salar! Pick up dessert to eat at home in your pjs & watch a movie",
	"Dinner and games @ Dave & Busters",
	"Take a hike @ Charleston Falls",
	"Go to a hockey game",
	"Picnic in the park with some wine & cheese (or beer)",
	"SPA & Margarita Day"}

type Template struct {
	Date string
	Idx  int
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	idx := r.FormValue("Not")
	var newIdx int

	if idx == "" {
		newIdx = rand.Intn(len(dates))
	} else {
		idxInt, err := strconv.Atoi(idx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		newIdx = idxInt
		for newIdx == idxInt {
			newIdx = rand.Intn(len(dates))
		}
	}

	htmlBytes, err := Asset("data/root.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t, err := template.New("root.html").Parse(string(htmlBytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := t.Execute(w, &Template{dates[newIdx], newIdx}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
