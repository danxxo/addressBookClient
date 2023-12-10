package main

import (
	"bufio"
	"fmt"
	"os"

	client "addressBookClient/client"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Create a Record")
		fmt.Println("2. Get a Record")
		fmt.Println("3. Update record by phone. Phone is NEEDABLE!")
		fmt.Println("4. Delete record by Phone. Phone is NEEDABLE!")
		fmt.Println("5. Exit")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			client.AddRecord(scanner)
		case "2":
			client.GetRecord(scanner)
		case "3":
			client.UpdateNotes(scanner)
		case "4":
			client.DeleteRecord(scanner)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}

		fmt.Println()
	}
}
