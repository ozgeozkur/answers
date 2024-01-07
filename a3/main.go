package main

import "fmt"

// generateOutput is a recursive function that prints the squares of numbers
// up to the given limit. It avoids printing squares that exceed the limit.
// It takes two parameters: n - current number, limit - the maximum limit.
func generateOutput(n, limit int) {
	if n > 0 {
		// Call the function for the previous number in a recursive manner
		generateOutput(n-1, limit)

		// Calculate the square of the current number
		square := n * n

		// Print the square if it's less than or equal to the limit
		if square <= limit {
			fmt.Println(square)
		}
	}
}

func main() {
	// Example input
	input := 9

	// Generate the output
	generateOutput(input, input)
}
