package main

import (
	"github.com/tomasen/fcgi_client"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"regexp"
)

// PhpHandler is a net/http Handler that starts the process for passing
// the request to PHP-FPM.
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

// PhpPost is called when the user submits a POST request to the website.
func PhpPost(env map[string]string, f *fcgiclient.FCGIClient, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := f.PostForm(env, r.Form)
	if err != nil {
		log.Println("Post Err:", err)
	}
	PhpProcessResponse(resp, w)
}

// PhpGet is called when a user visits any page and submits a GET request to the
// website.
func PhpGet(env map[string]string, f *fcgiclient.FCGIClient, w http.ResponseWriter) {
	resp, err := f.Get(env)
	if err != nil {
		log.Println("Get Err:", err)
	}
	PhpProcessResponse(resp, w)
}

// PhpProcessResponse is used by PhpPost and PhpGet to write the response back
// to the user's browser.
func PhpProcessResponse(resp *http.Response, w http.ResponseWriter) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err:", err)
	}

	w.Write(content)
}

// AnalyzeRequest will analyze the request for malicious intent.
func AnalyzeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Analyzing Request")
		r.ParseForm()
		log.Println(r.Form["secret"])
		re, _ := regexp.Compile("script")
		for k, v := range r.Form {
			if re.MatchString(strings.Join(v, "")) {
				log.Println("Attack Detected")
			}
			//r.Form[k] = []string{"hii"}
			r.Form[k] = v
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", AnalyzeRequest(PhpHandler()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
