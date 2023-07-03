package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// url = https://xkcd.com/571/info.0.json


type Comic struct{
	Num int
	Year string
	SafeTitle string `json:"safe_title"`
	Img string
	Title string
	Transcript string
}

var comics []Comic

func main(){
	var comic Comic
	for i:= 1 ; i<10; i++{
		link := "http://xkcd.com/"+fmt.Sprint(i)+"/info.0.json"
		resp, err := http.Get(link)
		if resp.StatusCode != http.StatusOK{
			resp.Body.Close()
			break
		}
		if err != nil{
			resp.Body.Close()
			log.Fatal(err)
		}
		err = json.NewDecoder(resp.Body).Decode(&comic)
		if err != nil{
			resp.Body.Close()
			log.Fatal(err)
		}

		comics = append(comics, comic)
	}

	searchTerm := os.Args[1]
	for _,v := range comics{
		if strings.Contains(v.Title,searchTerm) {
			fmt.Printf("%s\t%s",v.Img,v.Transcript)
		}
	}

}