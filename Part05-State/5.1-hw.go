package main

import (
	
	"encoding/json"
	"fmt"
	"os"

)

func main() {


	type loaction struct {
		Name string  `json:"name"`
		Lat float64  `json:"latitude"`
		Long float64 `json:"longitude"`

	}

	loactions := []loaction{
		{"New York", 40.7128, -74.0060},
		{"London", 51.5074, -0.1278},
		{"Paris", 48.8566, 2.3522},
	}

	bytes, err := json.Marshal(loactions)
	exit0nError(err)

	fmt.Println(string(bytes))
}



func exit0nError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
