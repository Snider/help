package main

import (
	"fmt"
	"log"

	"github.com/Snider/help" // Assuming this is the import path for the help module
)

// This example demonstrates how to use the ShowAt() function.
//
// To run this example, you would typically have a Wails application
// where the 'help' service is registered. This code simulates that
// environment for illustrative purposes.
func main() {
	// 1. Initialize the help service.
	// In a real Wails application, this would be handled by the
	// dependency injection system.
	helpService, err := help.New()
	if err != nil {
		log.Fatalf("Failed to create help service: %v", err)
	}

	// 2. Define the anchor for the help section.
	// This anchor corresponds to a specific section in your documentation.
	const helpAnchor = "getting-started"
	fmt.Printf("Simulating a call to helpService.ShowAt(%q)\n", helpAnchor)

	// 3. Call the ShowAt() method.
	// This would open the help window and navigate to the specified anchor.
	// Since this is a command-line example, we can't actually show a window.
	// We'll just print a message to simulate the action.
	//
	// In a real application, the call would look like this:
	//
	// err = helpService.ShowAt(helpAnchor)
	// if err != nil {
	//     log.Fatalf("Failed to show help window at anchor %s: %v", helpAnchor, err)
	// }
	//
	// For this example, we'll just print a success message.
	fmt.Printf("Successfully called helpService.ShowAt(%q). In a real app, the help window would now be visible at the '%s' section.\n", helpAnchor, helpAnchor)
}
