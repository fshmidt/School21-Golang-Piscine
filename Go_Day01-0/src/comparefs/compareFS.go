package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

/*checking format*/

func checkFileFormat(filePath string) {
	if utf8.RuneCountInString(filePath) < 4 {
		fmt.Println("File path is too short")
		os.Exit(3)
	}
	if utf8.RuneCountInString(filePath) > 3 && filePath[utf8.RuneCountInString(filePath)-4:] == ".txt" {
		return
	} else {
		fmt.Println("Wrong file format")
		os.Exit(4)
	}
}

func main() {
	var oldFilePath, newFilePath string

	/*recovery*/

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Unknown panic happend, but it's recovered already ^__^ :", err)
		}
	}()

	/*flags*/

	f1 := flag.NewFlagSet("f1", flag.ContinueOnError)
	old := f1.Bool("old", false, "take old data")
	neww := f1.Bool("new", false, "take new data")
	if len(os.Args) > 3 {
		f1.Parse(os.Args[1:])
		f1.Parse(os.Args[3:])
	} else {
		fmt.Println("Not enough args.")
		os.Exit(1)
	}
	if *old && *neww {
		if os.Args[1] != "--old" && os.Args[1] != "-old" {
			fmt.Println("Wrong order of args. Older is first.")
			os.Exit(2)
		} else if os.Args[1] == "-old" {
			fmt.Println("Use '--', not '-'.")
			os.Exit(2)
		}
		oldFilePath = os.Args[2]
		newFilePath = os.Args[4]

		/*checking format & reading entire 1st file*/

		checkFileFormat(oldFilePath)
		oldFile, errO := ioutil.ReadFile(oldFilePath)
		if errO != nil {
			fmt.Println("No such file")
			os.Exit(3)
		}

		/*making array from 1st file*/

		s := string(oldFile)
		arr := strings.Split(s, "\n")

		/*using bufio for reading line by line for 2nd file*/

		newFile, err := os.Open(newFilePath)
		if err != nil {
			fmt.Println("Can't open 2nd file")
			os.Exit(5)
		}
		sc := bufio.NewScanner(newFile)

		/*checking strings one by one one way*/

		for sc.Scan() {
			var num int
			var line string
			for _, line = range arr {
				if line == sc.Text() {
					num++
				}
			}
			if num == 0 {
				fmt.Println("ADDED ", sc.Text())
			}
		}

		/*closing file to reuse all the memory again*/

		newFile.Close()

		/*doing the same using alredy having resources*/

		oldFile, errO = ioutil.ReadFile(newFilePath)
		if errO != nil {
			fmt.Println("Can't read second file")
			os.Exit(6)
		}
		s = string(oldFile)
		arr = strings.Split(s, "\n")
		newFile, err = os.Open(oldFilePath)
		if err != nil {
			fmt.Println("Can't open first file")
			os.Exit(7)
		}
		defer newFile.Close()
		sc = bufio.NewScanner(newFile)

		/*all variables now contain vise versa data
		now we can make reversed checking of strings*/

		for sc.Scan() {
			var num int
			var line string
			for _, line = range arr {
				line = line + "\n"

				if line[0:utf8.RuneCountInString(line)-1] == sc.Text() {
					num++
				}
			}
			if num == 0 {
				fmt.Println("REMOVED ", sc.Text())
			}
		}
	}
}
