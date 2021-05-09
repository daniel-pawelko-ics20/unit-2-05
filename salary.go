// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that calculates salary

package main

import (
    "fmt"
    "github.com/leekchan/accounting"
)


// This function displays currency
func main() {
  // Defining variables
  accountingFormater := accounting.Accounting{Symbol: "$", Precision: 2}
	var hours float32
	var rate float32

	fmt.Println("Calculate Salary")
	fmt.Println()

	// User Input
	fmt.Println("Number of hours worked: ")
	fmt.Scanln(&hours)
  fmt.Println()

	fmt.Println("Hourly rate: $")
	fmt.Scanln(&rate)
  fmt.Println()

  // Print out pay and gov take away
  fmt.Println("Your pay will be:",accountingFormater.FormatMoney((hours*rate)*(1.00-0.18)))
  fmt.Println("The government will take:",accountingFormater.FormatMoney((hours*rate)*0.18)) 

}
