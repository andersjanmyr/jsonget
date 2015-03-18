# JsonGet

A small module to get values from a JSON structure without having
to create the Go types for them.


## Install

```
$ go get github.com/andersjanmyr/jsonget
```

## Usage

```json
// Example json
{
  "name":"Wednesday",
  "age":6,
  "color": {
    "red": "#ff0000",
    "green": "#00ff00",
    "blue": "#0000ff"
  },
  "parents":["Gomez","Morticia"],
  "pets":[{"name": "Fido"}, {"name": "Misse"}]
}
```


```go
value, err := JsonGet(data, "name")
// value = "Wednesday"

value, _ := JsonGet(data, "parents")
// value = []interface{}{"Gomez", "Morticia"}

value, _ := JsonGet(data, "color")
// value := map[string]interface{}{"red": "#ff0000", "green": "#00ff00", "blue": "#0000ff"}

value, _ := JsonGet(data, "color.red")
// value = "#ff0000"

value, _ := JsonGet(data, "parents.1")
// value = "Morticia"

value, _ := JsonGet(data, "pets.1.name")
// value := "Misse"

value, _ := JsonGet(data, "pets.*.name")
// value = []interface{}{"Fido", "Misse"}
```
