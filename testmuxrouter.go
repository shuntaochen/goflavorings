package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

var routes map[string][]interface{}

func testmux() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	spa := spaHandler{staticPath: "static", indexPath: "index.html"}
	router.PathPrefix("/static").Handler(spa)

	router.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("abc"))
	})

	routes = map[string][]interface{}{"/a": {routeAHandler, "get"}, "/b": {routeBHandler, "get"}}
	for k, v := range routes {
		mylog(k, v)
		router.HandleFunc(k, defh) //v.(func(w http.ResponseWriter, r *http.Request)))
	}
	mylog("listening on 8000")
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func defh(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/a" {
	}
	if r.URL.Path == "/b" {

	}
	if routes[r.URL.Path][1] != strings.ToLower(r.Method) {
		w.Write([]byte("Invalid http verb"))
		return
	}
	myfilter(routes[r.URL.Path][0], w, r)(w, r)

}

func myfilter(v interface{}, w http.ResponseWriter, r *http.Request) func(w http.ResponseWriter, r *http.Request) {
	//filter w,r
	return v.(func(w http.ResponseWriter, r *http.Request))
}

func mylog(entries ...interface{}) {
	for _, entry := range entries {
		fmt.Println("v:", entry)
	}
}

func routeAHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("saa"))
}

func routeBHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bbb"))
}

func testentry() {
	mylog(2, 3, 4)
	var routes map[string]interface{} = map[string]interface{}{"/a": routeAHandler, "/b": ""}
	for k, v := range routes {
		mylog(k, v)
		mylog(reflect.TypeOf(k), reflect.TypeOf(v))

	}
}
