# TIL

Things I learned from this

## Things I don't understand and need to do a deeper dive into

The entire Marshalling and Unmarshalling.

### Some questions that tripped me up

#### Type conversations:

Why did the examples want to encode things into a slice of bytes? 

```go
b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
```

I know that the Marshall function `func Marshal(v interface{}) ([]byte, error)` accepts any type so long as it has the implements the Marshaler Interface and by extension the MarshalJSON() method, but why is []byte advantageous in this example?

```go
// Marshaler is the interface implemented by types that
// can marshal themselves into valid JSON.
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}
```

#### Interface

Let's be honest, only have a tepid understanding of interfaces in the first place, so this shouldn't be surprising at all. I really don't understand the use of empty interface throughout the JSON encode/decode examples and code base.

```go
// Decoding arbitrary data
// Consider this JSON data, stored in the variable b:

b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

// Without knowing this data’s structure, we can decode it into an interface{} value with Unmarshal:

var f interface{}
err := json.Unmarshal(b, &f)
```

Another example 

```go 
func Marshal(v any) ([]byte, error) {
	e := newEncodeState()

	err := e.marshal(v, encOpts{escapeHTML: true})
	if err != nil {
		return nil, err
	}
	buf := append([]byte(nil), e.Bytes()...)

	encodeStatePool.Put(e)

	return buf, nil
}
```

For additional context - any is an alias of interface{}

```go
// builtin.go
// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}
```

#### DataStructures

I want to take json of an unknown structure, and map it to SOME data structure in go that can be transversed in order to work on a solution to the problem at of counting parent fields. I am a bit lost at how to convert the json encoded data into anything useful. There's additionally the challenge of picking the best data structure.

#### What ever the fuck this is

```go
	var out interface{}
    err := json.Unmarshal(jsonData, &out)
```

I still get confused about when to use the `&` and when not. Is this by reference, copy, or to lookup a pointer... Memory management is a thing. 