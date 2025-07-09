// cmd/nftmetadataindexercore/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"nftmetadataindexercore/internal/nftmetadataindexercore"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := nftmetadataindexercore.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
