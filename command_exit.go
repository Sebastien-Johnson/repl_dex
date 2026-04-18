package main

import (
	"fmt"
	"os"
)

//each callback func gets a file
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}




	

