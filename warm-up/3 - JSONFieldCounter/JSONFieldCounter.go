package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := `{"name": "John", "age": 30, "city": "New York"}`

	// encodes json data.
	jsonData, err := json.Marshal(b)

	if err != nil {
		fmt.Println(err)
	}

	output(jsonData)
}

func output(jsonData []byte) {
	// JSON encoded data can be unmarshalls to fill an object.
	// If b contains valid JSON that fits in [insert defined object 'm'], after the call err will be nil and the data from b will have been stored in the struct m.
	// If the
	var out interface{}
	err := json.Unmarshal(jsonData, &out)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", out)
}
