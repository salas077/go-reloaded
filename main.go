package main

import (
	"fmt" // For printing messages to the screen
	"os"  // For getting command line arguments
)

// main is the starting point of our program
// This is where everything begins when we run the program
func main() {
	// os.Args contains the command line arguments
	// os.Args[0] is the program name
	// os.Args[1] is the first argument (input file)
	// os.Args[2] is the second argument (output file)
	
	// Check if user ran program without any arguments
	if len(os.Args) == 1 {
		// No arguments provided, run our tests instead
		fmt.Println("No arguments provided, running tests...")
		testAll() // Run all our test cases
		return   // Exit the program
	}
	
	// Check if user provided exactly 2 arguments (input file and output file)
	if len(os.Args) != 3 {
		// Wrong number of arguments, show how to use the program
		fmt.Println("Usage: go run . <input_file> <output_file>")
		fmt.Println("Example: go run . sample.txt result.txt")
		return // Exit the program
	}

	// Get the file names from command line arguments
	inputFile := os.Args[1]  // First argument is input file
	outputFile := os.Args[2] // Second argument is output file
	
	// Create our text processing pipeline
	// This is like setting up a factory with different workers
	pipeline := NewPipeline()
	
	// Process the input file and create output file
	// This does all the text transformations
	err := pipeline.Process(inputFile, outputFile)
	
	// Check if there was an error during processing
	if err != nil {
		// ErrorHandler already displayed the error message
		return // Exit the program
	}
	
	// Success! Tell the user we're done
	fmt.Printf("Successfully processed %s and created %s\n", inputFile, outputFile)
}