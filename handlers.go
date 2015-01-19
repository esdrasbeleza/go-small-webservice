package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Dao NotesDao
}

func CreateHandler(dao NotesDao) Handler {
	return Handler{Dao: dao}
}

func (handler *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	result := handler.Dao.GetAllNotes()
	json.NewEncoder(w).Encode(result)
}

func (handler *Handler) RegisterNote(w http.ResponseWriter, r *http.Request) {
	note := NewEmptyNote()
	jsonError := json.NewDecoder(r.Body).Decode(&note)

	if jsonError != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Printf("Inserting new note %s \n", note)
		insertError := handler.Dao.StoreNote(note)
		if insertError != nil {
			fmt.Fprintln(w, insertError)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func (handler *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["noteId"]
	result, err := handler.Dao.GetNoteById(noteId)

	if err != nil {
		// TODO: handle errors different of 404
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}
