package main

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/jprobinson/go-x12"
)

func main() {
	// Load the X12 message as a string
	message := "ISA*00*          *00*          *ZZ*ReceiverID     *ZZ*Sender         *120918*1234*^*00501*000000001*0*T*:~GS*HC*ReceiverID*Sender*20120918*1234*1*X*005010X222A1~ST*837*0001*005010X222A1~"

	// Parse the message
	parser := x12.NewParser(x12.V5010)
	transactions, _ := parser.ParseString(message)
	
	// Validate the X12 transaction using an XML schema
	doc := etree.NewDocument()
	if err := doc.ReadFromString(message); err != nil {
		fmt.Println("Error reading X12 message:", err)
		return
	}

	schema := etree.NewXMLSchema(x12.V5010.Schema())
	if err := doc.Validate(schema); err != nil {
		fmt.Println("Error validating X12 message:", err)
		return
	}
	
	// Print the validation status
	fmt.Println("X12 837 transaction is valid.")
}
