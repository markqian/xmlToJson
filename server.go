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

type Order struct {
	Id string `xml:"id" json:"id"`
	Data string `xml:"data" json:"data"`
	CreatedAt string `xml:"createdAt" json:"createdAt"`
	UpdatedAt string `xml:"updatedAt" json:"updatedAt"`
}

type OrderList struct {
	Orders []Order `xml:"order"`
}


// process the given order
func processXMLOrder(order *Order, finish chan bool) {
	order.Data = strings.ToUpper(order.Data)
	// signal that this order has finished processing
	finish <- true
}


// Process xml and convert to anonymous JSON format
func processXML(data string) string {
	orderList := OrderList{}

	err := xml.Unmarshal([]byte(data), &orderList)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err.Error()
	}

	// check if each goroutine has finished
	finish := make(chan bool)

	for index, _ := range orderList.Orders {
		go processXMLOrder(&orderList.Orders[index], finish)
		<- finish 
	}

	// convert to byte array
	jsonByteArray, _ := json.MarshalIndent(&orderList.Orders, "", "    ")

	jsonString := "{}"
	if len(orderList.Orders) > 0 {
		// convert to string
		jsonString = "{" + string(jsonByteArray)[1:len(jsonByteArray) - 1] + "}"
	}
	
	// get json string
	return jsonString
}
	

func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, _ := ioutil.ReadAll(r.Body)

		result := processXML(string(body))

		fmt.Fprintf(w, "%s", result)	
    } 
}

func main() {
	http.HandleFunc("/process", processHandler)  
	log.Fatal(http.ListenAndServe(":8080", nil))
}

