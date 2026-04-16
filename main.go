package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ; ; {
		fmt.Println(("Pokedex > "))
		scanner.Scan();
		input := scanner.Text();
		word := cleanInput(input);
		response := fmt.Sprintf("Your command was: %s", word[0]);
		fmt.Println(response)
	}
}

