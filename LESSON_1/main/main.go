package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var pong = []byte("pong\n")

type Pinger struct {
}

func (p *Pinger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write(pong)
}

func main() {
	p := &Pinger{}
	http.Handle("/", p)
	http.HandleFunc("/ping", p.ServeHTTP)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
