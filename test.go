package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

// Thanks! https://github.com/trotha01/itermImage
func main() {
	path := "/Users/jordan1/Desktop/MSP798224cgg39g58hha470000547e9ih4gg38b9c4.gif"
	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading image %s. Error: %s", path, err.Error())
	}
	b64FileName := base64.StdEncoding.EncodeToString([]byte(path))
	b64FileContents := base64.StdEncoding.EncodeToString(body)
	fmt.Printf("\033]1337;File=name=%s;inline=1:%s\a\n", b64FileName, b64FileContents)
}