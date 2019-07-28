package lib

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

var typemap = map[string]string{
	"smallint":  "int8",
	"mediumint": "int16",
	"int":       "int32",
	"bigint":    "int64",
	"float":     "float32",
	"double":    "float64",
	"varchar":   "string",
	"boolean":   "bool",
	"timestamp": "time.Time",
	"date":      "time.Time",
}

func BuildStruct(table *Table) string {
	var modelStr string
	modelStr = fmt.Sprintf("type %v struct { \n", camel2underscore((*table).Name))

	for _, info := range (*table).Columns {
		var pkstr string = ""
		if info.IsPK() {
			pkstr = ";primary_key"
		}
		modelStr += fmt.Sprintf("\t%v %v `gorm:%v%v` \n", underscore2camel(info.Name), typemap[info.Datatype], info.Name, pkstr)
	}

	modelStr += "} \n"

	return modelStr
}

func camel2underscore(name string) string {
	buffer := bytes.NewBufferString("")
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteByte('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}

	return buffer.String()
}

func underscore2camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
