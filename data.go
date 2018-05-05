package fieldracer

import (
	"reflect"
	"strings"
)

var escaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&#34;", // "&#34;" is shorter than "&quot;".
)

// ScriptEscaper verify each string structure field and
// escapes special characters like "<" to become "&lt;".
// It escapes only five such characters: <, >, &, ' and ".
func StructScriptEscaper(structptr interface{}) {
	structEscape(structptr, escaper)
}

// CharEscaper works the same as ScriptEscaper, but escapes only character you need
func StructCharEscaper(structptr interface{}, oldstr string, newstr string){
	newEscaper := strings.NewReplacer(oldstr, newstr)
	structEscape(structptr, newEscaper)

}

func MapScriptEscaper(mapPtr *map[string]interface{}){
	mapEscape(mapPtr, escaper)
}

func structEscape(structptr interface{}, replacer *strings.Replacer){
	t := reflect.ValueOf(structptr).Elem()
	fieldsLen := t.NumField()

	for i := 0; i < fieldsLen; i++ {
		if t.Field(i).Type().String() == "string" {
			t.Field(i).SetString(replacer.Replace(t.Field(i).String()))
		}
	}
}

func mapEscape(mapPtr *map[string]interface{}, replacer *strings.Replacer){
	for k, v := range *mapPtr{
		if strValue, ok := v.(string); ok{
			(*mapPtr)[k] = interface{}(replacer.Replace(strValue))
		}
	}
}


