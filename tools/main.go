package main

import (
	//	"fmt"
	"github.com/chiradeep/go-nitro/config/lb"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

type Config struct {
	Package    string
	TfTitle    string
	TfName     string
	TfId       string
	StructName string
	Fields     map[string]string
}

func getFieldNames() map[string]string {
	result := make(map[string]string)
	t := reflect.TypeOf(&lb.Lbvserver{}).Elem()
	for index := 0; index < t.NumField(); index++ {
		field := t.Field(index)

		name := strings.ToLower(field.Name)
		typ := strings.Title(field.Type.Name())
		if typ != "" {
			result[name] = typ
		}
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
	funcMap := template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"neq": func(x, y interface{}) bool {
			return x != y
		},
	}
	writer, err := os.Create(filepath.Join("netscaler", "resource_lb.go"))
	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("resource.tmpl", "provider.tmpl"))
	err = t.ExecuteTemplate(writer, "resource.tmpl", cfg)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
	writer, err = os.Create(filepath.Join("netscaler", "provider.go"))
	err = t.ExecuteTemplate(writer, "provider.tmpl", cfg)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}
