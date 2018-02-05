package gowafp

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/tomasen/fcgi_client"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// AnalyzeRequest will analyze the request for malicious intent.
func AnalyzeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := bluemonday.UGCPolicy()
		r.ParseForm()
		for k, v := range r.Form {
			unSanitized := strings.Join(v, "")            // @TODO check this
			r.Form[k] = []string{p.Sanitize(unSanitized)} // @TODO check this
			// @TODO check if the input had malicious code and log it
		}
		next.ServeHTTP(w, r)
	})
}

// PhpHandler is a net/http Handler that starts the process for passing
// the request to PHP-FPM.
func PhpHandler(script string, protocol string, address string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env := make(map[string]string)
		env["SCRIPT_FILENAME"] = script

		fcgi, err := fcgiclient.Dial(protocol, address)
		defer fcgi.Close()

		if err != nil {
			log.Println("err:", err)
		}

		if r.Method == "POST" {
			phpPost(env, fcgi, w, r)

			return
		}

		phpGet(env, fcgi, w)
	})
}

// phpPost is called when the user submits a POST request to the website.
func phpPost(env map[string]string, f *fcgiclient.FCGIClient, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := f.PostForm(env, r.Form)
	if err != nil {
		log.Println("Post Err:", err)
	}
	phpProcessResponse(resp, w)
}

// phpGet is called when a user visits any page and submits a GET request to the
// website.
func phpGet(env map[string]string, f *fcgiclient.FCGIClient, w http.ResponseWriter) {
	resp, err := f.Get(env)
	if err != nil {
		log.Println("Get Err:", err)
	}
	phpProcessResponse(resp, w)
}

// phpProcessResponse is used by phpPost and phpGet to write the response back
// to the user's browser.
func phpProcessResponse(resp *http.Response, w http.ResponseWriter) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err:", err)
	}

	w.Write(content)
}
