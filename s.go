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

	"github.com/fatih/color"
	. "github.com/inturn/go-helpers"
)

var WRAUrl string = "https://2ylflv45i7.execute-api.us-west-2.amazonaws.com/prod/WolframalphaQuery?input="

var googleUrl string = "https://www.googleapis.com/customsearch/v1?key=AIzaSyB20e2VDjrUebicIJkA4MFH4WO4b8cEzQY&cx=013676722247143124300:dazj-lelyfy"

type GoogleResponse struct {
	Items []struct {
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
		Title   string `json:"title"`
	} `json:"items"`
}

func main() {

	if len(os.Args) > 2 {

		q := url.QueryEscape(strings.Join(os.Args[2:], " "))

		switch (os.Args[1]) {
		case "g":
			callGoogle(q)
		case "w":
			if (os.Getenv("TERM_PROGRAM") != "iTerm.app") {
				fmt.Println("this only works in iTerm 3")
				os.Exit(2)
			}
			callWRA(q)
		}
	}

	fmt.Println(`
		https://github.com/jwdeitch/search-cli


		Supported search providers:
		google (g):
		  -y=[int]		limit search to N years back
		  -n=[int]		Return N results (max 10)

		Wolfram Alpha (w)


		Example usage:
		s w time in israel
		s g cat day care nyc -n=5 -y=1 (limit search results to 1 year back, and only return 5 results)

		`)
}

func parseFlags(q string) string {
	q = q + " "

	// number of results
	if strings.Contains(q, "-n%3D") {
		numPosition := strings.Index(q, "-n%3D")
		num := q[numPosition + 5:numPosition + 7]
		num = strings.Replace(num, " ", "", 1)
		numInt, _ := strconv.Atoi(num)
		if numInt > 10 {
			numInt = 10
		}

		stringToRemove := q[numPosition:numPosition + 7]
		q = strings.Replace(q, stringToRemove, "", 1)
		q = q + "&num=" + strconv.Itoa(numInt)
	} else {
		q = q + "&num=3"
	}

	// year limit
	if strings.Contains(q, "-y%3D") {
		yearPosition := strings.Index(q, "-y%3D")
		year := q[yearPosition + 5:yearPosition + 6]
		stringToRemove := q[yearPosition:yearPosition + 6]
		q = strings.Replace(q, stringToRemove, "", 1)
		q = q + "&dateRestrict=y[" + year + "]"
		fmt.Println(q)
	}

	return q
}

func callGoogle(q string) {
	q = parseFlags(q)
	resp, err := http.Get(googleUrl + "&q=" + q)
	defer resp.Body.Close()
	Check(err)
	response, _ := ioutil.ReadAll(resp.Body)

	var responseItems GoogleResponse

	json.Unmarshal([]byte(response), &responseItems)

	//fmt.Println(responseItems)
	for _, responseItem := range responseItems.Items {
		fmt.Println()
		color.Cyan(responseItem.Title)
		fmt.Println(responseItem.Snippet)
		color.Green(responseItem.Link)
		fmt.Println()
	}

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