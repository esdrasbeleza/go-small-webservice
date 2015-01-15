package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func ListNotes(w http.ResponseWriter, r *http.Request) {
	result := []Note{}
	collection.Find(nil).All(&result)
	json.NewEncoder(w).Encode(result)
}

func RegisterNote(w http.ResponseWriter, r *http.Request) {
	note := NewEmptyNote()
	jsonError := json.NewDecoder(r.Body).Decode(&note)

	if jsonError != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Printf("Inserting new note %s \n", note)
		insertError := collection.Insert(note)
		if insertError != nil {
			fmt.Fprintln(w, insertError)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["noteId"]
	query := bson.M{"_id": bson.ObjectIdHex(noteId)}

	result := Note{}
	err := collection.Find(query).One(&result)

	if err != nil {
		// TODO: handle errors different of 404
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}
