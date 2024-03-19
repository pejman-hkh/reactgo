package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pejman-hkh/gdp/gdp"
	"github.com/pejman-hkh/gdp/gox"
)

func main() {
	dir := os.Args[1]

	files, _ := ioutil.ReadDir(dir)

	for _, file := range files {
		name := file.Name()
		len := len(name)

		if name[(len-3):len] == "gox" {
			fmt.Printf("%s\n", name)
			goxFile, _ := os.ReadFile(dir + name)
			document := gdp.Default(string(goxFile))

			out := gox.ToGo(&document)

			ioutil.WriteFile(dir+name[0:len-1], []byte(out[:]), 0644)
		}

	}
}
