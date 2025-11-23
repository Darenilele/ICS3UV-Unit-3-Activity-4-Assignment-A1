/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will demonstrate if statements 
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Variables
	var nutsAsString string = ""
	var boltAsString string = ""
	var washerAsString string = ""
	var nutsAsNumber int = 0
	var boltAsNumber int = 0
	var washerAsNumber int = 0
	var total int = 0

	// Create new reader
	var reader = bufio.NewReader(os.Stdin)

	// INPUT (Displaying Statements)
	fmt.Print("How many bolts would you like to purchase? ")
	boltAsString, _ = reader.ReadString('\n')
	boltAsString = strings.TrimSpace(boltAsString)
	boltAsNumber, _ = strconv.Atoi(boltAsString)

	fmt.Print("How many nuts would you like to purchase? ")
	nutsAsString, _ = reader.ReadString('\n')
	nutsAsString = strings.TrimSpace(nutsAsString)
	nutsAsNumber, _ = strconv.Atoi(nutsAsString)

	fmt.Print("How many washers would you like to purchase? ")
	washerAsString, _ = reader.ReadString('\n')
	washerAsString = strings.TrimSpace(washerAsString)
	washerAsNumber, _ = strconv.Atoi(washerAsString)

	// PROCESSING (Calculations)
	if boltAsNumber > nutsAsNumber {
		fmt.Println("Check the Order - not enough nuts for the bolts you purchased.")
	} else if boltAsNumber > washerAsNumber {
		fmt.Println("Check the Order - not enough washers for the bolts you purchased.")
	} else {
		fmt.Println("Your Order Is Okay")
			total = boltAsNumber + nutsAsNumber + washerAsNumber
		fmt.Println("Your total cost of the order is", total, "$")
	}

	fmt.Println("\nDone.")
}
