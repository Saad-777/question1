package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertToAbsolute(currentDir, relativePath string) string {
	// Normalize the current directory by removing trailing slash if present
	currentDir = strings.TrimSuffix(currentDir, "/")

	// If relative path starts with /, it's already absolute
	if strings.HasPrefix(relativePath, "/") {
		return normalizePath(relativePath)
	}

	// Split current directory and relative path into components
	currentDirParts := strings.Split(currentDir, "/")
	// Remove empty strings from currentDirParts
	currentDirParts = removeEmpty(currentDirParts)

	// Handle relative path components
	relativeParts := strings.Split(relativePath, "/")
	resultParts := make([]string, len(currentDirParts))
	copy(resultParts, currentDirParts)

	for _, part := range relativeParts {
		switch part {
		case "", ".":
			// Skip empty parts and current directory markers
			continue
		case "..":
			// Go up one directory level if possible
			if len(resultParts) > 0 {
				resultParts = resultParts[:len(resultParts)-1]
			}
		default:
			// Add the path component
			resultParts = append(resultParts, part)
		}
	}

	// Construct the final path
	if len(resultParts) == 0 {
		return "/"
	}
	return "/" + strings.Join(resultParts, "/")
}

// Helper function to remove empty strings from slice
func removeEmpty(s []string) []string {
	result := []string{}
	for _, str := range s {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// Helper function to normalize a path
func normalizePath(path string) string {
	// Split path into components
	parts := strings.Split(path, "/")
	resultParts := []string{}

	for _, part := range parts {
		switch part {
		case "", ".":
			continue
		case "..":
			if len(resultParts) > 0 {
				resultParts = resultParts[:len(resultParts)-1]
			}
		default:
			resultParts = append(resultParts, part)
		}
	}

	if len(resultParts) == 0 {
		return "/"
	}
	return "/" + strings.Join(resultParts, "/")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter current directory (e.g., /home/myhome/myfolder1): ")
	currentDir, _ := reader.ReadString('\n')
	currentDir = strings.TrimSpace(currentDir)

	fmt.Print("Enter relative path (e.g., ../myfolder2/mydocument.txt): ")
	relativePath, _ := reader.ReadString('\n')
	relativePath = strings.TrimSpace(relativePath)

	absPath := convertToAbsolute(currentDir, relativePath)
	fmt.Printf("\nAbsolute path: %s\n", absPath)
}
