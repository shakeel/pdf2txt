package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shakeel/pdf2txt/pdf"
)

var (
	verbose     = flag.Bool("verbose", false, "verbose output")
	outFilepath = flag.String("out", "./out/", "`filepath` for writing output")
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of pdf2txt:\n\n")
	fmt.Fprintf(os.Stderr, "pdf2txt [flags] pdf-file ...\n\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func fatalf(format string, args ...interface{}) {
	log.SetFlags(0)
	log.Fatalf(format, args...)
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}

	err := os.MkdirAll(*outFilepath, 0644)
	if err != nil {
		fatalf("fatal error %v", err)
	}

	for _, in := range flag.Args() {
		reader, err := pdf.Open(in)
		if err != nil {
			fatalf("fatal error %v", err)
		}

		out := strings.Replace(filepath.Join(*outFilepath, filepath.Base(in)), ".pdf", ".txt", 1)
		writer, err := os.Create(out)
		if err != nil {
			fatalf("fatal error %v", err)
		}

		var b strings.Builder
		for i := 1; i <= reader.NumPage(); i++ {
			// Initialize y co-ordinate for the page
			y := 0.0
			for _, t := range reader.Page(i).Content().Text {
				// Check if we are on a new line
				if t.Y != y {
					y = t.Y
					b.WriteString("\n")
				}
				b.WriteString(t.S)
			}
		}
		fmt.Fprintf(writer, "%v\n", b.String())
		writer.Close()
	}
}
