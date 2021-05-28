// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program shows what you're doing today.

package main

import "fmt"

func main() {
	var userAge int
	var userDay string

	fmt.Println("Enter your age:")
	fmt.Println()
	fmt.Print("Age: ")
	fmt.Scanln(&userAge)
  fmt.Println("Is it a weekday?:")
	fmt.Print("yes or no: ")
	fmt.Scanln(&userDay)

	if userDay == "no" {
		print("Time to relax for the weekend!! :D");
	} else if userAge >= 18 && userDay == "yes" {
		print("Time for work!");
	} else if userAge < 18 && userDay == "yes" {
		print("Time for school!");
	} else {
    print("Sorry, that's invalid. Try again.");
  }
}
