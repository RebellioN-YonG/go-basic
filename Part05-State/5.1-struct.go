package main

import (
	"os"
	"fmt"
	"encoding/json"
)

func main() {

	var ccp struct {
		name string
		age  int
	}

	ccp.name = "wennie"
	ccp.age = 71

	fmt.Println(ccp.name, ccp.age)

	fmt.Println(ccp)

	type location struct {
		lat float64
		lng float64
	
	}

	var spirit location

	spirit.lat = 37.422

	spirit.lng = -122.084

	fmt.Println(spirit.lat, spirit.lng)

	var opportunity location


	opportunity.lat = 19.8964


	opportunity.lng = -128.694
	fmt.Println(spirit, opportunity)

	// slice-struct
	type location2 struct {
		name string
		lat  float64
		long float64
	}

	location2s := []location2{
		{name:"Tokyo", lat: 35.6895, long: 139.6917},
		{name:"Delhi", lat: 28.6139, long: 77.2090},
	}

	fmt.Println(location2s)

	// encoding struct to json
	type lo struct{
		Lat, Long float64
	}

	cur := lo{37.422, -122.084}
	bytes, err := json.Marshal(cur)
	exit0nError(err)

	fmt.Println(string(bytes))

	// customize json tags
	type location3 struct {
		Name string `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longtitude"`
	}

	cur2 := location3{"Tokyo", 35.6895, 139.6917,}
	bytes2, err := json.Marshal(cur2)
	exit0nError(err)

	fmt.Println(string(bytes2))

	}	

	func exit0nError(err error)  {
		if err!= nil {
			fmt.Println(err)
			os.Exit(1)
		}	

		
}