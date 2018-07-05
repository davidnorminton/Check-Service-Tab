package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	domain    = "https://www.cs-catering-equipment.co.uk/"
	testUrl   = domain + "blizzard-h600wh-white-single-door-upright-refrigerator"
	backupUrl = domain + "interlevin-sc381-glass-door-merchandiser"
	tabClass  = "tab-2"
	pyMail    = "/home/dave/go/tabsDownEmail.py"
	pyPath    = "/usr/bin/python"
	href      = "<a href='{{url}}'>{{url}}</a><br />"
)

func elemLookup(body []byte, find string) bool {
	return strings.Contains(string(body), find)
}

func getPage(url string) *http.Response {

	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error: Downloading url")
	}

	return response
}

func readBody(response *http.Response) []byte {

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error: Reading html")
	}

	return body

}

func sendEmail(theUrl string) {

	link := strings.Replace(href, "{{url}}", theUrl, 2)
	cmd := exec.Command(pyPath, pyMail, link)

	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error: Starting python email script")
	}

	fmt.Println(string(output))

}

func main() {
	response := getPage(testUrl)

	body := readBody(response)

	if !elemLookup(body, tabClass) {
		sendEmail(testUrl)
	}

}
