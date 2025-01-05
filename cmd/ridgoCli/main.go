package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lj3954/ridgo/pkg/any"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		log.Fatalln("Usage: ridgo <release> [language] [arch]")
	}
	lang := "English"
	if len(os.Args) == 3 {
		lang = os.Args[2]
	}
	arch := "x86_64"
	if len(os.Args) == 4 {
		arch = os.Args[3]
	}
	url, err := any.Get(os.Args[1], lang, arch)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(url.URL)
}
