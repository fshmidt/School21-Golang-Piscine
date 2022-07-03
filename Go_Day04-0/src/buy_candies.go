package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Parameters struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type Resp struct {
	Change int    `json:"change"`
	Tnahks string `json:"thanks"`
}

type PReq struct {
	Err string `json:"error"`
}

func sp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		d := json.NewDecoder(r.Body)
		defer r.Body.Close()
		p := &Parameters{}
		if err := d.Decode(p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		money := p.Money
		candyType := p.CandyType
		candyCount := p.CandyCount

		if money <= 0 || candyCount <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			O := &PReq{}
			O.Err = "Use normal numbers for candies and money smart-ass"
			z := json.NewEncoder(w).Encode(O)
			if z != nil {
				log.Fatal(z)
			}
		} else {
			var cPrice int
			if candyType == "CE" {
				cPrice = 10
			} else if candyType == "AA" {
				cPrice = 15
			} else if candyType == "NT" {
				cPrice = 17
			} else if candyType == "DE" {
				cPrice = 21
			} else if candyType == "YR" {
				cPrice = 23
			}
			if cPrice == 0 {
				w.WriteHeader(http.StatusBadRequest)
				O := &PReq{}
				O.Err = "Wrong candy name!"
				z := json.NewEncoder(w).Encode(O)
				if z != nil {
					log.Fatal(z)
				}
			} else if money >= candyCount*cPrice {
				w.WriteHeader(http.StatusCreated)
				O := &Resp{}
				O.Tnahks = "Thank you!"
				O.Change = money - candyCount*cPrice
				z := json.NewEncoder(w).Encode(O)
				if z != nil {
					log.Fatal(z)
				}
			} else if money < candyCount*cPrice {
				w.WriteHeader(http.StatusPaymentRequired)
				O := &PReq{}
				need := candyCount*cPrice - money
				O.Err = fmt.Sprintf("You need %d more money!", need)
				z := json.NewEncoder(w).Encode(O)
				if z != nil {
					log.Fatal(z)
				}
			}
		}
	default:
		break
	}
}

func main() {
	http.HandleFunc("/buy_candy", sp)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
