package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// URL of the REST API endpoint you want to call
	ans := getPhoneNumbers("Afganistan", "1234567890")
	fmt.Println(ans)
	ans1 := getPhoneNumbers("Oceania", "1234567890")
	fmt.Println(ans1)

	ans2 := getPhoneNumbers("Puerto Rico", "1234567890")
	fmt.Println(ans2)
}

type Response struct {
	Data []Country `json:"data"`
}

type Country struct {
	Name         string   `json:"name"`
	CallingCodes []string `json:"callingCodes"`
}

func getPhoneNumbers(country string, phoneNumber string) string {

	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/countries?name=%s", url.QueryEscape(country))

	//url := url.QueryEscape(url1)

	// Print the encoded URL
	fmt.Println(url)
	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return "-1"
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		fmt.Println("Error reading response:", err)
		return "-1"
	}

	// Print response
	fmt.Println(string(body))

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return "-1"
	}

	if len(response.Data) == 0 {
		return "-1"
	}

	countryData := response.Data[0]
	if countryData.CallingCodes == nil || len(countryData.CallingCodes) == 0 {
		return "-1"
	}
	lastCallingCode := countryData.CallingCodes[len(countryData.CallingCodes)-1]
	return fmt.Sprintf("%s %s", lastCallingCode, phoneNumber)
}
