//Marshalling JSON in go

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Example from the doc @ https://godoc.org/encoding/json#Marshal
//Marshalling. -- converting a golang value (v interface{}) to JSON format
type color_values struct {
	Id     int      `json:"id"`
	Colors []string `json:"colors"`
}

func main() {
	c := color_values{
		Id:     432, //fields need to be available be
		Colors: []string{"blue", "green", "majenta", "gray"},
	}

	c_json, e := json.Marshal(c)
	if e != nil {
		fmt.Println(c_json)
	}
	os.Stdout.Write(c_json)

	//Unmarshall
	//Check: json-to-go website for quick cheat to see which data-structure should the JSon be converted to in case its not clear
	var c2 color_values
	err := json.Unmarshal(c_json, &c2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c2)
}
