// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains the JS functions for index.html

"use strict"

// Defining function that actives when user presses "Pay" Button
function userInputClick() {
  // Get input from user and store it in const variable
  const hours = parseInt(document.getElementById("hours-entered").value)
  const rate = parseInt(document.getElementById("rate-entered").value)

  // Making calculations
  const takeHome = (hours*rate)*(1.00-0.18)
  const govTake = (hours*rate)*0.18

  // Output calculations
  document.getElementById('pay').innerHTML = "Your pay will be: $" + takeHome.toFixed(2)
  document.getElementById('gov').innerHTML = "The government will take: $" + govTake.toFixed(2)
}