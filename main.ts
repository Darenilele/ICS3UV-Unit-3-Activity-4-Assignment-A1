/**
 * @author Daren Ileleji
 * @versopn 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will demonstarte if statments 
 */


// Variables
let nutsAsString: string = "";
let boltAsString: string = "";
let washerAsString: string = "";
let nutsAsNumber: number = 0;
let boltAsNumber: number = 0;
let washerAsNumber: number = 0;

// INPUT (Displaying Statements)
boltAsString = prompt("How many bolts would you like to purchase?") || "0";
boltAsNumber = parseInt(boltAsString);

nutsAsString = prompt("How many nuts would you like to purchase?") || "0";
nutsAsNumber = parseInt(nutsAsString);

washerAsString = prompt("How many washers would you like to purchase?") || "0";
washerAsNumber = parseInt(washerAsString);


// PROCESS (Calculations)
if (boltAsNumber > nutsAsNumber) {
  console.log("Check the Order - not enough nuts for the bolts you purchased.");
} else if (boltAsNumber > washerAsNumber) {
    console.log("Check the Order - not enough washers for the bolts you purchased.")
} else {
  console.log("Your Order Is Okay")
  let total: number = (boltAsNumber + nutsAsNumber + washerAsNumber);
  console.log("Your total cost of the order is " + total + "$")
}

console.log("\nDone.")
