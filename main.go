package main

import (
	"github.com/tomasen/fcgi_client"
	"io/ioutil"
	"log"
	"net/http"
)

func PhpHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env := make(map[string]string)
		env["SCRIPT_FILENAME"] = "/app/index.php"

		fcgi, err := fcgiclient.Dial("tcp", "php:9000")
		defer fcgi.Close()

		if err != nil {
			log.Println("err:", err)
		}

		if r.Method == "POST" {
			PhpPost(env, fcgi, w, r)

			return
		}

		PhpGet(env, fcgi, w)
	})
}

func PhpPost(env map[string]string, f *fcgiclient.FCGIClient, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := f.PostForm(env, r.Form)
	if err != nil {
		log.Println("Post Err:", err)
	}
	PhpProcessResponse(resp, w)
}

func PhpGet(env map[string]string, f *fcgiclient.FCGIClient, w http.ResponseWriter) {
	resp, err := f.Get(env)
	if err != nil {
		log.Println("Get Err:", err)
	}
	PhpProcessResponse(resp, w)
}

func PhpProcessResponse(resp *http.Response, w http.ResponseWriter) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err:", err)
	}

	w.Write(content)
}

func AnalyzeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Analyzing Request")
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", AnalyzeRequest(PhpHandler()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
