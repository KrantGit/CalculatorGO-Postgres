package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Input struct {
	FirstNumber  float64 `json:"first_number"`
	Sign         string  `json:"sign"`
	SecondNumber float64 `json:"second_number"`
}

type Output struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func calculatorHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		http.Error(w, "Not POST method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != io.EOF && err != nil {
		panic(err)
	}
	defer req.Body.Close()

	var input Input
	err = json.Unmarshal(body, &input)
	if err != nil {
		panic(err)
	}

	var result Output

	switch input.Sign {
	case "+":
		result.Result = input.FirstNumber + input.SecondNumber
	case "-":
		result.Result = input.FirstNumber - input.SecondNumber
	case "*":
		result.Result = input.FirstNumber * input.SecondNumber
	case "/":
		if input.SecondNumber == 0 {
			result.Error = "Division by zero"
		} else {
			result.Result = input.FirstNumber / input.SecondNumber
		}
	default:
		result.Error = "Unknown sign: " + input.Sign
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		panic(err)
	}

	insert_table(input, result)

}
