package client

import (
	dto "addressBookClient/models/dto"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AddRecord(scanner *bufio.Scanner) {
	fmt.Println("Lets create a new Record!")

	rec := getRecordFromTerminal(scanner)

	jsonRequestBytes, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("addRecord()", err)
	}

	url := "http://localhost:8000/add"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("addRecord()", err)
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("recordAdd()", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("addRecord()", err)
		return
	}
	responseHandler(body)
}

func GetRecord(scanner *bufio.Scanner) {
	rec := getRecordFromTerminal(scanner)

	jsonRequestBytes, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}

	url := "http://localhost:8000/get"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseHandler(body)
}

func UpdateNotes(scanner *bufio.Scanner) {
	fmt.Println("Let s Update records. Dont forget fill Phone field")
	rec := getRecordFromTerminal(scanner)

	jsonRequestBytes, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}

	url := "http://localhost:8000/update"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getRecord()", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseHandler(body)
}

func DeleteRecord(scanner *bufio.Scanner) {

	fmt.Println("To delete record Enter the Phone number")
	fmt.Print("Enter Phone: ")

	// Scanner
	scanner.Scan()
	Phone := scanner.Text()

	// Marshaling
	jsonRequestBytes, err := json.Marshal(dto.Record{Phone: Phone})
	if err != nil {
		fmt.Println("getRecord()", err)
	}

	// POST request
	url := "http://localhost:8000/delete"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBytes))
	if err != nil {
		fmt.Println("getRecord()", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Response handling
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseHandler(body)
}
