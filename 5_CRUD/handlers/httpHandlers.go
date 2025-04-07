package handlers

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "./pages/rootPage.html")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Route not valid"))
}

func MethodNotAllowed(w http.ResponseWriter) {
	w.Write([]byte("Method not allowed"))
	w.WriteHeader(http.StatusMethodNotAllowed)
}
