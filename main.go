package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antonivlev/skybet/bets"
	"github.com/gorilla/schema"

	"github.com/antonivlev/skybet/roulette"
)

func main() {
	rou := roulette.Roulette{
		RedNumbers:   []int{1, 2, 3, 4},
		BlackNumbers: []int{5, 6, 7, 8},
	}
	fmt.Printf("configured roulette: \n%+v\n\n", rou)

	// Register endpoints. Could create another roulette and register another set of
	// endpoints for it.
	http.HandleFunc("/", catchAll)
	handleBet(&rou, "/betSingle", bets.PlayBetOnSingleNumber)
	handleBet(&rou, "/betColour", bets.PlayColourBet)

	// Serve
	port := "8080"
	log.Println("Serving on " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Handles unknown requests. Would be nice to include some instructions in its response,
// for example listing all the endpoints, and their expected arguments
func catchAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use GET /betSingle or /betColour with some params. e.g:\n")
	fmt.Fprintf(w, "\nhttp://localhost:8080/betSingle?money=12.45&number=2\n")
}

// Registers a betting func to an endpoint. Takes care of parsing args from request.
// At the moment, parses args from URL, could be modified to accept a json payload.
func handleBet(rou *roulette.Roulette, urlPath string, handler bets.BettingFunc) {
	http.HandleFunc(urlPath, func(w http.ResponseWriter, r *http.Request) {
		// put the bet params into a bets.BetArgs struct
		var args bets.BetArgs
		paramsMap := r.URL.Query()
		// the gorilla library is used only here for parameter parsing
		err := schema.NewDecoder().Decode(&args, paramsMap)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			fmt.Fprintf(w, "\n-----------\n")
			fmt.Fprintf(w, "Are you passing correctly named and typed params? Will be decoded into:\n\n%#v", bets.BetArgs{})

			fmt.Println("error:", err)
			return
		}
		// TODO: should check here if the params are semantically ok, e.g. colour is contained in roulette
		// or should this be done in the individual betting funcs? for example a colour bet func could
		// tell you you shouldn't pass a number through

		// if successful, play the bet with these args
		win, outNum, outCol := handler(rou, args)
		// communicate result to user
		fmt.Fprintf(w, urlPath+"\n\n")
		fmt.Fprintf(w, "You bet %4.2f and got: %d %s\n", args.Money, outNum, outCol)
		fmt.Fprintf(w, "Winnings: %4.2f\n", win)
	})
}
