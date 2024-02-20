package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const Username = "YOUR_USERNAME"
	const Password = "YOUR_PASSWORD"

	payload := map[string]string{
		"source": "universal_ecommerce",
		"url": "https://www.hotels.com/hotel-search?adults=2&children=&d1=2023-12-01&d2=2023-12-05&destination=monaco%2c%20monaco&enddate=2024-02-24&latlong=43.731142%2c7.419758&mapbounds=&pwadialog=&regionid=6050684&rooms=1&semdtl=&sort=recommended&startdate=2024-02-20&theme=&userewards=false&userintent=",
	}

	jsonValue, _ := json.Marshal(payload)

	client := &http.Client{}
	request, _ := http.NewRequest("POST",
		"https://realtime.oxylabs.io/v1/queries",
		bytes.NewBuffer(jsonValue),
	)

	request.SetBasicAuth(Username, Password)
	response, _ := client.Do(request)

	responseText, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseText))
}