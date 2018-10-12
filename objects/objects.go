package objects

import (
	"io"
	"log"
	"os"
	"net/http"
	"strings"
)

const STORAGE_ROOT = "I:\code\storage\objects"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method

	if m == http.MethodPut {
		put(w, r)
		return
	}

	if m == http.MethodGet {
		get(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func put(w http.ResponseWriter, r *http.Request) {
	objectName := strings.Split(r.URL.EscapedPath(), "/")[2]
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" + objectName)

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()

	io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
  objectName := strings.Split(r.URL.EscapedPath(), "/")[2]
  f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" + objectName)
  
  if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
  }
  
  defer f.Close()

  io.Copy(w, f)
}