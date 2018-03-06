package main

import (
	"conf2go/config"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	D        map[string]interface{}
	path     = flag.String("config", "./config.json", "")
	confTmpl *template.Template
)

type SimpleField struct {
	Name string
	Type string
}

type StructField struct {
	Name         string
	SimpleFields []SimpleField
}

type Data struct {
	SimpleFields []SimpleField
	StructFields []StructField
	ConfPath     string
}

func InitConfig(way string) (err error) {
	D = make(map[string]interface{})

	*path, err = filepath.Abs(way)
	if err != nil {
		panic(fmt.Errorf("Error get abs path:%v", err))
	}

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		return fmt.Errorf("InitConfig -> couldn't read config file %s -> %s", path, err)
	}
	if err = json.Unmarshal(data, &D); err != nil {
		return fmt.Errorf("InitConfig -> couldn't unmarshal json -> %s", err)
	}

	return
}

func main() {
	flag.Parse()
	if err := InitConfig(*path); err != nil {
		fmt.Printf("Err while initializing config -> %s", err)
	}

	fmap := template.FuncMap{"goName": goName}
	confTmpl = template.Must(template.New("config.tmpl").Funcs(fmap).ParseFiles("config.tmpl"))

	confgo := generatePackage()
	defer confgo.Close()
	c := getData(D)
	fmt.Println("SIMPLE FIELDS")
	for _, v := range c.SimpleFields {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("STRUCT FIELDS")
	for _, v := range c.StructFields {
		fmt.Printf("%+v\n", v)
	}
	if err := confTmpl.Execute(confgo, c); err != nil {
		fmt.Printf("While executing template -> %s\n", err)
	}
	e := config.Init()
	if e != nil {
		fmt.Println(e)
	}
	//fmt.Println(config.C.Bd.A)
}

func generatePackage() *os.File {
	pathDir := filepath.Join(".", "config")
	if err := os.MkdirAll(pathDir, 0777); err != nil {
		fmt.Printf("while creating config directory -> %s", err)
	}

	pathFile := filepath.Join(pathDir, "config.go")
	conf, err := os.Create(pathFile)
	if err != nil {
		fmt.Printf("while creating conf.go -> %s", err)
	}
	return conf
}

func getData(rawFields map[string]interface{}) (data Data) {
	var simpleFields []SimpleField
	var structFields []StructField

	data.ConfPath = *path

	for k, v := range rawFields {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			sf := StructField{
				Name: k,
			}
			for j, b := range v.(map[string]interface{}) {
				smf := SimpleField{
					Name: j,
					Type: reflect.TypeOf(b).String(),
				}
				sf.SimpleFields = append(sf.SimpleFields, smf)
			}
			structFields = append(structFields, sf)
			continue
		}
		smf := SimpleField{
			Name: k,
			Type: reflect.TypeOf(v).String(),
		}
		simpleFields = append(simpleFields, smf)
	}
	data.StructFields = structFields
	data.SimpleFields = simpleFields
	return
}

func goName(jName string) (name string) {
	firstToUpper := func(s string) string {
		runes := []rune(s)
		return strings.ToUpper(string(runes[0])) + string(runes[1:])
	}
	words := strings.Split(jName, "_")
	for _, w := range words {
		name += firstToUpper(w)
	}
	return
}

func printConf(rawFields map[string]interface{}) {
	for k, v := range rawFields {

		fmt.Println(k, v)
		fmt.Println("TYPES: ", reflect.TypeOf(k), reflect.TypeOf(v))
		fmt.Println("VALUES: ", reflect.ValueOf(k), reflect.ValueOf(v))

		if reflect.TypeOf(v).Kind() == reflect.Map {
			fmt.Println("FOUND A MAP")
			for j, b := range v.(map[string]interface{}) {
				fmt.Println(j, b)
			}
			fmt.Println("______________________________________")
		}
	}
}
