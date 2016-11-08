package main

import (
	//	"fmt"
	"github.com/chiradeep/go-nitro/config/lb"
	"html/template"
	"log"
	"os"
	"reflect"
)

type Config struct {
	Package    string
	TfTitle    string
	TfName     string
	TfId       string
	StructName string
	Fields     map[string]reflect.Type
}

func getFieldNames() map[string]reflect.Type {
	result := make(map[string]reflect.Type)
	t := reflect.TypeOf(&lb.Lbvserver{}).Elem()
	for index := 0; index < t.NumField(); index++ {
		field := t.Field(index)

		name := field.Name
		typ := field.Type
		result[name] = typ
	}
	return result
}

func main() {
	cfg := Config{Package: "lb",
		TfName:     "lbvserver",
		TfTitle:    "Lbvserver",
		TfId:       "lbvserverName",
		StructName: "Lbvserver",
		Fields:     getFieldNames()}
	t := template.Must(template.ParseFiles("resource.tmpl"))
	err := t.Execute(os.Stdout, cfg)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}
