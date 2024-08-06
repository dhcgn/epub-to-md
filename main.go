package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/taylorskalyo/goreader/epub"
)

var (
	fileFlag = flag.String("file", "", "Path to the file to read")
	outFlag  = flag.String("out", "", "Path to the output file")
)

func main() {
	flag.Parse()

	rc, err := epub.OpenReader(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer rc.Close()

	// The rootfile (content.opf) lists all of the contents of an epub file.
	// There may be multiple rootfiles, although typically there is only one.
	book := rc.Rootfiles[0]

	// Print book title.
	// In main.go, replace the existing code after the book declaration with:

	var output io.Writer

	if *outFlag == "" {
		output = os.Stdout
	} else {
		file, err := os.Create(*outFlag)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		output = file
	}
	// Print book title.
	fmt.Fprintln(output, "# "+book.Title)

	// List the IDs of files in the book's spine.
	for _, item := range book.Spine.Itemrefs {
		fmt.Fprintln(output, item.ID)
		r, err := item.Open()
		if err != nil {
			panic(err)
		}

		converter := md.NewConverter("", true, nil)

		html, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		markdown, err := converter.ConvertString(string(html))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(output, markdown)
	}
}
