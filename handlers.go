package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func MaskIndex(w http.ResponseWriter, r *http.Request) {
     masks := Masks{
           Mask{Number: "3016864576"},
           Mask{Number: "3016864576"},
           Mask{Number: "3016864576"},
     }     

    if err := json.NewEncoder(w).Encode(masks); err != nil {
        panic(err)
    }
}