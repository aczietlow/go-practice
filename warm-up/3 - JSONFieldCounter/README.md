# Problem: JSON Field Counter

*Objective:*

Write a Go program that reads a JSON string and counts the number of top-level fields in it.

*Description:*

You are given a JSON string that represents an object. Your task is to determine how many top-level fields exist in this JSON object.

*Example:*

Input: {"name": "John", "age": 30, "city": "New York"}

Output: 3

Input: {"data": {"id": 1, "type": "message"}, "status": "success", "code": 200}

Output: 3

*Requirements:*

The input string will always represent a valid JSON object.
The JSON object might contain nested objects or arrays, but you only need to count the top-level fields.
Consider using the encoding/json package in Go for parsing the JSON string.

*Hints:*

Unmarshal the JSON string into a map of string to empty interface (map[string]interface{}) and then count the number of keys in the map.

*Bonus:*

Extend the program to also display the names of the top-level fields.
Implement error handling for invalid JSON strings.
Once you've given it a try, let me know if you'd like hints, feedback, or a sample solution!

## Resources

[JSON and GO (offical docs)](https://go.dev/blog/json)
[encoding/json](https://pkg.go.dev/encoding/json)
[encoding/json marshalling](https://pkg.go.dev/encoding/json#Marshaler)
[json spec](https://www.json.org/json-en.html)
[JSON RFC](https://www.rfc-editor.org/rfc/rfc7159.html)

[chatGPT problem thread](https://chat.openai.com/c/ae852a73-233d-41d6-bc2b-5c32fea1222c)

