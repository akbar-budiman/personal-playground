package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Document struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	LastUpdate time.Time `json:"lastUpdate"`
}

var documents []Document

func RegisterRouter() {
	documents = []Document{
		{Id: 1, Title: "Daftar Belanja", Content: "Telur, Roti tawar, Laundry", LastUpdate: time.Now()},
		{Id: 2, Title: "Agenda", Content: "Cari Kontrakan", LastUpdate: time.Now()},
	}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/documents", getAllDocument).Methods("GET")
	myRouter.HandleFunc("/documents", addDocument).Methods("POST")
	myRouter.HandleFunc("/documents/{id}", getCertainDocument).Methods("GET")
	myRouter.HandleFunc("/documents/{id}", removeCertainDocument).Methods("DELETE")

	myRouter.HandleFunc("/geometry", addGeometryObj).Methods("POST")
	myRouter.HandleFunc("/geometry/{id}", getCertainGeometryObj).Methods("GET")
	myRouter.HandleFunc("/geometry/{id}/properties", getCertainGeometryProps).Methods("GET")

	myRouter.HandleFunc("/saved-obj/properties", getSavedObjectsProps).Methods("GET")

	myRouter.HandleFunc("/goroutine", RunGoRoutineExampleController)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + " " + r.RequestURI)
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getAllDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + " " + r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(documents)
}

func addDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + " " + r.RequestURI)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newDocument Document
	json.Unmarshal(reqBody, &newDocument)
	newDocument.LastUpdate = time.Now()

	for _, data := range documents {
		if data.Id == newDocument.Id {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	documents = append(documents, newDocument)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(documents)
}

func getCertainDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + " " + r.RequestURI)
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStrToSearch := vars["id"]
	idIntToSearch, _ := strconv.ParseInt(idStrToSearch, 10, 0)

	for _, data := range documents {
		if data.Id == int(idIntToSearch) {
			json.NewEncoder(w).Encode(data)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func removeCertainDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + " " + r.RequestURI)

	vars := mux.Vars(r)
	idStrToSearch := vars["id"]
	idIntToSearch, _ := strconv.ParseInt(idStrToSearch, 10, 0)

	foundIndex := -1
	for index, data := range documents {
		if data.Id == int(idIntToSearch) {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		documents = append(documents[:foundIndex], documents[foundIndex+1:]...)
		w.WriteHeader(http.StatusNoContent)
	}

}

func addGeometryObj(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + " " + r.RequestURI)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newObj Rect
	json.Unmarshal(reqBody, &newObj)

	SetData(newObj.Id, string(reqBody[:]))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newObj)
}

func getCertainGeometryObj(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + " " + r.RequestURI)
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStrToSearch := vars["id"]
	idIntToSearch, _ := strconv.Atoi(idStrToSearch)

	data := GetData(idIntToSearch)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getCertainGeometryProps(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method + " " + r.RequestURI)
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStrToSearch := vars["id"]
	idIntToSearch, _ := strconv.Atoi(idStrToSearch)

	data := GetData(idIntToSearch)

	if data != nil {
		response := GeometryProps{
			// Area:      data.Area(),
			// Perimeter: data.Perimeter(),
			Area:      AreaCounter{data}.Count(),
			Perimeter: PerimeterCounter{data}.Count(),
		}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
