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

// ScriptEscaper receives structure pointer, verify each string field and
// escapes special characters like "<" to become "&lt;".
// It escapes only five such characters: <, >, &, ' and ".
func ScriptEscaper(structptr interface{}) {
	escape(structptr, escaper)
}

// CharEscaper works the same as ScriptEscaper, but escapes only character you need
func CharEscaper(structptr interface{}, oldstr string, newstr string){
	newEscaper := strings.NewReplacer(oldstr, newstr)
	escape(structptr, newEscaper)

}

func escape(structptr interface{}, replacer *strings.Replacer){
	t := reflect.ValueOf(structptr).Elem()
	fieldsLen := t.NumField()

	for i := 0; i < fieldsLen; i++ {
		if t.Field(i).Type().String() == "string" {
			t.Field(i).SetString(replacer.Replace(t.Field(i).String()))
		}
	}
}

func EqualFieldsSetter(fromstruct interface{}, tostruct interface{}){

}

func NotEqualFieldsSetter(fromstruct interface{}, tostruct interface{}){

}




