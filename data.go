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

// StructScriptEscaper verify each string structure field and
// escapes special characters like "<" to become "&lt;".
// It escapes only five such characters: <, >, &, ' and ".
func StructScriptEscaper(structptr interface{}) {
	structEscape(structptr, escaper)
}

// StructCharEscaper works the same as ScriptEscaper, but escapes only character you need
func StructCharEscaper(structptr interface{}, oldstr string, newstr string){
	newEscaper := strings.NewReplacer(oldstr, newstr)
	structEscape(structptr, newEscaper)
}

// MapScriptEscaper verify each string value of map and
// escapes special characters like "<" to become "&lt;".
// It escapes only five such characters: <, >, &, ' and ".
func MapScriptEscaper(mapPtr *map[string]interface{}){
	mapEscape(mapPtr, escaper)
}



// StringScriptEscaper verify single string and
// escapes special characters like "<" to become "&lt;".
// It escapes only five such characters: <, >, &, ' and ".
func StringScriptEscaper(stringWithSyms *string){
	*stringWithSyms = escaper.Replace(*stringWithSyms)
}

func structEscape(structptr interface{}, replacer *strings.Replacer){
	t := reflect.ValueOf(structptr).Elem()
	recursiveStructEscape(t, replacer)
}

func recursiveStructEscape(t reflect.Value, replacer *strings.Replacer){
	fieldsLen := t.NumField()

	for i := 0; i < fieldsLen; i++ {

		if isStructType(t.Field(i).Type().String()){
			recursiveStructEscape(t.Field(i), replacer)
		}

		if isPointerType(t.Field(i).Type().String()){
			recursiveStructEscape(t.Field(i).Elem(), replacer)
		}

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

func isStructType(typeName string)bool{
	var stdTypes = []string{"map", "int", "float", "map", "string", "bool", "*"}

	for _, stdTypeIterator := range stdTypes {
		if strings.Contains(typeName, stdTypeIterator){
			return false
		}
	}
	return true
}

func isPointerType(typeName string)bool{
	if strings.Contains(typeName, "*"){
		return true
	}
	return false
}



