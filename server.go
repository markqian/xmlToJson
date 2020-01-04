package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"net/http"
)

var c chan string

type Order struct {
	Id string `xml:"id" json:"id"`
	Data string `xml:"data" json:"data"`
	CreatedAt string `xml:"createdAt" json:"createdAt"`
	UpdatedAt string `xml:"updatedAt" json:"updatedAt"`
}

type OrderList struct {
	Orders []Order `xml:"order"`
}

// Process xml and convert to anonymous JSON format
func processXML(data string) {
	v := OrderList{}

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		c <- err.Error()
		return
	}

	for index, element := range v.Orders {
		// index is the index where we are
		// element is the element from Orders for where we are
		v.Orders[index].Data = strings.ToUpper(element.Data)
	}

	jsonByteArray, err := json.MarshalIndent(&v.Orders, "", "    ")

	jsonString := "{}"
	if len(v.Orders) > 0 {
		// convert to string
		jsonString = "{" + string(jsonByteArray)[1:len(jsonByteArray) - 1] + "}"
	}
	c <- jsonString
}
	

func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := ioutil.ReadAll(r.Body)

		go processXML(string(body))

		result := <- c

		fmt.Fprintf(w, "%s", result)	
    } 
}

func main() {
	c = make(chan string)
	http.HandleFunc("/process", processHandler)  
	log.Fatal(http.ListenAndServe(":8080", nil))
}

