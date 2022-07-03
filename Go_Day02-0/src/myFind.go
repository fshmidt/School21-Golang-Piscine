package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"os"
)

type Flags struct {
	f, d, sl *bool
	ext      *string
}

func main() {
	var fl Flags

	fl.f = flag.Bool("f", false, "show file")
	fl.d = flag.Bool("d", false, "show directories")
	fl.sl = flag.Bool("sl", false, "show symlink")
	fl.ext = flag.String("ext", "", "show extension")
	flag.Parse()

	if *fl.ext != "" && !*fl.f {
		fmt.Println("Flag error. Use -f for using -ext")
		os.Exit(1)
	}
	if !*fl.f && !*fl.d && !*fl.sl {
		*fl.f, *fl.d, *fl.sl = true, true, true
	}

	filepath.Walk(flag.Arg(0), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if *fl.sl && info.Mode()&fs.ModeSymlink != 0 {
			fmt.Print(path + "->")
			sLink, err := filepath.EvalSymlinks(path)
			if err != nil {
			fmt.Println("[broken]")
			} else {
				fmt.Println(sLink)
			}
		} else if *fl.d && info.IsDir() {
			fmt.Println(path)
		} else if *fl.f && info.Mode()>>9 == 0 {
			if *fl.ext != "" {
				if filepath.Ext(path) == "."+*fl.ext {
					fmt.Println(path)
				}
			} else {
			fmt.Println(path)
			}
		}
		return nil
	})
}
