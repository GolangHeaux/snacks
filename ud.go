//ud.go gets the most popular urban dictionary term and is for a video which covers how to make api requests and use json with golang

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//UrbanDictionary holds the JSON Response
type UrbanDictionary struct {
	List []struct {
		Definition string `json:"definition"`
		Example    string `json:"example"`
	} `json:"list"`
} //UrbanDictionary

//main demonstrates how to unmarshal json on the internet by using the undocumented urban dictionary API
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ud <term>")
		os.Exit(3)
	} //end-if no arguments were given

	term := strings.Join(args, " ")
	request, _ := http.Get("http://api.urbandictionary.com/v0/define?term=" + url.QueryEscape(term))
	thejson, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()

	ud := UrbanDictionary{}
	json.Unmarshal(thejson, &ud)
	if len(ud.List) == 0 {
		fmt.Println("no exact match found for " + term)
		os.Exit(3)
	} //end-if no results found

	fmt.Printf(strings.ToUpper(term)+": %s\nex: %s\n", ud.List[0].Definition, ud.List[0].Example)
} //main
