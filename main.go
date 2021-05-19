package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	idg, err := NewMonotonicIDGenerator("ids")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		id, err := idg.Generate(req.Context(), "scope-a")
		if err != nil {
			log.Default().Print(err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.Write([]byte(string(id)))
	})

	fmt.Println("Listening on port 8001")
	err = http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Default().Print(err)
	}
}
