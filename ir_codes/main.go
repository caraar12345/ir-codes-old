package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("usage: %s file_prefix\n", args[0])
	}

	filePrefix := os.Args[1]

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatalln(err)
	}

	prefixedFiles := []fs.FileInfo{}

	for _, f := range files {
		if strings.HasPrefix(f.Name(), filePrefix) {
			prefixedFiles = append(prefixedFiles, f)
		}
	}

	for _, f := range prefixedFiles {
		fileByteArray, err := ioutil.ReadFile(f.Name())
		if err != nil {
			log.Fatalln(err)
		}

		stringHex := string(fileByteArray)

		hexByteArray, err := hex.DecodeString(stringHex)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("#", f.Name())
		fmt.Println(b64.StdEncoding.EncodeToString(hexByteArray))
	}
}
