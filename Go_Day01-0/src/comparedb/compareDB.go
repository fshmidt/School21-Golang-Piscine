package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/r3labs/diff"
	"io/ioutil"
	"os"
	"unicode/utf8"
	"strings"
)

/*structs of db*/

type Item struct {
	Itemname  string `xml:"itemname" json:"ingredient_name" diff:"ingedient, identifier"`
	Itemcount string `xml:"itemcount" json:"ingredient_count" diff:"item_count"`
	Itemunit  string `xml:"itemunit" json:"ingredient_unit,omitempty" diff:"unit"`
}

type Recipe struct {
	Name       string `xml:"name" json:"name" diff: "recipe, indentifier"`
	Stovetime  string `xml:"stovetime" json:"time" diff: "time"`
	Ingredient []Item `xml:"ingredients>item" json:"ingredients" diff:"ingredients, indentifier"`
}

type Recipes struct {
	Recipes []Recipe   `xml:"cake" json:"cake"`
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

func pairs(p []string) string {
	pairs:= make([]string, len(p)/2 + len(p)%2)
	var a,b int
	for a = len(pairs) -1; b < len(p)&^1; b, a = b + 2, a -1 {
		pairs[a] = fmt.Sprintf("%s %s", p[b], p[b+1])
	}
	if a == 0 {
		pairs[a] = p[b]
	}
	return strings.Join(pairs, " for ")
}

func main() {
	var oldFilePath, newFilePath string
	var oldRes, newRes []byte
	var changelog diff.Changelog
	var ups error
	var jsonP *Json
	var xmlP *XML

	/*recovery*/

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Unknown panic happend, but it's recovered already ^__^ :", err)
		}
	}()

	/*flags*/

	f1 := flag.NewFlagSet("f1", flag.ContinueOnError)
	old  := f1.Bool("old", false, "take old data")
	neww := f1.Bool("new", false, "take new data")

	if len(os.Args) > 3{
		f1.Parse(os.Args[1:])
		f1.Parse(os.Args[3:])
	} else {
		fmt.Println("Not enough args.")
		os.Exit(1)
	}
	if *old && *neww {
		if os.Args[1] != "--old" && os.Args[1] != "-old"{
			fmt.Println("Wrong order of args. Older is first.")
			os.Exit(2)
		} else if os.Args[1] == "-old"{
			fmt.Println("Use '--', not '-'.")
			os.Exit(2)
		}
		oldFilePath = os.Args[2]
		newFilePath = os.Args[4]

		/*reading format & opening files*/

		oldFormat := fileFormat(oldFilePath)
		oldFile, errO := ioutil.ReadFile(oldFilePath)
		if errO != nil {
			fmt.Println("No such file")
			os.Exit(3)
		}
		newFile, errN := ioutil.ReadFile(newFilePath)
		if errN != nil {
			fmt.Println("No such file")
			os.Exit(3)
		}
		switch oldFormat {

		/*xml*/

		case "xml":
			myOldStruct := new(XML)
			oldRes = Rewrite(myOldStruct, oldFile)
			removeShit(&newFile)
			myNewStruct := new(Json)
			newRes = Rewrite(myNewStruct, newFile)
			xmlP = myOldStruct
			jsonP = myNewStruct

			/*json*/

		case "json":
			fmt.Println("here")
			removeShit(&oldFile)
			myOldStruct := new(Json)
			oldRes = Rewrite(myOldStruct, oldFile)
			myNewStruct := new(XML)
			newRes = Rewrite(myNewStruct, newFile)
			xmlP = myNewStruct
			jsonP = myOldStruct
		default:
			break
		}
		_, _ = oldRes, newRes
		differ, err := diff.NewDiffer(diff.DisableStructValues())
		if err != nil {
			fmt.Println("No such file")
				os.Exit(5)
			}
		changelog, ups = differ.Diff(xmlP, jsonP)
		if ups != nil {
				fmt.Println("No such file")
				os.Exit(4)
			}
		for _, change := range changelog{
			a := change.Path
			switch change.Type {
			case diff.CREATE:
				fmt.Printf("ADDED %s\n", pairs(a))
			case diff.UPDATE:
				fmt.Printf("CHANGED %s - %s instead of %s\n", pairs(a), change.To, change.From)
			case diff.DELETE:
				switch n:= len(a) -1; a[n] {
				case "unit":
					a = append(a, change.From.(string))
				case "ingredient":
					a = a[:n]
				}
				fmt.Printf("REMOVED %s\n", pairs(a))
			}
		}
	} else {

		/*forgot to use flag for args*/

		fmt.Println("Use '--old' & '--new' flags for passing path to Args.")
		os.Exit(5)
	}
}
