package main

import (
	"fmt"
	"net/url"
	"net/http"
	"os"
	"encoding/base64"
	"strings"
	"encoding/json"
	"io/ioutil"
	"strconv"

	. "github.com/inturn/go-helpers"
)

var WRAUrl string = "https://2ylflv45i7.execute-api.us-west-2.amazonaws.com/prod/WolframalphaQuery?input="

var googleUrl string = "https://www.googleapis.com/customsearch/v1?key=AIzaSyB20e2VDjrUebicIJkA4MFH4WO4b8cEzQY&cx=013676722247143124300:dazj-lelyfy&num=3"

func main() {

	if (os.Getenv("TERM_PROGRAM") != "iTerm.app") {
		fmt.Println("this only works in iTerm")
		os.Exit(2)
	}

	q := url.QueryEscape(strings.Join(os.Args[2:], " "))

	switch (os.Args[1]) {
	case "g":
		callGoogle(q)
	case "w":
		callWRA(q)
	}
}

func callGoogle(q string) {
	resp, err := http.Get(googleUrl + "&q=" + q)
	Check(err)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(contents))
}

func callWRA(q string) {
	resp, err := http.Get(WRAUrl + q)
	Check(err)
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)

	var responseItems []string

	json.Unmarshal([]byte(response), &responseItems)

	for imgKey, resultImg := range responseItems {
		imgResp, _ := http.Get(string(resultImg))
		defer resp.Body.Close()

		contents, _ := ioutil.ReadAll(imgResp.Body)
		filename := "/tmp/" + strconv.Itoa(imgKey) + "s_search_utility.gif"
		err = ioutil.WriteFile(filename, contents, 0644)

		printImg(filename)

		Check(err)
	}

}

// Thanks! https://github.com/trotha01/itermImage
func printImg(path string) {
	body, err := ioutil.ReadFile(path)
	Check(err)
	b64FileName := base64.StdEncoding.EncodeToString([]byte(path))
	b64FileContents := base64.StdEncoding.EncodeToString(body)
	fmt.Printf("\033]1337;File=name=%s;inline=1:%s\a\n", b64FileName, b64FileContents)
	defer os.Remove(path)
}