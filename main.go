package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kellegous/fcgi"

	"strings"
)

func php(wd string, w http.ResponseWriter, r *http.Request, c *fcgi.Client) {
	params := fcgi.ParamsFromRequest(r)
	params["SCRIPT_FILENAME"] = []string{filepath.Join(wd, "index.php")}
	c.ServeHTTP(params, w, r)
}

func checkRequest(r *http.Request) bool {
	fmt.Println("Checking Request")

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		if strings.Join(v, "") == "hack" {
			fmt.Println("We have been hacked")
			return false
		}
	}

	return true
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	c, err := fcgi.NewClient("tcp", "php:9000")
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			if r.Method == "POST" {
				if !checkRequest(r) {
					http.NotFound(w, r)
					return
				}
			}
			php(wd, w, r, c)
		})
	log.Panic(http.ListenAndServe(":8080", nil))
}
