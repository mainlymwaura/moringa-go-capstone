package main

import "fmt"

func main() {
    fmt.Println("🎉 Welcome to Moringa AI Capstone!")
    fmt.Println("Learning Go with GenAI is awesome!")
    
    // Bonus: Show current time
    fmt.Println("\n📅 Submitted on:", getDate())
}

// A simple function to demonstrate Go syntax
func getDate() string {
    return "April 27, 2026"
}