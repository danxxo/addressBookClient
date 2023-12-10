package client

import (
	"bufio"
	"fmt"

	dto "addressBookClient/models/dto"
)

// func for scanning terminal input, reurns the gotten record and num of fields
func getRecordFromTerminal(scanner *bufio.Scanner) dto.Record {

	fmt.Print("Enter Name: ")
	scanner.Scan()
	Name := scanner.Text()

	fmt.Print("Enter LastName: ")
	scanner.Scan()
	LastName := scanner.Text()

	fmt.Print("Enter MiddleName: ")
	scanner.Scan()
	MiddleName := scanner.Text()

	fmt.Print("Enter Address: ")
	scanner.Scan()
	Address := scanner.Text()

	fmt.Print("Enter Phone: ")
	scanner.Scan()
	Phone := scanner.Text()

	return dto.Record{Name: Name, LastName: LastName, MiddleName: MiddleName, Address: Address, Phone: Phone}
}
