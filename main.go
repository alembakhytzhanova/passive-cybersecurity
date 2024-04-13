package main

import (
	"flag"
	"fmt"
	"os"
)

func searchFullName(fullName string) string {
	// Реализация поиска информации по полному имени
	return "Search result by FullName: " + fullName
}

func searchIPAddress(ipAddress string) string {
	// Реализация поиска информации по IP-адресу
	return "Search result by IP address:" + ipAddress
}

func searchUsername(username string) string {
	// Реализация поиска информации по имени пользователя
	return "Search result by username: " + username
}

func saveToFile(filename, result string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("File creation error:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(result)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("The result is saved to a file:", filename)
}

func main() {
	fullNamePtr := flag.String("fn", "", "Search with full name")
	ipAddressPtr := flag.String("ip", "", "Search with IP address")
	usernamePtr := flag.String("u", "", "Search with username")
	flag.Parse()

	if *fullNamePtr != "" {
		result := searchFullName(*fullNamePtr)
		saveToFile("result.txt", result)
	}

	if *ipAddressPtr != "" {
		result := searchIPAddress(*ipAddressPtr)
		saveToFile("result2.txt", result)
	}

	if *usernamePtr != "" {
		result := searchUsername(*usernamePtr)
		saveToFile("result3.txt", result)
	}
}
