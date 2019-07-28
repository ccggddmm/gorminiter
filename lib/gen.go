package lib

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

//todo 字段映射表可配置
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

func BuildFile(table *Table) {

}

func BuildStruct(table *Table) string {
	var buf bytes.Buffer
	buf.WriteString("type ")
	buf.WriteString(camel2underscore((*table).Name))
	buf.WriteString(" struct { \n")

	for _, info := range (*table).Columns {
		var pkstr string = ""
		if info.IsPK() {
			pkstr = ";primary_key"
		}
		tmpstr := fmt.Sprintf("\t%v %v `gorm:\"column:%v%v\"` \n",
			underscore2camel(info.Name), typemap[info.Datatype], info.Name, pkstr)
		buf.WriteString(tmpstr)
	}
	buf.WriteString("} \n")

	return buf.String()
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
