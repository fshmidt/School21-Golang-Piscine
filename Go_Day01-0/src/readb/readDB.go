package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

/*structs of db*/

type Item struct {
	Itemname  string `xml:"itemname" json:"ingredient_name"`
	Itemcount string `xml:"itemcount" json:"ingredient_count"`
	Itemunit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
}

type Cake struct {
	Name       string `xml:"name" json:"name"`
	Stovetime  string `xml:"stovetime" json:"time"`
	Ingredient []Item `xml:"ingredients>item" json:"ingredients"`
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Recipes []Cake   `xml:"cake" json:"cake"`
}

/*structs with their methods*/

type XML Recipes

func (p *XML) Read(file []byte) (Recipes, error) {
	err := xml.Unmarshal(file, p)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Recipes(*p), err
	}
	return Recipes(*p), err
}

type Json Recipes

func (p *Json) Read(file []byte) (Recipes, error) {
	err := json.Unmarshal(file, p)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Recipes(*p), err
	}
	return Recipes(*p), nil
}

type DBReader interface {
	Read(file []byte) (Recipes, error)
}

/*func of DBReader interface*/

func Rewrite(a DBReader, file []byte) []byte {
	var res []byte
	rcps, err := a.Read(file)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil
	}
	switch a.(type) {
	case *Json:
		res, err = xml.MarshalIndent(rcps, "", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return nil
		}
	case *XML:
		res, err = json.MarshalIndent(rcps, "", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return nil
		}
	default:
		break
	}
	return res
}

/*clearing from comments*/

func removeShit(file *[]byte) {
	s := string(*file)
	var x, y int

	r := []rune(s)
	var buf []rune
	for i := range r {
		if i > 0 && r[i] == '/' && r[i-1] == '/' {
			x = i
		}
		if x > 0 && r[i] == '\n' {
			y = i
			buf = append(r[0:x-1], r[y:]...)
			x = 0
		}
	}

	if len(buf) > 0 {
		newS := string(buf)
		*file = []byte(newS)
	}
}

/*checking format*/

func fileFormat(filePath string) string {
	if utf8.RuneCountInString(filePath) < 4 {
		fmt.Println("File path is too short")
		os.Exit(3)
	}
	if utf8.RuneCountInString(filePath) > 4 && filePath[utf8.RuneCountInString(filePath)-5:] == ".json" {
		return "json"
	} else if filePath[utf8.RuneCountInString(filePath)-4:] == ".xml" {
		return "xml"
	} else {
		fmt.Println("Wrong file format")
		os.Exit(4)
		return "0"
	}
}

func main() {
	var f bool
	var filePath string
	var res []byte

	/*recovery*/

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Unknown panic happend, but it's recovered already ^__^ :", err)
		}
	}()

	/*flags*/

	flag.BoolVar(&f, "f", false, "display path")
	flag.Parse()
	if f {
		filePath = os.Args[2]

		/*reading format & opening file*/

		format := fileFormat(filePath)
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("No such file")
			os.Exit(3)
		}
		//myStruct := new(Recipes)
		switch format {

		/*xml*/

		case "xml":
			myStruct := new(XML)
			res = Rewrite(myStruct, file)
		/*json*/

		case "json":
			removeShit(&file)
			myStruct := new(Json)
			res = Rewrite(myStruct, file)
		default:
			break
		}
		fmt.Printf("%s\n", res)
	} else {

		/*forgot to use flag for args*/

		fmt.Println("Use '-f' flag for passing path to Args.")
		os.Exit(5)
	}
}
