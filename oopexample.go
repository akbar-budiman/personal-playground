package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getSavedObjectsProps(w http.ResponseWriter, r *http.Request) {
	savedObjects := []Geometry{
		Rect{GeometryObject: GeometryObject{Id: 1}, Width: 2, Length: 3},
		Circle{GeometryObject: GeometryObject{Id: 2}, Radius: 10},
	}

	var response []GeometryProps
	for _, obj := range savedObjects {
		response = append(
			response,
			GeometryProps{
				Area:      obj.Area(),
				Perimeter: obj.Perimeter(),
			},
		)
	}

	fmt.Println(r.Method + " " + r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
