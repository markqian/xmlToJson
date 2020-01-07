package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
    "unicode"

)

func TestEmptyXML(t *testing.T) {

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/process", strings.NewReader("<orderList></orderList>")) // step 1

	processHandler(res, req) 
	content, _ := ioutil.ReadAll(res.Body)
	expected := "{}"
	if string(content) != expected { 
		t.Errorf("Expected %s, got %s", expected, string(content))
	}
}

func stripSpaces(str string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            // if the character is a space, drop it
            return -1
        }
        // else keep it in the string
        return r
    }, str)
}

func TestXMLWithValidFormat(t *testing.T) {

	data := `<orderList>
	<order>
			<id>aeffb38f-a1a0-48e7-b7a8-2621a2678534</id>
			<data>first_Order_Data</data>
			<createdAt>0001-01-01T00:00:00Z</createdAt>
			<updatedAt>0001-01-01T00:00:00Z</updatedAt>
		</order>
		<order>
			<id>beffb38f-b1a0-58e7-c7a8-3621a2678534</id>
			<data>second_Order_Data</data>
			<createdAt>0001-01-01T00:00:00Z</createdAt>
			<updatedAt>0001-01-01T00:00:00Z</updatedAt>
		</order>
	</orderList>`

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/process", strings.NewReader(data)) 

	processHandler(res, req) 
	content, _ := ioutil.ReadAll(res.Body)
	// strip spaces to compare string.
	expected := stripSpaces(`{
		{
			"id": "aeffb38f-a1a0-48e7-b7a8-2621a2678534",
			"data": "FIRST_ORDER_DATA",
			"createdAt": "0001-01-01T00:00:00Z",
			"updatedAt": "0001-01-01T00:00:00Z"
		},
		{
			"id": "beffb38f-b1a0-58e7-c7a8-3621a2678534",
			"data": "SECOND_ORDER_DATA",
			"createdAt": "0001-01-01T00:00:00Z",
			"updatedAt": "0001-01-01T00:00:00Z"
		}
	}`)


	if stripSpaces(string(content)) != expected { 
		t.Errorf("Expected %s, got %s", expected, string(content))
	}
}

func TestXMLWithIvalidFormat(t *testing.T) {

	data := `<orderList>
	<Order>
			<id>aeffb38f-a1a0-48e7-b7a8-2621a2678534</id>
			<data>first_Order_Data</data>
			<createdAt>0001-01-01T00:00:00Z</createdAt>
			<updatedAt>0001-01-01T00:00:00Z</updatedAt>
		</order>
	</orderList>`

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/process", strings.NewReader(data)) 

	processHandler(res, req) 
	content, _ := ioutil.ReadAll(res.Body)
	// strip spaces to compare string.
	expected := "XML syntax error on line 7: element <Order> closed by </order>"


	if string(content) != expected { 
		t.Errorf("Expected %s, got %s", expected, string(content))
	}

}

func TestXMLNull(t *testing.T) {

	data := ""

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://localhost:8080/process", strings.NewReader(data)) 

	processHandler(res, req) 
	content, _ := ioutil.ReadAll(res.Body)
	// strip spaces to compare string.
	expected := "EOF"


	if string(content) != expected { 
		t.Errorf("Expected %s, got %s", expected, string(content))
	}

}