package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func main() {
	router := mux.NewRouter().StrictSlash(true)
	staticDir := "/static/"

	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	spa := spaHandler{staticPath: "static", indexPath: "index.html"}
	router.PathPrefix("/spa").Handler(spa)

	router.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("abc"))
	})
	vipSite := VipSiteHandlers{}

	routes = map[string][]interface{}{api("/login"): {vipSite.Login, "get"}}
	for k := range routes {
		router.HandleFunc(k, defh)
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

	if routes[r.URL.Path][1] != strings.ToLower(r.Method) {
		w.Write([]byte("Invalid http verb"))
		return
	}
	//verify jwt verity for certain routes requiring authentication,
	myfilter(routes[r.URL.Path][0], w, r)(w, r)

}

func myfilter(v interface{}, w http.ResponseWriter, r *http.Request) func(w http.ResponseWriter, r *http.Request) {
	//filter w,r

	return v.(func(w http.ResponseWriter, r *http.Request))
}

func api(path string) string {
	return "/api" + path
}
