package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"rsc.io/pdf"
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

	reader, err := pdf.Open(flag.Arg(0))
	if err != nil {
		fatalf("file error %s", flag.Arg(0))
	}

	for i := 1; i <= reader.NumPage(); i++ {
		for _, t := range reader.Page(i).Content().Text {
			fmt.Println(t)
		}
	}
	fmt.Println("--")

}
