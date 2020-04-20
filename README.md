## Example
```go 
package main

//go:generate typedjson -interface Data -structs *Foo,*Bar
type Data interface {
	typedjson(*DataTyped) string
}

type Foo struct {
	Name  string
	Value int
}

type Bar struct {
	Name1   string
	Another []string
}
```

After generation you will be able to use your `Foo` and `Bar in typed manner:

```go 
func main() {
	foo := &Bar{
		Name1: "asd",
	}
	d := DataTyped{}

	d.Data = foo
	b, err := json.MarshalIndent(&d, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

	d1 := DataTyped{}
	err = json.Unmarshal(b, &d1)
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(d, d1) {
		panic("not equal")
	}

}
```
