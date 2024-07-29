package main

import (
	"flag"
	"fmt"
	"os"

	headers "github.com/Jesserc/package-headers-go"
)

func main() {
	// Define command-line flags
	packageName := flag.String("package", "", "The name of the package")
	description := flag.String("description", "", "Description of the package")
	example := flag.String("example", "", "Example usage of the package")
	stdout := flag.Bool("stdout", false, "Whether to copy and also print result to standard output")

	flag.Parse()

	if *packageName == "" || *description == "" {
		fmt.Println("Package name and description are required.")
		flag.Usage()
		os.Exit(1)
	}

	headers.GenerateHeader(packageName, description, example, stdout)
}
