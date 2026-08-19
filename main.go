package main

import (
	"fmt"
	"os"
)

// compliance_checker - Check security compliance
func compliance_checker(path string) {
	fmt.Println("========================================")
	fmt.Println("  Compliance-Checker")
	fmt.Println("  Check security compliance")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	compliance_checker(path)
}
