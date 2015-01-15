package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

func AddNote(w http.ResponseWriter, r *http.Request) {
	note := Note{Id: bson.NewObjectId(), Added: time.Now(), Updated: time.Now()}
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Printf("Inserting new note %s \n", note)
		err2 := collection.Insert(note)
		if err2 != nil {
			fmt.Fprintln(w, err2)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func ShowNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["noteId"]
	oid := bson.ObjectIdHex(noteId)
	query := bson.M{"_id": oid}

	result := Note{}
	err := collection.Find(query).One(&result)

	if err == nil {
		json.NewEncoder(w).Encode(result)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}
}
