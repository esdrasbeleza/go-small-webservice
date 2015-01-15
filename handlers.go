package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	result := []Note{}
	collection.Find(nil).All(&result)
	json.NewEncoder(w).Encode(result)
}

func ShowNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId, _ := strconv.Atoi(vars["noteId"])
	query := bson.M{"id": noteId}

	result := Note{}
	err := collection.Find(query).One(&result)

	if err == nil {
		json.NewEncoder(w).Encode(result)
	} else {
		log.Println(err)
		w.WriteHeader(404)
	}
}
