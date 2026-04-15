package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type data map[string]any

func main() {
	js := `{"data": {"id": 1, "type": "message"}, "status": "success", "code": 200}`

	d := parseJson(js)
	fmt.Printf("JSON Object has %d top level fields\n", countTopLevelFields(d))

	tlf := getTopLevelFields(d)
	fmt.Printf("Top level fields are: %v", tlf)
}

func parseJson(jsonString string) data {
	d := data{}
	err := json.Unmarshal([]byte(jsonString), &d)
	if err != nil {
		log.Panic(err)
	}
	return d
}

func countTopLevelFields(parsedData data) int {
	return len(parsedData)
}

func getTopLevelFields(parsedData data) []string {
	fields := []string{}
	for i, _ := range parsedData {
		fields = append(fields, i)
	}

	return fields
}
