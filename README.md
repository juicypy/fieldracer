# fieldracer

juicypy/fieldracer is a golang library for easy parsing structure fields
It helps to escape special script characters demanding not a lot of code from developer 
using only "reflect" and "strings" go std libraries

## Getting Started

From your GOPATH:

```bash
go get github.com/juicypy/fieldracer
```
## example

```go
package main

import (
	"github.com/juicypy/fieldracer"
	"fmt"
)

type structWithScpriptedFields struct {
	FieldOne       string
	FieldTwo       string
	NotStringField int64
}

func main(){

	withScript := structWithScpriptedFields{"<danger script>", "not a script", 42}

	fieldracer.StructScriptEscaper(&withScript) //call it with pointer of your structure

	fmt.Printf("Escaped structure: %+v", withScript)
	//-> Escaped structure: &{FieldOne:&lt;danger script&gt; FieldTwo:not a script NotStringField:42}
}
```
# it even works with nested structures or it's pointers

```go
type structWithScpriptedFields struct {
	FieldOne       string
	FieldTwo       string
	NotStringField int64
	NestedStruct *nestedStructureWithScripts
}

type nestedStructureWithScripts struct{
	FieldThree string
	FieldFour string
}

func main(){

	nestedStruct := nestedStructureWithScripts{"user data", "<one more danger script>"}

	withScript := structWithScpriptedFields{"<danger script>", "not a script", 42,
		&nestedStruct}

	fieldracer.StructScriptEscaper(&withScript) //call it with pointer of your structure

	fmt.Printf("Escaped structures: %+v \n Nested structure: %+v", withScript, withScript.NestedStruct)
	//-> Escaped structures: {FieldOne:&lt;danger script&gt; FieldTwo:not a script NotStringField:42 NestedStruct:0xc42000a060}
	//   Nested structure: &{FieldThree:user data FieldFour:&lt;one more danger script&gt;}
}
```

# it's also works with map[string]interface{} in the same way
# you can escape only character you need using: 
```go
  fieldracer.StructCharEscaper(&withScript, "*", "(")
  //and fields with "*" character will have "(" in their place
```
# and you can escape single string
```go
  someScriptedString := "<one more danger script>"
  fieldracer.StringScriptEscaper(&someScriptedString)
```


