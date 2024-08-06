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

const (
	AppName = "epub-to-md"
)

var (
	fileFlag    = flag.String("file", "", "Path to the file to read")
	outFlag     = flag.String("out", "", "Path to the output file")
	versionFlag = flag.Bool("version", false, "Print version and exit")
)

var (
	version = "dev" // set by ldflags at build time
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(AppName + " " + version)
		return
	}

	rc, err := epub.OpenReader(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer rc.Close()

	// There may be multiple rootfiles, although typically there is only one.
	book := rc.Rootfiles[0]

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
	fmt.Fprintln(output, "")
	fmt.Fprintln(output, "> Created with "+AppName+" "+version)
	fmt.Fprintln(output, "> [https://github.com/dhcgn/epub-to-md](https://github.com/dhcgn/epub-to-md)")
	fmt.Fprintln(output, "")

	for _, item := range book.Spine.Itemrefs {
		r, err := item.Open()
		if err != nil {
			log.Fatal(err)
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
