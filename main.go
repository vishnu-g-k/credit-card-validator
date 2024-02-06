package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type cardNumber struct {
	Number string `json:"number"`
}

type Response struct {
	Valid bool `json:"valid"`
}

func luhnAlgorithm(cardNumber string) bool {
	total := 0
	isSecondDigit := false

	if len(cardNumber) == 0 {
		return false
	}

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit
		isSecondDigit = !isSecondDigit
	}
	fmt.Println("total: ", total)
	return total%10 == 0
}

func validateCreditCard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside validate function")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}
	var cc cardNumber
	err := json.NewDecoder(r.Body).Decode(&cc)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	isValid := luhnAlgorithm(cc.Number)
	response := Response{Valid: isValid}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	fmt.Println("inside main")
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("routre: ", router)
	router.HandleFunc("/validate", validateCreditCard).Methods("POST")
	fmt.Println("after handle func")

	err := http.ListenAndServe("0.0.0.0:8080", router)
	fmt.Println("err: ", err)
	if err != nil {
		fmt.Println("****Error : ", err)
	}
}
