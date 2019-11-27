package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mikey-wotton/indexer"
	"go.uber.org/zap"
)

//go:generate go build -o indexer.exe

func main() {
	dirPath := flag.String("directory", "./", "The Path to index, default current folder")
	outPath := flag.String("output", "index.html", "Output index, default index.html")
	flag.Parse()

	var logger, _ = zap.NewProduction()

	if !strings.HasSuffix(*outPath, ".html") {
		*outPath = *outPath + ".html"
	}

	err := indexer.CreateIndex(*dirPath, *outPath)
	if err != nil {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	} else {
		log.Print(fmt.Sprintf("Success! Your index file was created at ./%s\n", *outPath))
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
