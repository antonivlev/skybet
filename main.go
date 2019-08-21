package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antonivlev/skybet/roulette"
)

func main() {
	rou := roulette.Roulette{
		RedNumbers:   []int{1, 2, 3, 4},
		BlackNumbers: []int{5, 6, 7, 8},
	}
	fmt.Printf("roulette: \n%+v\n\n", rou)

	// playBetOnSingleNumber(&rou, 3, 13.50)
	// playBetOnSingleNumber(&r, 3, 13.50)

	// playColourBet(&r, "red", 13.50)
	// playColourBet(&r, "red", 13.50)
	// playColourBet(&r, "black", 13.50)

	http.HandleFunc("/", catchAll)
	http.HandleFunc("/bet/single", func(w http.ResponseWriter, r *http.Request) {
		playBetOnSingleNumber(&rou, 3, 13.50)
		fmt.Fprintf(w, "inner func")
	})

	// Serve
	port := "8080"
	log.Println("Serving on " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	log.Println("all")
	fmt.Fprintf(w, "Use GET /bet with some params")
}

func handleBet(w http.ResponseWriter, r *http.Request) {
	log.Println("bet")

	fmt.Fprintf(w, "this is a bet")
}
