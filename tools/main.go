package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

type Config struct {
	Package        string
	TfTitle        string
	TfName         string
	TfID           string
	NsID           string
	StructName     string
	Fields         map[string]string
	BindingName    string
	BindingPkg     string
	BindingType    string
	BoundType      string
	KeyFields      map[string]interface{} //important fields and their values for documentation and testing
	KeyFieldsBound map[string]interface{} //important fields and their values for documentation and testing for bound type if any
}

var (
	i = flag.String("i", "", "The input JSON Schema file.")
	d = flag.String("d", "", "The NS identifier for the object: e.g., name, policyname, etc")
	b = flag.String("b", "", "The JSON schema file for the binding if any")
	n = flag.String("n", "", "The name for the HCL field that specifies the binding")
	k = flag.String("k", "", "JSON string mapping key fields to values for testing and documentation")
	K = flag.String("K", "", "If this resource is always bound to another, this parameter provides a JSON string mapping key fields to values for the Bound type")
)

func parseSchema(inputFile string) *Schema {
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read the input file with error ", err)
		return nil
	}
	schema, err := Parse(string(b))

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse the input JSON schema with error ", err)
		return nil
	}
	//fmt.Printf("Parse schema: %v\n", schema.Properties["appflowlog"].Readonly)
	return &schema.Schema

}

func getFieldNames(obj interface{}) map[string]string {
	result := make(map[string]string)
	t := reflect.TypeOf(obj).Elem()
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

func getFieldNamesFromSchema(schema Schema) map[string]string {
	result := make(map[string]string)
	for key, value := range schema.Properties {
		fieldName := strings.Join(strings.Split(strings.ToLower(key), "_"), "")
		typ := getPrimitiveTypeName(value.Type)
		readonly := value.Readonly
		if typ != "" && !readonly {
			result[fieldName] = strings.Title(typ)
		}
	}
	return result
}

func getConfigFromSchema(pkg string, schema Schema, keyFieldsJSON string, boundKeyFieldsJSON string, nsID string) *Config {
	fields := getFieldNamesFromSchema(schema)
	cfg := Config{Package: pkg,
		TfName:      schema.ID,
		TfTitle:     strings.Title(schema.ID),
		TfID:        schema.ID + "Name",
		NsID:        nsID,
		StructName:  strings.Title(schema.ID),
		Fields:      fields,
		BindingName: "",
	}
	keyFieldValues := make(map[string]interface{})
	boundKeyFieldValues := make(map[string]interface{})
	cfg.KeyFields = keyFieldValues
	if keyFieldsJSON != "" {
		err := json.Unmarshal([]byte(keyFieldsJSON), &keyFieldValues)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to parse keyfield JSON ", err)
			return &cfg
		}
		cfg.KeyFields = keyFieldValues
	}
	if boundKeyFieldsJSON != "" {
		err := json.Unmarshal([]byte(boundKeyFieldsJSON), &boundKeyFieldValues)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to parse bound keyfield JSON ", err)
			return &cfg
		}
		cfg.KeyFieldsBound = boundKeyFieldValues
	}
	return &cfg
}

func getConfig(pkg string, tfName string, structName string, configObj interface{}) *Config {
	cfg := Config{Package: pkg,
		TfName:      tfName,
		TfTitle:     structName,
		TfID:        tfName + "Name",
		StructName:  structName,
		Fields:      getFieldNames(configObj),
		BindingName: "",
	}
	return &cfg
}

func main() {
	flag.Parse()

	funcMap := template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"neq": func(x, y interface{}) bool {
			return x != y
		},
		"isInt": func(x interface{}) bool {
			switch x.(type) {
			case int:
				return true
			case float64:
				return true
			default:
				return false
			}
			return false
		},
		"isPresent": func(x map[string]interface{}) bool {
			return len(x) > 0
		},
	}
	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("resource.tmpl", "provider.tmpl", "resource_test.tmpl"))

	if *i == "" {
		log.Fatal("No input schema file provided")
	}
	schema := parseSchema(*i)
	pkg := filepath.Base(filepath.Dir(*i))
	keyFields := ""
	boundKeyFields := ""
	nsID := "name"
	if *k != "" { //if key fields are provided
		keyFields = *k
	}
	if *K != "" { //if bound key fields are provided
		boundKeyFields = *K
	}
	if *d != "" {
		nsID = *d
	}
	cfg := getConfigFromSchema(pkg, *schema, keyFields, boundKeyFields, nsID)
	if *n != "" && *b != "" { //if binding is required
		bindingSchema := parseSchema(*b)
		cfg.BindingName = *n
		cfg.BindingPkg = filepath.Base(filepath.Dir(*b))
		cfg.BindingType = strings.Title(strings.Join(strings.Split(bindingSchema.ID, "_"), ""))
		cfg.BoundType = strings.Title(strings.Split(bindingSchema.ID, "_")[0])
	}

	writer, err := os.Create(filepath.Join("netscaler", "resource_"+schema.ID+".go"))
	err = t.ExecuteTemplate(writer, "resource.tmpl", *cfg)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
	writer, err = os.Create(filepath.Join("netscaler", "provider.go"))
	err = t.ExecuteTemplate(writer, "provider.tmpl", *cfg)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}

	writer, err = os.Create(filepath.Join("netscaler", "resource_"+schema.ID+"_test.go"))
	err = t.ExecuteTemplate(writer, "resource_test.tmpl", *cfg)
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}
