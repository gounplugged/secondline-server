package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome fool!")
}

func MaskIndex(w http.ResponseWriter, r *http.Request) {
     masks := Masks{
           Mask{Number: "3016864576", CountryCode: "1"},
           Mask{Number: "3016864576", CountryCode: "1"},
           Mask{Number: "3016864576", CountryCode: "1"},
      	   Mask{Number: "475932921", CountryCode: "+32"},
      	   Mask{Number:	"475932921", CountryCode: "+32"},
      	   Mask{Number:	"475932921", CountryCode: "+32"},
     }

    if err := json.NewEncoder(w).Encode(masks); err != nil {
        panic(err)
    }
}