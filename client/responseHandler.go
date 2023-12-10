package client

import (
	"encoding/json"
	"fmt"

	dto "addressBookClient/models/dto"
)

func responseHandler(body []byte) {
	fmt.Println("--Response--")

	var response dto.Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\tResult: ", response.Result)
	if response.Result == "error" {
		fmt.Println("\tserver returned error: ", response.Error)
		fmt.Println("--End Response--")
		return
	}
	fmt.Println("--End Response--")

	if response.Data == nil {
		return
	}

	var records dto.Records
	err = json.Unmarshal(response.Data, &records)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(records) == 0 {
		return
	}

	fmt.Println("--Get Records--")
	for i, record := range records {
		fmt.Println("\tRecord ", i)
		fmt.Println("\t\tName: ", record.Name)
		fmt.Println("\t\tLastName: ", record.LastName)
		fmt.Println("\t\tLastName: ", record.MiddleName)
		fmt.Println("\t\tAddress: ", record.Address)
		fmt.Println("\t\tPhone: ", record.Phone)
	}
	fmt.Println("--END--")
}
