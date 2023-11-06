package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Act as json object, don't need to unmarshal it...
	jsonResponse := fakeJsonResponse()

	var out map[string]any
	// aka: var out map[string]interface{} // I just think the above is cleaner

	err := json.Unmarshal([]byte(jsonResponse), &out)

	if err != nil {
		fmt.Println("failed to unmarshal json")
	}

	fmt.Printf("%#v\n", out)

	count := len(out)

	fmt.Println("Top level fields - ", count)
}

func fakeJsonResponse() string {
	// Just pretend we made an http request to a REST API and handled all the good shit here.

	//jsonString := `{"name": "John", "age": 30, "city": "New York"}`
	jsonString := `{"data": {"id": 1, "type": "message"}, "status": "success", "code": 200}`
	return jsonString
}
